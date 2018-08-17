package simboss

import (
	"time"
	"strconv"
	"net/http"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/url"
)

const API_ROOT string = "https://api.simboss.com/2.0"

type Client struct {
	appId     string
	appSecret string
	User      UserService
	Device    DeviceService
	Pool      PoolService
	Sms       SmsService
	Realname  RealnameService
}

func NewClient(appId, appSecret string) *Client {
	c :=  &Client{
		appId:     appId,
		appSecret: appSecret,
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
	data.Set("timestamp", strconv.FormatInt(time.Now().UnixNano()/1e6, 10))
	data.Set("sign", c.sign(data))

	req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8;")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	businessResponse := Response{}

	if err := json.Unmarshal(body, &businessResponse); err != nil {
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
	Message string `json:"string"`
	Detail string `json:"detail"`
}

type ResponseError struct {
	Code string
	Message string
	Detail string
}

func (b ResponseError) Error() string {
	return fmt.Sprintf("Code: %s Message: %s Detail: %s", b.Code, b.Message, b.Detail)
}
