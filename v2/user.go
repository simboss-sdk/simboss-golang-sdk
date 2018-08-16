package simboss

import "encoding/json"

type UserService struct {
	Client *Client
}

type CardSummary struct {
	Cmcc int `json:"cmcc"`
	Unicom int `json:"unicom"`
	Telcom int `json:"telcom"`
}

type Dashboard struct {
	UserAccount string `json:"userAccount"`
	Company string `json:"company"`
	Name string `json:"name"`
	Balance string `json:"balance"`
	Mobile string `json:"mobile"`
	RegisterTime string `json:"registerTime"`
	CardSummary string `json:"cardSummary"`
}


/*
 * 账户总览接口
 */
func (u * UserService) DashboardGet() (Dashboard, error) {
	dashboard := Dashboard{}
	body, err := u.Client.Post("/user/dashboard/get", nil)
	if err != nil {
		return dashboard, err
	}
	if err := json.Unmarshal(body, &dashboard); err != nil {
		return dashboard, err
	}
	return dashboard, nil
}