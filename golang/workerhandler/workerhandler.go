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

func HandleJob(data []byte) error {
	log.Printf("Received a message: %s", data)
	var job dto.Job
	if err := json.Unmarshal(data, &job); err != nil {
		return fmt.Errorf("invalid job input %s", err)
	}
	if job.Data.Type != 0 {
		log.Printf("unsuported job type %d\n", job.Data.Type)
		return nil
	}
	return nil
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
