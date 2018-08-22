package simboss

import (
	"testing"
			"net/url"
	"encoding/json"
	"github.com/simboss-sdk/simboss-golang-sdk/utils"
)

func TestClient_sign(t *testing.T) {
	const appId string = "1001"
	const  appSecret = "xxxx"
	client := NewClient(appId, appSecret)
	data := url.Values{}
	data.Set("iccid", "1001")
	data.Set("type", "cmcc")
	data.Set("appid", appId)
	data.Set("timestamp", utils.GetNonce())
	sign := client.sign(data)
	t.Log(sign)
}

func TestResponse_Unmarshal(t *testing.T) {
	resp := Response{
		Code:    "0",
		Data:    Pool{Id: 1001},
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}

	nResp := Response{}

	if err := json.Unmarshal(respBytes, &nResp); err != nil {
		t.Fatal(err)
	}

	dataBytes, err := json.Marshal(nResp.Data)
	if err != nil {
		t.Fatal(err)
	}

	detail := Pool{}

	if err := json.Unmarshal(dataBytes, &detail); err != nil {
		t.Fatal(err)
	}

	if detail.Id != 1001 {
		t.Error("id should be 1001")
	}
}