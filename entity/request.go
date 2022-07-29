package entity

import "time"

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
	To            string     `json:"to"`
	GeneralStatus string     `json:"generalStatus"`
	SentSince     *time.Time `json:"sentSince"`
	SentUntil     *time.Time `json:"sentUntil"`
	Limit         int        `json:"limit"`
	MCC           string     `json:"mcc"`
	MNC           string     `json:"mnc"`
}
