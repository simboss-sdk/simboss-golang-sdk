package v2

import (
	"testing"
	"net/http"
	"github.com/simboss-sdk/simboss-golang-sdk/v2"
	"net/url"
	"io/ioutil"
	"bytes"
	)

func TestSms_Send(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail":""
}`
		resp := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}

		return &resp, nil
	}

	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)

	params := url.Values{}
	err := client.Sms.Send(params)
	if err != simboss.ErrRequiredCardId {
		t.Fatal("should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")
	params.Set("text", "sms content")

	err = client.Sms.Send(params)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSms_List(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail":"",
    "data": {
        "page": {
             "pageNo": 1,
             "pageSize": 50,
             "total": 102 
          },
          "list": [
              {
                   "id": 123,
                   "iccid": "898606111111132823",
                   "carrier": "cmcc",
                   "text": "status ok",
                   "sendTime": "2017-11-23 12:34:56",
                   "type": "send"
              }
          ]
    }
}`
		resp := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}

		return &resp, nil
	}

	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)

	params := url.Values{}
	params.Set("iccid", "1001")
	params.Set("pageNo", "1")

	list, err := client.Sms.List(params)
	if err != nil {
		t.Fatal(err)
	}

	if list.Page.Total != 102 {
		t.Error("list.Page.Total should be 102")
	}

	if len(list.List) != 1 {
		t.Fatal("list.List should exist one record")
	}

	if list.List[0].SendTime.String() != "2017-11-23 12:34:56" {
		t.Error("list.List[0].SendTime should be 2017-11-23 12:34:56")
	}
}

func TestSms_List_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	params := url.Values{}
	params.Set("pageNo", "1")
	params.Set("iccid", poolIccid)
	smsList, err := client.Sms.List(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", smsList)
}