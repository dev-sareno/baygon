package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
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

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			handle(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err.Error())
	}
}

func handle(domain string) {
	log.Printf("lookup: %s\n", domain)

	// CNAME first
	var cname *CNAME = nil
	if v, err := lookupCNAME(domain); err != nil {
		// CNAME found
		log.Printf("CNAME found %s\n", v)

		// resolve IPs
		ips := getIPs(cname.Value)

		// create CNAME
		cname = &CNAME{v, ips}
	}

	// then IP
	ips := getIPs(domain)

	// create nodes
	nodes := createNodeChain(domain, ips, cname)

	// generate Mermaid output
	GenerateOutput(nodes)
}

func getIPs(domain string) []IP {
	log.Println("looking for IP")
	result := make([]IP, 0)
	if ips, err := lookupA(domain); err != nil {
		log.Printf("lookup failed. %s\n", err.Error())
	} else {
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

	if cname != nil {
		// resolves to CNAME

		// node: CNAME
		result = append(result, Node{cname.Value})

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
		text := strings.Join(v, "\n")
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

func GenerateOutput(nodes []Node) {
	for _, node := range nodes {
		log.Println(node)
	}
}
