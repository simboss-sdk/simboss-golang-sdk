package simboss

import "net/url"

type RealnameService struct {
	Client *Client
}

func (r *RealnameService) Submit(params url.Values) error {
	_, err := r.Client.Post("/realname/submitRealname", params)
	if err != nil {
		return err
	}
	return nil
}