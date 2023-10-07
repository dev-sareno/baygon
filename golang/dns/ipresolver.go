package dns

import (
	"fmt"
	"net"
	"strings"
)

type IpResolver struct {
	DnsResolver
	value string
}

func (slf *IpResolver) SetValue(v string) {
	slf.value = v
}

func (slf *IpResolver) Resolve() ([]Resolution, error) {
	fmt.Printf("ip lookup for %s..\n", slf.value)
	result, err := net.LookupIP(slf.value)
	if err != nil {
		return []Resolution{}, fmt.Errorf("ip lookup failed for %s. %s", slf.value, err.Error())
	}
	fmt.Printf("\tvalue: %s\n", result)
	fmt.Println("\tdone.")

	var ips []string
	for _, v := range result {
		ips = append(ips, v.String())
	}
	return []Resolution{{Type: "IP", Value: strings.Join(ips, "\n")}}, nil
}
