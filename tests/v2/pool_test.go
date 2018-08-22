package v2

import (
	"testing"
	"net/http"
	"io/ioutil"
	"bytes"
	"github.com/simboss-sdk/simboss-golang-sdk/v2"
	"net/url"
)

func TestPool_Detail(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
  "message": "OK",
  "detail": "normally invoke",
  "data": {
    "id": 908,
    "poolSpecification": -1,
    "carrier": "cmcc",
    "totalVolume": 15.0,
    "useVolume": 3.264,
    "leftVolume": 11.736,
    "packageVolume": 0.0,
    "useRate": 0.2176,
    "totalCount": 1,
    "currentActivationCount": 1,
    "currentDeactivationCount": 0,
    "currentInventoryCount": 0,
    "currentTestingCount": 0,
    "currentRetiredCount": 0,
    "activationReadyCount": 0
  },
  "code": "0",
  "success": true
}`
		resp := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}

		return &resp, nil
	}

	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)

	params := url.Values{}

	pool, err := client.Pool.Detail(params)
	if err == nil {
		t.Fatal("should throw a error")
	}
	if err != simboss.ErrRequiredCardId {
		t.Fatal("should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")
	pool, err = client.Pool.Detail(params)
	if err != nil {
		t.Fatal(err)
	}

	if pool.Id != 908 {
		t.Error("id should be 908")
	}
}

func TestPool_List_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	pools, err := client.Pool.List()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", pools)
}

func TestPool_Detail_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	params := url.Values{}
	params.Set("iccid", poolIccid)
	pool, err := client.Pool.Detail(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", pool)
}

func TestPool_List(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
  "message": "OK",
  "detail": "normally invoke",
  "data": [
    {
      "id": 908,
      "poolSpecification": -1,
      "carrier": "cmcc",
      "totalVolume": 15.0,
      "useVolume": 3.264,
      "leftVolume": 11.736,
      "packageVolume": 0.0,
      "useRate": 0.2176,
      "totalCount": 1,
      "currentActivationCount": 1,
      "currentDeactivationCount": 0,
      "currentInventoryCount": 0,
      "currentTestingCount": 0,
      "currentRetiredCount": 0,
      "activationReadyCount": 0
    },
    {
      "id": 909,
      "poolSpecification": 30,
      "carrier": "chinanet",
      "totalVolume": 15.0,
      "useVolume": 3.264,
      "leftVolume": 11.736,
      "packageVolume": 0.0,
      "useRate": 0.2176,
      "totalCount": 1,
      "currentActivationCount": 1,
      "currentDeactivationCount": 0,
      "currentInventoryCount": 0,
      "currentTestingCount": 0,
      "currentRetiredCount": 0,
      "activationReadyCount": 0
    }
  ],
  "code": "0",
  "success": true
}`
		resp := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}

		return &resp, nil
	}

	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)
	poolList, err := client.Pool.List()
	if err != nil {
		t.Fatal(err)
	}

	if len(poolList) != 2 {
		t.Error("length of pooList should be 2")
	}

	if poolList[0].Id != 908 {
		t.Error("id should be 908")
	}
}
