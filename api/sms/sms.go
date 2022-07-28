package sms

import (
	"encoding/json"

	"github.com/sofyan48/go-viro/consts"
	"github.com/sofyan48/go-viro/pkg/requester"
)

type smsAPI struct {
	http     requester.Contract
	senderID string
	header   map[string]string
	payload  map[string]string
	url      string
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

func (s *smsAPI) Single(to, text string) *smsAPI {
	url := consts.BASE_URL + consts.SMS_URL_PATH_SINGLE
	payload := map[string]string{
		"from": s.senderID,
		"to":   to,
		"text": text,
	}
	s.payload = payload
	s.url = url
	return s
}

func (s *smsAPI) Send() ([]byte, error) {
	payload, err := json.Marshal(s.payload)
	if err != nil {
		return nil, err
	}
	return s.http.POST(s.url, s.header, payload)
}
