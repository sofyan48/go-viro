package sms

import "github.com/sofyan48/go-viro/entity"

type SMSInterface interface {
	Single(to, text string) *smsAPI
	Advance(payload []entity.AdvancePayload) *smsAPI
	Multi(payload []entity.MultiPayload) *smsAPI
	GetReports() ([]entity.ResultsFormat, error)
}
