package simboss

import "net/url"

type RealnameService struct {
	client *Client
}

func (r *RealnameService) Submit(params url.Values) error {
	_, err := r.client.Post("/realname/submitRealname", params)
	if err != nil {
		return err
	}
	return nil
}