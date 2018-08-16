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

func (c *Client) sign(data map[string]string) string {
	params := make([]string, 0)

	for key, value := range data {
		params = append(params, fmt.Sprintf("%s=%s", key, value))
	}

	sort.Strings(params)

	params = append(params, c.appSecret)
	paramsString := strings.Join(params, "&")

	h := sha256.New()
	h.Write([]byte(paramsString))

	return hex.EncodeToString(h.Sum(nil))
}

func (c *Client) Post(url string, data map[string]string) ([]byte, error) {
	url = API_ROOT + url

	if data == nil {
		data = make(map[string]string)
	}

	data["appid"] = c.appId
	data["timestamp"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	data["sign"] = c.sign(data)

	req, err := http.NewRequest("POST", url, nil)
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

	return body, nil
}
