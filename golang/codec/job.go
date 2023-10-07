package codec

import (
	"encoding/json"
	"github.com/dev-sareno/ginamus/dto"
	"log"
)

func Encode(job *dto.Job) string {
	data, err := json.MarshalIndent(job, "", "  ")
	if err != nil {
		log.Println("unable to json-encode job")
		return "{}" // return empty json
	}
	s := string(data)
	log.Printf("encoded job: %s\n", s)
	return s
}
