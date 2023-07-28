package mtypus

type RequestPostDo struct {
	Input string `json:"input" binding:"required"`
}
