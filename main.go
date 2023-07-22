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
		log.Fatal(err)
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
		log.Fatal(err)
	}
}

func handle(domain string) {
	cname, _ := lookupCNAME(domain)
	fmt.Println(cname)

	a, _ := lookupA(domain)
	fmt.Println(a)
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
