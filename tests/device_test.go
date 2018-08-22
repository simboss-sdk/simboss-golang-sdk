package tests

import (
	"testing"
	"net/http"
	"io/ioutil"
	"bytes"
	"github.com/simboss-sdk/simboss-golang-sdk/v2"
	"net/url"
	"strings"
)

func TestDevice_DetailBatch(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
  "message": "OK",
  "detail": "normally invoke",
  "data": [
    {
      "iccid": "898602B7091701051989",
      "imsi": "460040783801989",
      "msisdn": "1064878381989",
      "carrier": "cmcc",
      "type": "SINGLE",
      "status": "deactivation",
      "deviceStatus": "DEACTIVATED_NAME",
      "openDate": "2018-03-14 14:38:31",
      "startDate": "2018-03-14 14:38:56",
      "expireDate": "2018-03-31 23:59:59",
      "lastSyncDate": "2018-03-14 16:00:01",
      "activeDuration": 0,
      "realnameRequired": true,
      "testingExpireDate": "2018-03-14 14:38:31",
      "ratePlanId": 1982,
      "iratePlanName": "30.0M/月",
      "usedDataVolume": 0.0,
      "totalDataVolume": 30.0,
      "ratePlanEffetiveDate": "2018-03-14 14:38:56",
      "ratePlanExpirationDate": "2018-03-31 23:59:59",
      "dataUsage": 0.0,
      "realName": "张三",
      "realNameCertifystatus": "not-submit", 
      "remark":"根据 xxxxxxxxxx 订单号出库",
      "memo":"xxx",
      "functions": ["FUN1", "FUN2"]
    },
    {
      "iccid": "898607B2091700080201",
      "imsi": "460041201100201",
      "msisdn": "1064720110201",
      "carrier": "cmcc",
      "type": "POOL",
      "status": "activation",
      "deviceStatus": "DEACTIVATED_NAME",
      "openDate": "2018-03-13 19:46:54",
      "startDate": "2018-03-13 19:52:42",
      "expireDate": "2018-05-31 23:59:59",
      "lastSyncDate": "2018-03-13 22:55:37",
      "activeDuration": 0,
      "realnameRequired": false,
      "cardPoolId": 898,
      "testingExpireDate": "2018-03-13 19:46:54",
      "ratePlanId": 2032,
      "iratePlanName": "10.0M/月",
      "usedDataVolume": 149.900390625,
      "totalDataVolume": 10.0,
      "ratePlanEffetiveDate": "2018-03-13 19:52:41",
      "ratePlanExpirationDate": "2018-03-31 23:59:59",
      "dataUsage": 149.900390625,
      "realName": "张三",
      "realNameCertifystatus": "not-submit", 
      "remark":"根据 xxxxxxxxxx 订单号出库",
      "memo":"xxx",
      "functions": ["FUN1", "FUN2"]
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

	params := url.Values{}

	_, err := client.Device.DetailBatch(params)
	if err != simboss.ErrRequiredBatchCardId {
		t.Fatal("error should be ErrRequiredBatchCardId")
	}

	params.Set("iccids", "898606182182822223,898606182111132824")
	deviceList, err := client.Device.DetailBatch(params)
	if err != nil {
		t.Fatal(err)
	}
	if len(deviceList) != 2 {
		t.Error("length should be 2")
	}
	if deviceList[0].Iccid != "898602B7091701051989" {
		t.Error("deviceList[0].Iccid should be 898602B7091701051989")
	}
	if deviceList[0].OpenDate.String() != "2018-03-14 14:38:31" {
		t.Error("deviceList[0].OpenDate should be 2018-03-14 14:38:31")
	}
}

func TestDevice_DetailBatch_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	params := url.Values{}
	params.Set("iccids", strings.Join(iccids, ","))
	details, err := client.Device.DetailBatch(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", details)
}

func TestDevice_Detail(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
  "code": "0",
  "message": "OK",
  "detail": "",
  "data": {
    "iccid": "898607B2091700080201",
    "imsi": "460041201100201",
    "msisdn": "1064720110201",
    "carrier": "cmcc",
    "type": "POOL",
    "status": "activation",
    "deviceStatus": "DEACTIVATED_NAME",
    "openDate": "2018-03-13 19:46:54",
    "startDate": "2018-03-13 19:52:42",
    "expireDate": "2018-05-31 23:59:59",
    "lastSyncDate": "2018-03-13 22:55:37",
    "activeDuration": 0,
    "realnameRequired": false,
    "cardPoolId": 898,
    "testingExpireDate": "2018-03-13 19:46:54",
    "ratePlanId": 2032,
    "iratePlanName": "10.0M/月",
    "usedDataVolume": 149.900390625,
    "totalDataVolume": 10.0,
    "ratePlanEffetiveDate": "2018-03-13 19:52:41",
    "ratePlanExpirationDate": "2018-03-31 23:59:59",
    "dataUsage": 149.900390625,
    "realName": "张三",
    "realNameCertifystatus": "not-submit",
    "remark": "根据 xxxxxxxxxx 订单号出库",
    "memo": "xxx",
    "functions": [
      "FUN1",
      "FUN2"
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

	_, err := client.Device.Detail(params)
	if err != simboss.ErrRequiredCardId {
		t.Fatal("error should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")

	device, err := client.Device.Detail(params)
	if err != nil {
		t.Fatal(err)
	}
	if device.Iccid != "898607B2091700080201" {
		t.Error("iccid should be 898607B2091700080201")
	}
	if device.OpenDate.String() != "2018-03-13 19:46:54" {
		t.Error("OpenDate should be 2018-03-13 19:46:54")
	}
}

func TestDevice_Detail_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	params := url.Values{}
	params.Set("iccid", poolIccid)
	detail, err := client.Device.Detail(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", detail)
}

func TestDevice_OrderPlans(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail": "",
    "data":    [
      {
          "ratePlanId": 64, 
          "price": 100,  
          "timeLength": 3,  
          "timeUnit": "month",
          "name": "1季度300M-100元",
          "type": "monthly", 
          "description": "1季度300M每月不清零套餐",
          "dataVolume": 300,
           "unlimitedVolume": true,
          "effectiveDate": "2017-06-12 00:12:34",
          "expirationDate": "2017-07-12 00:12:34",
          "status" : "active"
      }
  ]
}`
		resp := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}

		return &resp, nil
	}

	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)

	params := url.Values{}

	_, err := client.Device.OrderedPlans(params)
	if err != simboss.ErrRequiredCardId {
		t.Fatal("error should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")
	orderedPlanList, err := client.Device.OrderedPlans(params)
	if err != nil {
		t.Fatal(err)
	}
	if len(orderedPlanList) != 1 {
		t.Fatal("length should be 1")
	}
	if orderedPlanList[0].RatePlanId != 64 {
		t.Fatal("orderedPlanList[0].RatePlanId should be 64")
	}
	if orderedPlanList[0].EffectiveDate.String() != "2017-06-12 00:12:34" {
		t.Fatal("orderedPlanList[0].EffectiveDate should be 2017-06-12 00:12:34")
	}
}

func TestDevice_OrderPlans_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	params := url.Values{}
	params.Set("iccid", singerIccid)
	plans, err := client.Device.OrderedPlans(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", plans)
}

func TestDevice_Rateplans(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code": "0",
    "message": "OK",
    "detail": "",
    "data":    [
      {
          "ratePlanId": 0, 
          "price": 100,  
          "timeLength": 3,  
          "timeUnit": "month",
          "name": "1季度300M-100元",
          "type": "monthly", 
          "description": "1季度300M每月不清零套餐",
          "dataVolume": 1024,
          "unlimitedVolume": true 
      }
  ]
}`
		resp := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}

		return &resp, nil
	}

	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)

	params := url.Values{}

	_, err := client.Device.Rateplans(params)
	if err != simboss.ErrRequiredCardId {
		t.Fatal("error should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")
	rateplansList, err := client.Device.Rateplans(params)
	if err != nil {
		t.Fatal(err)
	}
	if len(rateplansList) != 1 {
		t.Fatal("length should be 1")
	}
	if rateplansList[0].RatePlanId != 0 {
		t.Fatal("orderedPlanList[0].RatePlanId should be 0")
	}
}

func TestDevice_Rateplans_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	params := url.Values{}
	params.Set("iccid", singerIccid)
	plans, err := client.Device.Rateplans(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", plans)
}

func TestDevice_Recharge(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
  "code": "0",
  "message": "OK",
  "detail": "",
  "data": "20170426733333231"
}`
		resp := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}

		return &resp, nil
	}


	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)

	params := url.Values{}

	_, err := client.Device.Recharge(params)
	if err != simboss.ErrRequiredCardId {
		t.Fatal("error should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")
	id, err := client.Device.Recharge(params)
	if err != nil {
		t.Fatal(err)
	}
	if id != "20170426733333231" {
		t.Error("id should be 20170426733333231")
	}
}

func TestDevice_RechargeRecords(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail":"",
    "data": [
        {
            "sequence":"2016033031121471",  
            "orderName":"1 季度 300M-100 元",  
            "orderType":"open",
            "price": 300, 
            "createTime":"2015-11-09 11:12:13",  
            "unlimitedVolume": true,
            "period": 12,
            "timeLenth": 12,
            "timeUnit": "month",
            "dataVolume": 12,
            "volumePlanType": "monthly",
            "volumePlanName": "30M/个月"
        }
    ]
}`
		resp := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}

		return &resp, nil
	}


	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)

	params := url.Values{}

	_, err := client.Device.RechargeRecords(params)
	if err != simboss.ErrRequiredCardId {
		t.Fatal("error should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")
	records, err := client.Device.RechargeRecords(params)
	if err != nil {
		t.Fatal(err)
	}

	if len(records) != 1 {
		t.Error("length should be 1")
	}
	if records[0].Sequence != "2016033031121471" {
		t.Error("Sequence should be 2016033031121471")
	}
	if records[0].UnlimitedVolume != true {
		t.Error("UnlimitedVolume should be true")
	}
}

func TestDevice_RechargeRecords_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	params := url.Values{}
	params.Set("iccid", singerIccid)
	result, err := client.Device.RechargeRecords(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}

func TestDevice_GprsStatus(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail":"",
    "data" : {
        "iccid": "898606111111132823",
        "ipAddr": "10.172.27.33",
        "status": "ACTIVATED_NAME",
        "apn": "cmiot"
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

	_, err := client.Device.GprsStatus(params)
	if err != simboss.ErrRequiredCardId {
		t.Fatal("error should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")
	status, err := client.Device.GprsStatus(params)
	if err != nil {
		t.Fatal(err)
	}
	if status.Iccid != "898606111111132823" {
		t.Error("Iccid should be 898606111111132823")
	}
}

func TestDevice_GprsStatus_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	params := url.Values{}
	params.Set("iccid", singerIccid)
	result, err := client.Device.GprsStatus(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}

func TestDevice_UserStatus(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail":"",
    "data": {
        "iccid": "898606111111132823",
        "status": "ACTIVATED_NAME" 
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

	_, err := client.Device.UserStatus(params)
	if err != simboss.ErrRequiredCardId {
		t.Fatal("error should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")
	status, err := client.Device.UserStatus(params)
	if err != nil {
		t.Fatal(err)
	}
	if status.Iccid != "898606111111132823" {
		t.Error("Iccid should be 898606111111132823")
	}
}

func TestDevice_UserStatus_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	params := url.Values{}
	params.Set("iccid", singerIccid)
	result, err := client.Device.UserStatus(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}

func TestDevice_RunningStatus(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail":"",
    "data" : {
        "iccid": "898606111111132823",
        "status": "ACTIVATED_NAME" 
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

	_, err := client.Device.RunningStatus(params)
	if err != simboss.ErrRequiredCardId {
		t.Fatal("error should be ErrRequiredCardId")
	}

	params.Set("iccid", "898606111111132823")
	status, err := client.Device.RunningStatus(params)
	if err != nil {
		t.Fatal(err)
	}

	if status.Iccid != "898606111111132823" {
		t.Error("Iccid should be 898606111111132823")
	}
}

func TestDevice_RunningStatus_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	params := url.Values{}
	params.Set("iccid", poolIccid)
	result, err := client.Device.RunningStatus(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}

func TestDevice_RatePlanSummary(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail":"",
    "data" : {
        "iccid": "898606111111132823",
        "volumeSummary": 323, 
        "expirationDate": "2017-08-31 23:59:59" 
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

	_, err := client.Device.RatePlanSummary(params)
	if err != simboss.ErrRequiredCardId {
		t.Fatal("error should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")
	summary, err := client.Device.RatePlanSummary(params)
	if err != nil {
		t.Fatal(err)
	}
	if summary.Iccid != "898606111111132823" {
		t.Error("Iccid should be 898606111111132823")
	}
	if summary.ExpirationDate.String() != "2017-08-31 23:59:59" {
		t.Error("ExpirationDate should be 2017-08-31 23:59:59")
	}
}

func TestDevice_RatePlanSummary_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	params := url.Values{}
	params.Set("iccid", singerIccid)
	result, err := client.Device.RatePlanSummary(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}

func TestDevice_ModifyDeviceStatus(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail":"",
    "data" : "success"
}`
		resp := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}

		return &resp, nil
	}


	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)

	params := url.Values{}

	err := client.Device.ModifyDeviceStatus(params)
	if err != simboss.ErrRequiredCardId {
		t.Fatal("error should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")
	err = client.Device.ModifyDeviceStatus(params)
	if err != simboss.ErrRequired {
		t.Fatal("error should be ErrRequired")
	}

	params.Set("status", "ACTIVATED_NAME")
	err = client.Device.ModifyDeviceStatus(params)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDevice_DailyUsage(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail":"",
    "data" : {
        "usage": 12.33,
        "date": "2017-11-09"
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

	_, err := client.Device.DailyUsage(params)
	if err != simboss.ErrRequiredCardId {
		t.Fatal("error should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")
	_, err = client.Device.DailyUsageByDateRange(params)
	if err != simboss.ErrRequired {
		t.Fatal("error should be ErrRequired")
	}

	params.Set("date", "2017-11-09")
	dailyUsage, err := client.Device.DailyUsage(params)
	if err != nil {
		t.Fatal(err)
	}
	if dailyUsage.Usage != 12.33 {
		t.Error("dailyUsage.Usage should be 12.33")
	}
	if dailyUsage.Date.String() != "2017-11-09" {
		t.Error("dailyUsage.Date should be 2017-11-09")
	}
}

func TestDevice_DailyUsage_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	params := url.Values{}
	params.Set("iccid", singerIccid)
	params.Set("date", "2017-08-01")
	result, err := client.Device.DailyUsage(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}

func TestDevice_DailyUsageByDateRange(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail":"",
    "data" : [
        {
            "usage": 12.33,
            "date": "2017-11-08"
        },
        {
            "usage": 13.33,
            "date": "2017-11-09"
        }
     ]
}`
		resp := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}

		return &resp, nil
	}


	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)

	params := url.Values{}

	_, err := client.Device.DailyUsageByDateRange(params)
	if err != simboss.ErrRequiredCardId {
		t.Fatal("error should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")
	_, err = client.Device.DailyUsageByDateRange(params)
	if err != simboss.ErrRequired {
		t.Fatal("error should be ErrRequired")
	}

	params.Set("startDate", "2017-11-08")
	params.Set("endDate", "2017-11-09")
	dailyUsageList, err := client.Device.DailyUsageByDateRange(params)
	if err != nil {
		t.Fatal(err)
	}
	if len(dailyUsageList) != 2 {
		t.Error("length should be 2")
	}
	if dailyUsageList[0].Usage != 12.33 {
		t.Error("dailyUsageList[0].Usage should be 12.33")
	}
	if dailyUsageList[0].Date.String() != "2017-11-08" {
		t.Error("dailyUsageList[0].Date should be 2017-11-08")
	}
}

func TestDevice_DailyUsageByDateRange_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	params := url.Values{}
	params.Set("iccid", singerIccid)
	params.Set("startDate", "2018-07-01")
	params.Set("endDate", "2018-08-01")
	result, err := client.Device.DailyUsageByDateRange(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", result)
}

func TestDevice_CancelTesting(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail":"",
    "data": "success"
}`
		resp := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}

		return &resp, nil
	}


	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)

	params := url.Values{}

	err := client.Device.CancelTesting(params)
	if err != simboss.ErrRequiredCardId {
		t.Fatal("error should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")
	err = client.Device.CancelTesting(params)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDevice_MemoUpdate(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail":"",
    "data": "success"
}`
		resp := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}

		return &resp, nil
	}

	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)

	params := url.Values{}

	err := client.Device.MemoUpdate(params)
	if err != simboss.ErrRequiredCardId {
		t.Fatal("error should be ErrRequiredCardId")
	}

	params.Set("iccid", "1001")
	err = client.Device.MemoUpdate(params)
	if err != simboss.ErrRequired {
		t.Fatal("error should be ErrRequired")
	}

	params.Set("memo", "hello world")
	err = client.Device.MemoUpdate(params)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDevice_MemoUpdate_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	params := url.Values{}
	params.Set("iccid", singerIccid)
	params.Set("memo", "ych sdk:MemoUpdate test")
	err := client.Device.MemoUpdate(params)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDevice_MemoBatchUpdate(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `
{
    "code":"0",
    "message": "OK",
    "detail":"",
    "data": "success"
}`
		resp := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}

		return &resp, nil
	}


	client := simboss.NewClientWithHttpDo(appId, appSecret, httpDo)

	params := url.Values{}

	err := client.Device.MemoBatchUpdate(params)
	if err != simboss.ErrRequiredBatchCardId {
		t.Fatal("error should be ErrRequiredBatchCardId")
	}

	params.Set("iccids", "898606182182822223,898606182111132824")
	err = client.Device.MemoBatchUpdate(params)
	if err != simboss.ErrRequired {
		t.Fatal("error should be ErrRequired")
	}

	params.Set("memo", "hello world")
	err = client.Device.MemoBatchUpdate(params)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDevice_MemoBatchUpdate_Response(t *testing.T) {
	client := simboss.NewClient(appId, appSecret)
	params := url.Values{}
	params.Set("iccids", strings.Join(iccids, ","))
	params.Set("memo", "ych sdk:MemoBatchUpdate test")
	err := client.Device.MemoBatchUpdate(params)
	if err != nil {
		t.Fatal(err)
	}
}