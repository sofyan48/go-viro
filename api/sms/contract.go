package sms

type SMSInterface interface {
	Single(to, text string) *smsAPI
}
