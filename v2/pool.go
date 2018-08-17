package simboss

import (
	"encoding/json"
	"net/url"
)

type PoolService struct {
	Client *Client
}

type Pool struct {
	Id int `json:"id"`
	PoolSpecification int `json:"poolSpecification"`
	Carrier string `json:"carrier"`
	TotalVolume float64 `json:"totalVolume"`
	UseVolume float64 `json:"useVolume"`
	LeftVolume float64 `json:"leftVolume"`
	PackageVolume float64 `json:"packageVolume"`
	UseRate float64 `json:"useRate"`
	TotalCount float64 `json:"totalCount"`
	CurrentActivationCount int `json:"currentActivationCount"`
	CurrentDeactivationCount int `json:"currentDeactivationCount"`
	CurrentInventoryCount int `json:"currentInventoryCount"`
	CurrentTestingCount int `json:"currentTestingCount"`
	CurrentRetiredCount int `json:"currentRetiredCount"`
	ActivationReadyCount int `json:"activationReadyCount"`
}

// 流量池详情
func (p * PoolService) Detail(params url.Values) (*Pool, error) {
	poolDetail := &Pool{}
	body, err := p.Client.Post("/card/pool/detail", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, poolDetail); err != nil {
		return nil, err
	}
	return poolDetail, nil
}

//用户下所有流量池信息
func (p * PoolService) List() ([]Pool, error) {
	poolDetailList := make([]Pool, 0)
	body, err := p.Client.Post("/card/pool/list", nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &poolDetailList); err != nil {
		return nil, err
	}
	return poolDetailList, nil
}