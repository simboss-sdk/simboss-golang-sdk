package tests

import (
	"testing"
	"net/http"
	"io/ioutil"
	"bytes"
	"github.com/simboss-sdk/simboss-golang-sdk/v2"
	"net/url"
)

func TestRealname_Sumbit(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail":"",
    "data" : ""
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
	params.Set("name", "xxxxxx")
	params.Set("licenseType", "xxxxxx")
	params.Set("licenseCode", "xxxxxx")
	params.Set("phone", "xxxxxx")
	params.Set("pic1", "xxxxxx")
	params.Set("pic2", "xxxxxx")

	err := client.Realname.Submit(params)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRealanme_Params_Error(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail":"",
    "data" : ""
}`
		resp := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}

		return &resp, nil
	}

	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)

	params := url.Values{}

	err := client.Realname.Submit(params)
	if err == nil {
		t.Fatal("should throw a error")
	}
	if err != simboss.ErrRequiredCardId {
		t.Fatal("should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")
	err = client.Realname.Submit(params)
	if err == nil {
		t.Fatal("should throw a error")
	}
	if err != simboss.ErrRequired {
		t.Fatal("should be ErrRequired")
	}

}