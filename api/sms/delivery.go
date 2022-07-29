package sms

import (
	"encoding/json"

	"github.com/sofyan48/go-viro/consts"
	"github.com/sofyan48/go-viro/entity"
)

func (s *smsAPI) GetReports() ([]entity.ResultsFormat, error) {
	url := consts.BASE_URL + consts.SMS_GET_REPORTS
	response, err := s.http.GET(url, s.header)
	if err != nil {
		return nil, err
	}
	resutl := []entity.ResultsFormat{}
	err = json.Unmarshal(response, resutl)
	if err != nil {
		return nil, err
	}
	return resutl, nil
}
