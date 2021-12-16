# 说明
用于便捷解析JSON字符串

#使用方法

```go
package main

import (
	"encoding/json"
	"fmt"
	"github-universaljson/universaljson"
)

var (
	data = `{
		"store_code": "xGpS0mPbNqo5FMQcdTk4nR1DJr9LUEfa",
		"main_orderno": "LWM2112149515095973",
		"payment_method":4,
		"orderno": "LWS2112148253015566",
		"fictitious_id": "4",
		"platform": 1,
		"is_expedited": 0,
		"is_topay": 0,
		"storage":0,
		"attach_info": [],
		"pay_money": 1.01,
		"country":"中国",
		"province": "重庆",
		"city": "重庆市",
		"town": "渝北区",
		"address": "XXX街道XXX小区XXX单元X-X",
		"receiver_name": "张三",
		"receiver_phone": "13456789100",
		"logistics_id": 2,
		"product": [{
			"goods_code": "010046701091",
			"number": 1,
			"short_code": "010046701091",
			"goods_years": 2009,
			"volume": "750ml"
		}]
	}`
)

func p(v interface{}) {
	fmt.Println(v)
}

func main() {
	parser := universaljson.ParseJSON(&data)

	if parser == nil {
		return
	}
	id, err := parser.GetInt64("logistics_id")
	if err == nil {
		p(id)
	} else {
		fmt.Println(err)
	}
	pm, e := parser.GetFloat64("pay_money")
	if e != nil {
		p(e)
	} else {
		p(pm)
	}
	product, err := parser.GetArray("product")[0].GetString("volume")
	if err == nil {
		p(product)
	}
	goods_years, err := parser.GetArray("product")[0].GetInt64("goods_years")
	if err == nil {
		p(goods_years)
	}
}

```
