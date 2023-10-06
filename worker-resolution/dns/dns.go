package dns

import (
	"fmt"
	"os"
	"strings"
)

type Resolution struct {
	Type  string
	Value string
}

type DnsResolver interface {
	SetValue(v string)
	Resolve() ([]Resolution, error)
}

type DnsRecord struct {
	Value string
	Child DnsResolver
}

func (slf *DnsRecord) Run() {
	slf.Child.SetValue(slf.Value)
	results, err := slf.Child.Resolve()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var items []string
	for _, result := range results {
		items = append(items, result.Value)
	}
	fmt.Printf("%v\n", strings.Join(items, "\n"))
}
