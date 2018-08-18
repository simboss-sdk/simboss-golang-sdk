package utils

import (
	"net/url"
	"time"
	"strconv"
)

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

func GetNonce() string {
	return strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
}