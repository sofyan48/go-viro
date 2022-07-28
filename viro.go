package viro

import "github.com/sofyan48/go-viro/api/sms"

type Viro struct {
	ApiKey   string
	SenderID string
}

func NewViro(senderID, apiKey string) *Viro {
	return &Viro{
		ApiKey:   apiKey,
		SenderID: senderID,
	}
}

func (v *Viro) SMS() sms.SMSInterface {
	return sms.NewSmsAPI(v.SenderID, v.ApiKey)
}
