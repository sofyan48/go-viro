package sms

import (
	"github.com/sofyan48/go-viro/pkg/requester"
)

type smsAPI struct {
	http     requester.Contract
	senderID string
	header   map[string]string
	payload  map[string]interface{}
	url      string
	typeSend string
}

func NewSmsAPI(senderID, apikey string) SMSInterface {
	return &smsAPI{
		http:     requester.New(),
		senderID: senderID,
		header: map[string]string{
			"Content-Type":  "application/json",
			"Accept":        "application/json",
			"Authorization": apikey,
		},
	}
}
