package sms

import (
	"encoding/json"

	"github.com/sofyan48/go-viro/consts"
	"github.com/sofyan48/go-viro/entity"
	"github.com/sofyan48/go-viro/pkg/utils"
)

func (s *smsAPI) Advance(payload []entity.AdvancePayload) *smsAPI {
	var postData []map[string]interface{}
	for k, i := range payload {
		postField := map[string]interface{}{
			"from": s.senderID,
			"destination": map[string]interface{}{
				"to": utils.SerializeNumber(i.To),
			},
			"text":     i.Text,
			"smsCount": k,
		}
		postData = append(postData, postField)
	}
	s.url = consts.BASE_URL + consts.SMS_URL_PATH_ADVANCE
	s.payload = map[string]interface{}{
		"messages": postData,
	}
	s.typeSend = "advance"
	return s
}

func (s *smsAPI) Multi(payload []entity.MultiPayload) *smsAPI {
	var postData []map[string]interface{}
	for k, i := range payload {
		var msisdn []string
		for c := range i.To {
			msisdn = append(msisdn, utils.SerializeNumber(i.To[c]))
		}
		postField := map[string]interface{}{
			"from": s.senderID,
			"destination": map[string]interface{}{
				"to": msisdn,
			},
			"text":     i.Text,
			"smsCount": k,
		}
		postData = append(postData, postField)
	}
	s.url = consts.BASE_URL + consts.SMS_URL_PATH_MULTI
	s.payload = map[string]interface{}{
		"messages": postData,
	}
	s.typeSend = "multi"
	return s
}

func (s *smsAPI) Single(to, text string) *smsAPI {
	msidn := utils.SerializeNumber(to)
	url := consts.BASE_URL + consts.SMS_URL_PATH_SINGLE
	payload := map[string]interface{}{
		"from": s.senderID,
		"to":   msidn,
		"text": text,
	}

	s.payload = payload
	s.url = url
	s.typeSend = "single"
	return s
}

func (s *smsAPI) Send() (*entity.ResultsFormat, error) {
	payload, err := json.Marshal(s.payload)
	if err != nil {
		return nil, err
	}
	response, err := s.http.POST(s.url, s.header, payload)
	resultFormat := &entity.ResultsFormat{}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, resultFormat)
	if err != nil {
		return nil, err
	}
	return resultFormat, nil
}
