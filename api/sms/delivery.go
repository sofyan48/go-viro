package sms

import (
	"encoding/json"
	"strconv"

	urlTools "net/url"

	"github.com/sofyan48/go-viro/consts"
	"github.com/sofyan48/go-viro/entity"
)

func (s *smsAPI) GetReports(limit int) ([]entity.ResultsFormat, error) {
	url := consts.BASE_URL + consts.SMS_GET_REPORTS
	if limit != 0 {
		url = url + "?limit=" + strconv.Itoa(limit)
	}

	response, err := s.http.GET(url, s.header)
	if err != nil {
		return nil, err
	}
	resutl := []entity.ResultsFormat{}
	err = json.Unmarshal(response, &resutl)
	if err != nil {
		return nil, err
	}
	return resutl, nil
}

func (s *smsAPI) GetLogs(params *entity.LogsParameters) ([]entity.ResultsFormat, error) {
	url := consts.BASE_URL + consts.SMS_GET_LOGS
	paramFormat := urlTools.Values{}
	paramFormat.Add("from", s.senderID)
	paramFormat.Add("to", params.To)
	paramFormat.Add("generalStatus", params.GeneralStatus)
	paramFormat.Add("sentSince", params.SentSince.String())
	paramFormat.Add("sentUntil", params.SentUntil.String())
	paramFormat.Add("limit", strconv.Itoa(params.Limit))
	paramFormat.Add("mcc", params.MCC)
	paramFormat.Add("mnc", params.MNC)
	urlWithParams := url + paramFormat.Encode()

	response, err := s.http.GET(urlWithParams, s.header)
	if err != nil {
		return nil, err
	}
	resutl := []entity.ResultsFormat{}
	err = json.Unmarshal(response, &resutl)
	if err != nil {
		return nil, err
	}
	return resutl, nil
}
