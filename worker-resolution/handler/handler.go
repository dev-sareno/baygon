package handler

import "log"

func HandleJob(data []byte) error {
	log.Printf("Received a message: %s", data)
	return nil
}
