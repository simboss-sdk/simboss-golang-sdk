package simboss

import (
	"testing"
	"strconv"
	"time"
)

func TestClient_sign(t *testing.T) {
	const appId string = "1111"
	const  appSecret = "2222"
	client := NewClient(appId, appSecret)
	data := map[string]string{
		"iccid":     "1001",
		"type":      "cmcc",
		"appid":     appId,
		"timestamp": strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
	}
	sign := client.sign(data)
	t.Log(sign)
}