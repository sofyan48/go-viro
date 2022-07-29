package entity

//go:generate easytags $GOFILE json

type AdvancePayload struct {
	To   string `json:"to"`
	Text string `json:"text"`
}

type MultiPayload struct {
	To   []string `json:"to"`
	Text string   `json:"text"`
}
