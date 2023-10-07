package workerhandler

import (
	"github.com/dev-sareno/ginamus/codec"
	"github.com/dev-sareno/ginamus/context"
	"github.com/dev-sareno/ginamus/dns"
	"github.com/dev-sareno/ginamus/dto"
	"log"
)

func Lookup(ctx *context.WorkerContext, activityId string, resolver dns.DnsResolver) *dto.Job {
	job := ctx.Job
	job.LastActivityId = activityId // set activity id

	jobInput := job.Data.Input

	jobOutput := dto.ActivityOutput{
		Index:   int32(len(job.Data.Outputs)),
		Id:      activityId,
		Data:    []string{},
		IsOk:    true,
		Message: job.LastActivityMessage,
	}
	var lookupResult []string // list of the resolved values
	hasError := false
	for _, v := range jobInput.Domains {
		resolver.SetValue(v)
		result, err := resolver.Resolve()
		if err != nil {
			// lookup failed
			log.Printf("lookup failed. %s\n", err)
			lookupResult = append(lookupResult, "") // append empty
			hasError = true
		} else {
			// lookup successful
			if len(result) > 0 {
				first := result[0] // length is always 1
				lookupResult = append(lookupResult, first.Value)
			} else {
				// value not found
				lookupResult = append(lookupResult, "") // append empty
			}
		}
	}

	var msg string
	if hasError {
		msg = "completed with errors"
	} else {
		msg = "completed"
	}

	// finalize job output
	jobOutput.Data = lookupResult
	jobOutput.Message = msg
	jobOutput.IsOk = true
	job.LastActivityMessage = msg
	job.LastActivityIsOk = true

	job.Data.Outputs = append(job.Data.Outputs, jobOutput)

	codec.Encode(job)

	return ctx.Job
}
