package workerhandler

import (
	"encoding/json"
	"fmt"
	"github.com/dev-sareno/ginamus/context"
	"github.com/dev-sareno/ginamus/db"
	"github.com/dev-sareno/ginamus/dns"
	"github.com/dev-sareno/ginamus/dto"
	"github.com/dev-sareno/ginamus/mq"
	"log"
	"os"
	"strings"
)

func HandleJob(ctx *context.WorkerContext, data []byte) {
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

	ctx.Job = &job // assign job to context

	// handle resolution
	handleDnsResolution(ctx)
}

func handleDnsResolution(ctx *context.WorkerContext) {
	lookupType := os.Getenv("WORKER_DNS_LOOKUP_TYPE")
	switch lookupType {
	case "A":
		const activityId = "lookup-a"
		resolver := dns.IpResolver{}
		job := Lookup(ctx, activityId, &resolver)

		// update database
		if ok := db.UpdateJob(job); !ok {
			return
		}

		// move job to lookup cname
		mq.PublishToLookupCname(ctx.MqChannel, job)
		break
	case "CNAME":
		const activityId = "lookup-cname"
		resolver := dns.CnameResolver{Child: &dns.EmptyResolver{}}
		_ = Lookup(ctx, activityId, &resolver)
		// done. cname lookup is the end of lookup
		break
	default:
		log.Printf("invalid dns lookup type %s\n", lookupType)
	}
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
