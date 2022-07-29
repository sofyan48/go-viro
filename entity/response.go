package entity

import "time"

//go:generate easytags $GOFILE json
type ResultsFormat struct {
	Result  []Result `json:"result,omitempty"`
	Message []Result `json:"message,omitempty"`
}

type Result struct {
	BulkID    string        `json:"bulkId,omitempty"`
	MessageID string        `json:"messageId,omitempty"`
	To        string        `json:"to,omitempty"`
	SentAt    *time.Time    `json:"sentAt,omitempty"`
	DoneAt    *time.Time    `json:"doneAt,omitempty"`
	SmsCount  uint64        `json:"smsCount,omitempty"`
	Price     *PriceFormat  `json:"price,omitempty"`
	Status    *StatusFormat `json:"status,omitempty"`
	Error     *ErrorFormat  `json:"error,omitempty"`
}

type PriceFormat struct {
	PricePerMessage float64 `json:"pricePerMessage"`
	Currency        string  `json:"currency"`
}

type StatusFormat struct {
	GroupID     uint64 `json:"groupId"`
	GroupName   string `json:"groupName"`
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ErrorFormat struct {
	GroupID     uint64 `json:"groupId"`
	GroupName   string `json:"groupName"`
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Permanent   bool   `json:"permanent"`
}
