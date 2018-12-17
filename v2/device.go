package simboss

import (
	"net/url"
	"github.com/simboss-sdk/simboss-golang-sdk/utils/time"
	"encoding/json"
	"github.com/simboss-sdk/simboss-golang-sdk/utils"
	)

type DeviceService struct {
	client *Client
}

type Device struct {
	Iccid string `json:"iccid"`
	Imsi string `json:"imsi"`
	Msisdn string `json:"msisdn"`
	Carrier string `json:"carrier"`
	Type string `json:"type"`
	Status string `json:"status"`
	DeviceStatus string `json:"deviceStatus"`
	OpenDate time.Time `json:"openDate"`
	StartDate string `json:"startDate"`
	ExpireDate string `json:"expireDate"`
	LastSyncDate string `json:"lastSyncDate"`
	ActiveDuration int `json:"activeDuration"`
	RealnameRequired bool `json:"realnameRequired"`
	CardPoolId int `json:"cardPoolId"`
	TestingExpireDate time.Time `json:"testingExpireDate"`
	RatePlanId int `json:"ratePlanId"`
	IratePlanName string `json:"iratePlanName"`
	UsedDataVolume float64 `json:"usedDataVolume"`
	TotalDataVolume float64 `json:"totalDataVolume"`
	RatePlanEffetiveDate time.Time `json:"ratePlanEffetiveDate"`
	RatePlanExpirationDate time.Time `json:"ratePlanExpirationDate"`
	DataUsage float64 `json:"dataUsage"`
	RealName string `json:"realName"`
	RealNameCertifystatus string `json:"realNameCertifystatus"`
	Remark string `json:"remark"`
	Memo string `json:"memo"`
	Functions []string `json:"functions"`
}

// 批量卡详情
func (d * DeviceService) DetailBatch(params url.Values) ([]Device, error) {
	if err := RequiredBatchCardId(params); err != nil {
		return nil, err
	}
	deviceList := make([]Device, 0)
	body, err := d.client.Post("/device/detail/batch", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &deviceList); err != nil {
		return nil, err
	}
	return deviceList, nil
}

// 单卡详情
func (d *DeviceService) Detail(params url.Values) (*Device, error) {
	if err := RequiredCardId(params); err != nil {
		return nil, err
	}
	device := &Device{}
	body, err := d.client.Post("/device/detail", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, device); err != nil {
		return nil, err
	}
	return device, nil
}

type OrderedPlan struct {
	RatePlanId int `json:"ratePlanId"`
	Price float64 `json:"price"`
	TimeLength int `json:"timeLength"`
	TimeUnit string `json:"timeUnit"`
	Name string `json:"name"`
	Type string `json:"type"`
	Description string `json:"description"`
	DataVolume float64 `json:"dataVolume"`
	UnlimitedVolume bool `json:"unlimitedVolume"`
	EffectiveDate time.Time `json:"effectiveDate"`
	ExpirationDate time.Time `json:"expirationDate"`
	Status string `json:"status"`
}

// 单卡已订购套餐列表
func (d *DeviceService) OrderedPlans(params url.Values) ([]OrderedPlan, error) {
	if err := RequiredCardId(params); err != nil {
		return nil, err
	}
	orderedPlanList := make([]OrderedPlan, 0)
	body, err := d.client.Post("/device/orderedPlans", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &orderedPlanList); err != nil {
		return nil, err
	}
	return orderedPlanList, nil
}

type RatePlan struct {
	RatePlanId int `json:"ratePlanId"`
	Price float64 `json:"price"`
	TimeLength int `json:"timeLength"`
	TimeUnit string `json:"timeUnit"`
	Name string `json:"name"`
	Type string `json:"type"`
	Description string `json:"description"`
	DataVolume float64 `json:"dataVolume"`
	UnlimitedVolume bool `json:"unlimitedVolume"`
}

// 单卡可续费套餐信息
func (d *DeviceService) Rateplans(params url.Values) ([]RatePlan, error) {
	if err := RequiredCardId(params); err != nil {
		return nil, err
	}
	ratePlanList := make([]RatePlan, 0)
	body, err := d.client.Post("/device/rateplans", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &ratePlanList); err != nil {
		return nil, err
	}
	return ratePlanList, nil
}

// 单卡续费
func (d *DeviceService) Recharge(params url.Values) (string, error) {
	if err := RequiredCardId(params); err != nil {
		return "", err
	}
	cashFlowUuid := ""
	body, err := d.client.Post("/device/recharge", params)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(body, &cashFlowUuid); err != nil {
		return "", err
	}
	return cashFlowUuid, nil
}

type RechargeRecord struct {
	Sequence string `json:"sequence"`
	OrderName string `json:"orderName"`
	OrderType string `json:"orderType"`
	Price float64 `json:"price"`
	CreateTime time.Time `json:"createTime"`
	UnlimitedVolume bool `json:"unlimitedVolume"`
	Period int `json:"period"`
	TimeLenth int `json:"timeLenth"`
	TimeUnit string `json:"timeUnit"`
	DataVolume float64 `json:"dataVolume"`
	VolumePlanType string `json:"volumePlanType"`
	VolumePlanName string `json:"volumePlanName"`
}

// 单卡续费记录
func (d *DeviceService) RechargeRecords(params url.Values) ([]RechargeRecord, error) {
	if err := RequiredCardId(params); err != nil {
		return nil, err
	}
	rechargeRecordList := make([]RechargeRecord, 0)
	body, err := d.client.Post("/device/recharge/records", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &rechargeRecordList); err != nil {
		return nil, err
	}
	return rechargeRecordList, nil
}

type GprsStatus struct {
	Iccid string `json:"iccid"`
	IpAddr string `json:"ipAddr"`
	Status string `json:"status"`
	Apn string `json:"apn"`
}

// 实时连接状态查询
func (d *DeviceService) GprsStatus(params url.Values) (*GprsStatus, error) {
	if err := RequiredCardId(params); err != nil {
		return nil, err
	}
	gprsStatus := &GprsStatus{}
	body, err := d.client.Post("/device/gprsStatus", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, gprsStatus); err != nil {
		return nil, err
	}
	return gprsStatus, nil
}

type UserStatus struct {
	Iccid string `json:"iccid"`
	Status string `json:"status"`
}

// 实时用户状态查询
func (d *DeviceService) UserStatus(params url.Values) (*UserStatus, error) {
	if err := RequiredCardId(params); err != nil {
		return nil, err
	}
	userStatus := &UserStatus{}
	body, err := d.client.Post("/device/userStatus", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, userStatus); err != nil {
		return nil, err
	}
	return userStatus, nil
}

type RunningStatus struct {
	Iccid string `json:"iccid"`
	Status string `json:"status"`
}

// 设备实时开关机状态查询
func (d *DeviceService) RunningStatus(params url.Values) (*RunningStatus, error) {
	if err := RequiredCardId(params); err != nil {
		return nil, err
	}
	runningStatus := &RunningStatus{}
	body, err := d.client.Post("/device/runningStatus", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, runningStatus); err != nil {
		return nil, err
	}
	return runningStatus, nil
}

type RatePlanSummary struct {
	Iccid string `json:"iccid"`
	VolumeSummary float64 `json:"volumeSummary"`
	ExpirationDate time.Time `json:"expirationDate"`
}

// 查询设备套餐概要
func (d *DeviceService) RatePlanSummary(params url.Values) (*RatePlanSummary, error) {
	if err := RequiredCardId(params); err != nil {
		return nil, err
	}
	ratePlanSummary := &RatePlanSummary{}
	body, err := d.client.Post("/device/ratePlan/summary", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, ratePlanSummary); err != nil {
		return nil, err
	}
	return ratePlanSummary, nil
}

// 流量池卡开关网络
func (d *DeviceService) ModifyDeviceStatus(params url.Values) (error) {
	if err := RequiredCardId(params); err != nil {
		return err
	}
	if !utils.Required(params, "status") {
		return ErrRequired
	}
	_, err := d.client.Post("/device/modifyDeviceStatus", params)
	if err != nil {
		return err
	}
	return nil
}

type DailyUsage struct {
	Usage float64 `json:"usage"`
	Date time.Date `json:"date"`
}

// 日用量查询
func (d *DeviceService) DailyUsage(params url.Values) (*DailyUsage, error) {
	if err := RequiredCardId(params); err != nil {
		return nil, err
	}
	if !utils.Required(params, "date") {
		return nil, ErrRequired
	}
	dailyUsage := &DailyUsage{}
	body, err := d.client.Post("/device/dailyUsage", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, dailyUsage); err != nil {
		return nil, err
	}
	return dailyUsage, nil
}

// 日用量按照时间范围查询
func (d *DeviceService) DailyUsageByDateRange(params url.Values) ([]DailyUsage, error) {
	if err := RequiredCardId(params); err != nil {
		return nil, err
	}
	if !utils.Required(params, "startDate", "endDate") {
		return nil, ErrRequired
	}
	dailyUsageList := make([]DailyUsage, 0)
	body, err := d.client.Post("/device/dailyUsageByDateRange", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &dailyUsageList); err != nil {
		return nil, err
	}
	return dailyUsageList, nil
}

type DailyUsageWithCardId struct {
	Usage float64 `json:"usage"`
	Iccid string `json:"iccid"`
	Imsi string `json:"imsi"`
	Msisdn string `json:"msisdn"`
}

type DailyUsageBatch struct {
	Date string `json:"date"`
	DailyUsageList []DailyUsageWithCardId `json:"dailyUsageList"`
}

// 批量查询日用量查询
func (d *DeviceService) DailyUsageBatch(params url.Values) (*DailyUsageBatch, error) {
	if err := RequiredBatchCardId(params); err != nil {
		return nil, err
	}
	if !utils.Required(params, "date") {
		return nil, ErrRequired
	}
	dailyUsageList := &DailyUsageBatch{
		Date: "",
		DailyUsageList: make([]DailyUsageWithCardId, 0),
	}
	body, err := d.client.Post("/device/dailyUsage/batch", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &dailyUsageList); err != nil {
		return nil, err
	}
	return dailyUsageList, nil
}

// 取消测试期
func (d *DeviceService) CancelTesting(params url.Values) (error) {
	if err := RequiredCardId(params); err != nil {
		return err
	}
	_, err := d.client.Post("/device/cancelTesting", params)
	if err != nil {
		return err
	}
	return nil
}

// 更新备注
func (d *DeviceService) MemoUpdate(params url.Values) (error) {
	if err := RequiredCardId(params); err != nil {
		return err
	}
	if !utils.Required(params, "memo") {
		return ErrRequired
	}
	_, err := d.client.Post("/device/memo/update", params)
	if err != nil {
		return err
	}
	return nil
}

// 批量更新备注
func (d *DeviceService) MemoBatchUpdate(params url.Values) (error) {
	if err := RequiredBatchCardId(params); err != nil {
		return err
	}
	if !utils.Required(params, "memo") {
		return ErrRequired
	}
	_, err := d.client.Post("/device/memo/batchUpdate", params)
	if err != nil {
		return err
	}
	return nil
}

type IccidList struct {
	Page Page `json:"page"`
	List []string `json:"list"`
}

// iccid列表查询
func (d *DeviceService) IccidList(params url.Values) (*IccidList, error) {
	if !utils.Required(params,"pageNo") {
		return nil, ErrRequired
	}
	iccidList := &IccidList{
		Page: Page{},
		List: make([]string, 0),
	}
	body, err := d.client.Post("/device/iccid/list", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, iccidList); err != nil {
		return nil, err
	}
	return iccidList, nil
}

// 强制临时激活
func (d *DeviceService) Activate(params url.Values) error {
	if err := RequiredCardId(params); err != nil {
		return err
	}
	_, err := d.client.Post("/device/activate", params)
	if err != nil {
		return err
	}
	return nil
}