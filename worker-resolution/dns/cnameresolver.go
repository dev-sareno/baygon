package dns

import (
	"fmt"
	"net"
	"strings"
)

type CnameResolver struct {
	DnsResolver
	value string
	Child DnsResolver
}

func (slf *CnameResolver) SetValue(v string) {
	slf.value = v
}

func (slf *CnameResolver) Resolve() ([]Resolution, error) {
	fmt.Printf("cname lookup for %s..\n", slf.value)
	result, err := net.LookupCNAME(slf.value)
	if err != nil {
		return []Resolution{}, fmt.Errorf("cname lookup failed for %s. %s", slf.value, err.Error())
	}
	if strings.HasSuffix(result, ".") {
		// remove trailing dot
		result = result[:len(result)-1]
	}

	if result == slf.value {
		// cname value not found
		fmt.Println("\tvalue not found")
	} else {
		// cname value found
		fmt.Printf("\tvalue: %s\n", result)
	}
	fmt.Println("\tdone.")

	slf.Child.SetValue(result)
	childValue, childError := slf.Child.Resolve()
	if childError != nil {
		return []Resolution{}, fmt.Errorf("failed to resolve child. %s", childError.Error())
	}

	if result == slf.value {
		return childValue, nil
	} else {
		return append([]Resolution{{Type: "CNAME", Value: result}}, childValue...), nil
	}
}
