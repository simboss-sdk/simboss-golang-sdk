package simboss

import (
	"net/url"
	"github.com/simboss-sdk/simboss-golang-sdk/utils"
	)


type RealnameService struct {
	client *Client
}

func (r *RealnameService) Submit(params url.Values) error {
	if err := RequiredCardId(params); err != nil {
		return err
	}
	if !utils.Required(params,"name","licenseType","licenseCode","phone","pic1","pic2") {
		return ErrRequired
	}
	_, err := r.client.Post("/realname/submitRealname", params)
	if err != nil {
		return err
	}
	return nil
}