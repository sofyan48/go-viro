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

type LogsParameters struct {
	To            string   `json:"to"`
	BulkID        []string `json:"bulkId"`
	MessageID     []string `json:"messageId"`
	GeneralStatus string   `json:"generalStatus"`
	SentSince     string   `json:"sentSince"`
	SentUntil     string   `json:"sentUntil"`
	Limit         int      `json:"limit"`
	MCC           string   `json:"mcc"`
	MNC           string   `json:"mnc"`
}
