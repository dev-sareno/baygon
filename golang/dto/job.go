package dto

type ActivityOutput struct {
	Index   int32    `json:"index"`
	Id      string   `json:"id"`
	Data    []string `json:"data"`
	IsOk    bool     `json:"isOk"`
	Message string   `json:"message"`
}

type JobInput struct {
	Domains []string   `json:"domains"`
	Filler  [][]string `json:"filler"`
}

type JobData struct {
	Type    int32            `json:"type"`
	Input   JobInput         `json:"input"`
	Outputs []ActivityOutput `json:"outputs"`
}

type Job struct {
	Id                  string  `json:"id"`
	CreatedAt           string  `json:"createdAt"`
	Data                JobData `json:"data"`
	LastActivityId      string  `json:"lastActivityId"`
	LastActivityIsOk    bool    `json:"lastActivityIsOk"`
	LastActivityMessage string  `json:"lastActivityMessage"`
}
