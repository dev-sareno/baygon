package workerhandler

import (
	"encoding/json"
	"github.com/dev-sareno/ginamus/codec"
	"github.com/dev-sareno/ginamus/context"
	"github.com/dev-sareno/ginamus/dns"
	"github.com/dev-sareno/ginamus/dto"
	"log"
)

func lookupCname(ctx *context.WorkerContext) {
	const activityId = "lookup-cname"

	job := ctx.Job
	job.LastActivityId = activityId // set activity id

	jobInput := job.Data.Input

	jobOutput := dto.ActivityOutput{
		Index:   int32(len(job.Data.Outputs)),
		Id:      activityId,
		Data:    "",
		IsOk:    true,
		Message: job.LastActivityMessage,
	}
	var lookupResult []string // list of the resolved values
	hasWarning := false
	hasError := false
	for _, v := range jobInput.Domains {
		lookup := dns.CnameResolver{Child: &dns.EmptyResolver{}}
		lookup.SetValue(v)
		result, err := lookup.Resolve()
		if err != nil {
			// lookup failed
			log.Printf("cname lookup failed. %s\n", err)
			lookupResult = append(lookupResult, "") // append empty
			hasError = true
		} else {
			// lookup successful
			// ipv4 lookup is expected length of one since it  doesn't have a child
			if len(result) != 1 {
				log.Println("WARNING: cnameresolver is expected to return a length of 1")
				// this is invalid, consider as failed
				lookupResult = append(lookupResult, "") // append empty
				hasWarning = true
				continue
			}
			lookupResult = append(lookupResult, result[0].Value)
		}
	}

	var msg string
	if hasWarning {
		msg = "completed with warning"
	} else if hasError {
		msg = "completed with errors"
	} else {
		msg = "completed"
	}

	// encode result
	b, _ := json.Marshal(&lookupResult)

	// finalize job output
	jobOutput.Data = string(b)
	jobOutput.Message = msg
	jobOutput.IsOk = true
	job.LastActivityMessage = msg
	job.LastActivityIsOk = true

	job.Data.Outputs = append(job.Data.Outputs, jobOutput)

	codec.Encode(job)

	// done. cname lookup is the end of lookup
}
