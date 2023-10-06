package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"sort"
	"strings"
)

type IP struct {
	Value  string
	IsIPv4 bool
}

type CNAME struct {
	Value string
	IPs   []IP
}

type Node struct {
	Text string
}

func main() {
	var targetPath string
	{ // validate arguments
		args := os.Args
		idx := -1
		for i, arg := range args {
			if arg == "-f" || arg == "--file" {
				idx = i
				break
			}
		}
		fileArgIsValid := idx != -1 && len(args) > (idx+1)
		if fileArgIsValid {
			// check if file exists
			fileArgValue := args[idx+1]
			if _, err := os.Stat(fileArgValue); errors.Is(err, os.ErrNotExist) {
				// file not found
				log.Fatalf(err.Error())
			}
			targetPath = fileArgValue
		} else {
			log.Fatalln("Argument missing --file/-f")
		}
	}

	file, err := os.Open(targetPath)
	if err != nil {
		// maybe access denied?
		log.Fatal(err.Error())
	}
	defer file.Close()

	outFile := makeOutputPath()
	defer outFile.Close()

	{ // write mermaid header
		mermaidHeader := make([]string, 0)
		mermaidHeader = append(mermaidHeader, "```mermaid")
		mermaidHeader = append(mermaidHeader, "flowchart LR")
		mermaidHeader = append(mermaidHeader, "")
		if _, err := outFile.WriteString(strings.Join(mermaidHeader, "\n")); err != nil {
			log.Fatalf("unable to write to output file. %s", err.Error())
		}
	}

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			mermaidText := handle(strings.TrimSpace(line))
			mermaidText = fmt.Sprintf("%s\n", mermaidText) // add new line
			if _, err := outFile.WriteString(mermaidText); err != nil {
				log.Fatalf("unable to write to output file. %s", err.Error())
			}
		}
	}

	{ // write mermaid footer
		mermaidFooter := make([]string, 0)
		mermaidFooter = append(mermaidFooter, "")
		mermaidFooter = append(mermaidFooter, "```")
		mermaidFooter = append(mermaidFooter, "")
		if _, err := outFile.WriteString(strings.Join(mermaidFooter, "\n")); err != nil {
			log.Fatalf("unable to write to output file. %s", err.Error())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err.Error())
	}
}

func makeOutputPath() *os.File {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	outputPath := fmt.Sprintf("%s/%s", cwd, "OUTPUT.md")
	log.Println(outputPath)
	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("failed to create output file. %s\n", err.Error())
	} else if file == nil {
		log.Fatalln("failed to create output file. error: file is nil")
	}
	return file
}

func handle(domain string) string {
	log.Printf("lookup: %s\n", domain)

	// CNAME first
	var cname *CNAME = nil
	if v, err := lookupCNAME(domain); err == nil {
		// CNAME found
		log.Printf("CNAME found %s\n", v)

		// check if identical
		if !strings.HasPrefix(v, domain) {
			// create CNAME
			cname = &CNAME{v, make([]IP, 0)}

			// resolve IPs
			cname.IPs = getIPs(cname.Value)
		}
	}

	// then IP
	ips := getIPs(domain)

	// create nodes
	nodes := createNodeChain(domain, ips, cname)

	// generate Mermaid output
	mermaidText := GenerateOutput(nodes)
	return mermaidText
}

func getIPs(domain string) []IP {
	log.Println("looking for IP")
	result := make([]IP, 0)
	if ips, err := lookupA(domain); err != nil {
		log.Printf("lookup failed. %s\n", err.Error())
	} else {
		sort.Slice(ips, func(i, j int) bool {
			return ips[i] > ips[j] // sort ascending
		})
		for _, ip := range ips {
			isIPv4 := strings.Contains(ip, ".")
			if isIPv4 {
				log.Printf("IPv4 found %s\n", ip)
			} else {
				log.Printf("IPv6 found %s\n", ip)
			}
			result = append(result, IP{ip, isIPv4})
		}
	}
	return result
}

func createNodeChain(domain string, ips []IP, cname *CNAME) []Node {
	var result []Node

	// node: domain
	result = append(result, Node{domain})

	if cname != nil && len(cname.IPs) > 0 {
		// resolves to CNAME

		// node: IP
		cnameResolvesTo := createNodeChain(cname.Value, cname.IPs, nil)
		result = append(result, cnameResolvesTo...)
	} else if len(ips) > 0 {
		// resolves to A/AAAA
		// node: IP
		var v []string
		for _, ip := range ips {
			v = append(v, ip.Value)
		}
		text := strings.Join(v, "<br/>")
		result = append(result, Node{text})
	}
	return result
}

func lookupCNAME(domain string) (string, error) {
	result, err := net.LookupCNAME(domain)
	if err != nil {
		return "", fmt.Errorf("lookup failed. %s", err.Error())
	}

	cname := ""
	for _, v := range result {
		cname += string(v)
	}
	return cname, nil
}

func lookupA(domain string) ([]string, error) {
	returnArr := make([]string, 0)
	result, err := net.LookupIP(domain)
	if err != nil {
		return returnArr, fmt.Errorf("lookup failed. %s", err.Error())
	}

	for _, v := range result {
		returnArr = append(returnArr, v.String())
	}
	return returnArr, nil
}

func GenerateOutput(nodes []Node) string {
	mermaid := make([]string, 0)
	for _, node := range nodes {
		// use hash value as ID
		h := sha1.New()
		if _, err := h.Write([]byte(node.Text)); err != nil {
			log.Printf("failed to create hash. %s\n", err.Error())
			continue
		}
		sha1Hash := hex.EncodeToString(h.Sum(nil))
		shortSha1Hash := sha1Hash[:5]
		nodeText := fmt.Sprintf("%s[%s]", shortSha1Hash, node.Text)
		mermaid = append(mermaid, nodeText)
	}
	mermaidText := strings.Join(mermaid, " --> ")
	return mermaidText
}
