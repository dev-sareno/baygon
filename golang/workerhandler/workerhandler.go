package workerhandler

import (
	"encoding/json"
	"fmt"
	"github.com/dev-sareno/ginamus/dns"
	"github.com/dev-sareno/ginamus/dto"
	"log"
	"os"
	"strings"
)

func HandleJob(data []byte) {
	log.Printf("Received a message: %s", data)

	var job dto.Job
	if err := json.Unmarshal(data, &job); err != nil {
		log.Printf("invalid job input %s\n", err)
		return
	}
	if job.Data.Type != 0 {
		log.Printf("unsuported job type %d\n", job.Data.Type)
		return
	}

	result, err := handleDnsResolution(job)
	if err != nil {
		log.Printf("dns resolution failed. %s\n", err)
		return
	}

	log.Printf("dns resolution result: %s\n", result)
}

func handleDnsResolution(job dto.Job) (string, error) {
	lookupType := os.Getenv("WORKER_DNS_LOOKUP_TYPE")
	switch lookupType {
	case "a":
		// TODO: implement me
		//resolver := dns.IpResolver{}
		//resolver.SetValue(job.Data.Input.Domains)
		break
	case "cname":
	default:
		return "", fmt.Errorf("invalid dns lookup type %s\n", lookupType)
	}
	return "", nil
}

func test() {
	domain := "github.com"

	ipResolver := dns.IpResolver{}
	cnameResolver := dns.CnameResolver{Child: &ipResolver}
	recordResolver := dns.RecordResolver{Child: &cnameResolver}
	recordResolver.SetValue(domain)
	result, err := recordResolver.Resolve()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var items []string
	for _, item := range result {
		items = append(items, item.Value)
	}
	fmt.Printf("%v\n", strings.Join(items, "\n"))
}
