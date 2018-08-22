package tests

import (
	"testing"
	"net/http"
	"io/ioutil"
	"bytes"
	"github.com/simboss-sdk/simboss-golang-sdk/v2"
)

func TestUser_DashboardGet(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `{"message":"OK","detail":"normally invoke","data":{"userAccount":"13800138000","company":"上海某某有限公司","name":"张三","balance":100.3,"mobile":"13800138000","registerTime":"2016-10-20 23:45:55","cardSummary":{"cmcc":0,"unicom":0,"telcom":12060}},"code":"0","success":true}`
		resp := http.Response{
			Body:       ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}
		return &resp, nil
	}

	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)

	dashboard, err := client.User.DashboardGet()
	if err != nil {
		t.Fatal(err)
	}

	if dashboard.UserAccount != "13800138000" {
		t.Errorf("expected 13800138000, but got %s", "13800138000")
	}
}

func TestUser_DashboardGet_Response_Error(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"610",
    "message": "server error",
    "detail":""
}`
		resp := http.Response{
			Body:       ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}
		return &resp, nil
	}

	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)

	dashboard, err := client.User.DashboardGet()
	if err == nil {
		t.Fatal("should throw a error")
	}

	responseError, ok := err.(simboss.ResponseError)
	if !ok {
		t.Error("should be a response error")
	}

	if responseError.Code != "610" || responseError.Message != "server error" {
		t.Error("code should equal '610'; messge should eauql to 'server error'")
	}

	if dashboard != nil {
		t.Error("dashboard should be nil")
	}
}