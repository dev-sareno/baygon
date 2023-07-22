package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	result, err := net.LookupCNAME("google.com")
	if err != nil {
		fmt.Printf("Lookup failed. %s\n", err.Error())
		os.Exit(1)
	}

	cname := ""
	for _, v := range result {
		cname += string(v)
	}
	fmt.Println(cname)
}
