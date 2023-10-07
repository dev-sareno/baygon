package workerhandler

import (
	"encoding/json"
	"github.com/dev-sareno/ginamus/context"
	"github.com/dev-sareno/ginamus/db"
	"github.com/dev-sareno/ginamus/dns"
	"github.com/dev-sareno/ginamus/dto"
	"github.com/dev-sareno/ginamus/mq"
	"log"
	"os"
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
	case "CNAME":
		const activityId = "lookup-cname"
		resolver := dns.CnameResolver{Child: &dns.EmptyResolver{}}
		job := Lookup(ctx, activityId, &resolver)

		// update database
		if ok := db.UpdateJob(job); !ok {
			return
		}

		// move job to lookup A
		mq.PublishToLookupA(ctx.MqChannel, job)
		break
	case "A":
		const activityId = "lookup-a"
		resolver := dns.IpResolver{}
		job := Lookup(ctx, activityId, &resolver)

		// update database
		if ok := db.UpdateJob(job); !ok {
			return
		}

		// done. A lookup is the end of lookup
		break
	default:
		log.Printf("invalid dns lookup type %s\n", lookupType)
	}
}
