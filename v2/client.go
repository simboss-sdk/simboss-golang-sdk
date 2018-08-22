package simboss

import (
			"net/http"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"github.com/simboss-sdk/simboss-golang-sdk/utils"
	"errors"
)

const API_ROOT string = "https://api.simboss.com/2.0"

var ErrRequired = errors.New("缺少必填参数")
var ErrRequiredCardId = errors.New("iccid, imsi, msisdn 三者选一")
var ErrRequiredBatchCardId = errors.New("iccids, imsis, msisdns 三者选一")

func RequiredCardId(values url.Values) error {
	if !utils.RequiredAtLeastOne(values, "iccid", "imsi", "msisdn") {
		return ErrRequiredCardId
	}
	return nil
}

func RequiredBatchCardId(values url.Values) error {
	if !utils.RequiredAtLeastOne(values, "iccids", "imsis", "msisdns") {
		return ErrRequiredBatchCardId
	}
	return nil
}

type Client struct {
	appId     string
	appSecret string
	httpDo func(*http.Client, *http.Request) (*http.Response, error)
	User      UserService
	Device    DeviceService
	Pool      PoolService
	Sms       SmsService
	Realname  RealnameService
}

func NewClient(appId, appSecret string) *Client {
	httpDo := func(c *http.Client, req *http.Request) (*http.Response, error) {
		return c.Do(req)
	}
	return NewClientWithHttpDo(appId, appSecret, httpDo)
}

func NewClientWithHttpDo(appId, appSecret string, httpDo func(*http.Client, *http.Request) (*http.Response, error)) *Client {
	c :=  &Client{
		appId:     appId,
		appSecret: appSecret,
		httpDo: httpDo,
	}
	c.User = UserService{c}
	c.Device = DeviceService{c}
	c.Pool = PoolService{c}
	c.Sms = SmsService{c}
	c.Realname = RealnameService{c}
	return c
}

func (c *Client) sign(data url.Values) string {
	params := make([]string, 0)

	for key, value := range data {
		params = append(params, fmt.Sprintf("%s=%s", key, strings.Join(value, ",")))
	}

	sort.Strings(params)

	params = append(params, c.appSecret)
	paramsString := strings.Join(params, "&")

	h := sha256.New()
	h.Write([]byte(paramsString))

	return hex.EncodeToString(h.Sum(nil))
}

func (c *Client) Post(path string, data url.Values) ([]byte, error) {
	path = API_ROOT + path

	if data == nil {
		data = make(url.Values)
	}

	data.Set("appid", c.appId)
	data.Set("timestamp", utils.GetNonce())
	data.Set("sign", c.sign(data))

	req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8;")

	res, err := c.httpDo(http.DefaultClient, req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	businessResponse := Response{}

	if err := json.Unmarshal(body, &businessResponse); err != nil {
		fmt.Println(body)
		return nil, err
	}

	if businessResponse.Code != "0" {
		return nil, ResponseError{businessResponse.Code, businessResponse.Message, businessResponse.Detail}
	}

	dataBytes, err := json.Marshal(businessResponse.Data)
	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}

type Response struct {
	Code string `json:"code"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
	Detail string `json:"detail"`
	Success bool `json:"success"`
}

type ResponseError struct {
	Code string
	Message string
	Detail string
}

func (b ResponseError) Error() string {
	return fmt.Sprintf("Code: %s Message: %s Detail: %s", b.Code, b.Message, b.Detail)
}
