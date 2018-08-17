package simboss

import (
	"net/url"
	"time"
	"encoding/json"
	)

type SmsService struct {
	Client *Client
}

// 短信下发接口
func (s *SmsService) Send(params url.Values) error {
	_, err := s.Client.Post("/sms/send", params)
	if err != nil {
		return err
	}
	return nil
}

type Page struct {
	PageNo int `json:"pageNo"`
	PageSize int `json:"pageSize"`
	Total int `json:"total"`
}

type Sms struct {
	Id int64 `json:"id"`
	Iccid string `json:"iccid"`
	Carrier string `json:"carrier"`
	Text string `json:"text"`
	SendTime time.Time `json:"sendTime"`
	Type string `json:"type"`
}

type SmsList struct {
	Page Page `json:"page"`
	List []Sms `json:"list"`
}

// 短信查询
func (s *SmsService) List(params url.Values) (*SmsList, error) {
	smsList := &SmsList{
		Page: Page{},
		List: make([]Sms, 0),
	}
	body, err := s.Client.Post("/sms/list", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, smsList); err != nil {
		return nil, err
	}
	return smsList, nil
}