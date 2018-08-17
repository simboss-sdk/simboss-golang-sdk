package utils

import "net/url"

func Required(params url.Values, names ...string) bool {
	var flag = true
	for _, name := range names {
		value := params.Get(name)
		if value == "" {
			flag = false
			break
		}
	}
	return flag
}
