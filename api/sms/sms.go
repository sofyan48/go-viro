package sms

import (
	"encoding/json"

	"github.com/sofyan48/go-viro/consts"
	"github.com/sofyan48/go-viro/entity"
	"github.com/sofyan48/go-viro/pkg/requester"
	"github.com/sofyan48/go-viro/pkg/utils"
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
