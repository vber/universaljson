# 说明
用于便捷解析JSON字符串

#使用方法

```go
package main

import (
	"fmt"
	"github-universaljson/universaljson"
)

var (
	data = `{
		"store_code": "xGpS0mPbNqo5FMQcdTk4nR1DJr9LUEfa",
		"main_orderno": "LWM2112149515095973",
		"attach_info": [],
		"pay_money": 1.01,
		"country":"中国",
		"province": "重庆",
		"city": "重庆市",
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
```go
s := `{"error_code":-1,"error_msg":"订单未支付，请稍后再试","data":{"main_order_no":"LWM2112149515095973","sub_order_no":"LWS2112148253015566"}}`

e, err := universaljson.ParseJSON(s).GetObject("data").GetString("main_order_no")
if err == nil {
   fmt.Println(e)
}
```
