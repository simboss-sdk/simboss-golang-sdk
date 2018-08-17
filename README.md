## simboss-golang-sdk


#### Installation
go get github.com/simboss-sdk/simboss-golang-sdk

#### Usage
```go
package main

import (
	"fmt"
	"github.com/simboss-sdk/simboss-golang-sdk/v2"
)

func main() {
	client := simboss.NewClient(("api-key", "api-secret")
	info, err := client.User.DashboardGet()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(info)
	}
}
```

#### API
| API名称 | 请求参数 | 返回值 |    
| ------- | --------- | ---------:|
|1.1 账户总览接口           			|  无                                                                                                   |  (*Dashboard,error)           |  
|2.1 批量卡详情			        	|  url.Values{iccids,imsis,msisdns}                                                                     |  ([]Device, error))           |
|2.2 单卡详情				        	|  url.Values{iccid,imsi,msisdn}                                                                        |  (*Device, error)             |
|2.3 单卡已订购套餐列表	       		|  url.Values{iccid,imsi,msisdn}                                                                        |  ([]OrderedPlan, error)       |  
|2.4 单卡可续费套餐信息	        	|  url.Values{iccid,imsi,msisdn}                                                                        |  ([]RatePlan, error)          |
|2.5 单卡续费				        	|  url.Values{iccid,imsi,msisdn,ratePlanId,month,externalOrder}                                         |  (string, error)              |
|2.6 单卡续费记录			    		|  url.Values{iccid,imsi,msisdn}                                                                        |  ([]RechargeRecord, error)    |
|2.7 实时连接状态查询		        	|  url.Values{iccid,imsi,msisdn}                                                                        |  (*GprsStatus, error)         |
|2.8 实时用户状态查询		        	|  url.Values{iccid,imsi,msisdn}                                                                        |  (*UserStatus, error)         |
|2.9 设备实时开关机状态查询      		|  url.Values{iccid,imsi,msisdn}                                                                        |  (*RunningStatus, error)      |
|2.10 查询设备套餐概要        			|  url.Values{iccid,imsi,msisdn}                                                                        |  (*RatePlanSummary, error)    |
|2.11 流量池卡开关网络     			|  url.Values{iccid,imsi,msisdn,status}                                                                 |  (error)                      |
|2.12 日用量查询                  	|  url.Values{iccid,imsi,msisdn,date}                                                                   |  (*DailyUsage, error)         |
|2.13 日用量按照时间范围查询           |  url.Values{iccid,imsi,msisdn,startDate,endDate}                                                      |  ([]DailyUsage, error)        |
|2.14 取消测试期               		|  url.Values{iccid,imsi,msisdn}                                                                        |  (error)                      |
|2.15 更新备注                    	|  url.Values{iccid,imsi,msisdn,memo}                                                                   |  (error)                      |
|2.16 批量更新备注            		|  url.Values{iccids,imsis,msisdns,memo}                                                                |  (error)                      |
|3.1 流量池详情			           	|  url.Values{iccid,imsi,msisdn}                                                                        |  (*Pool, error)               |
|3.2 用户下所有流量池信息             	|  无                                                                                                   |  ([]Pool, error)              |
|4.1 提交实名认证信息		   			|  url.Values{iccid,imsi,msisdn,name,licenseType,licenseCode,phone,extenalUserName,pic1,pic2,pic3}      |  (error)                      |
|5.1 短信下发接口			            |  url.Values{iccid,imsi,msisdn,text,msgId}                                                             |  (error)                      |
|5.2 短信查询                        |  url.Values{iccid,imsi,msisdn,pageNo}                                                                 |  (*SmsList, error)            |
