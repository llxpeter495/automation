package main

import (
	"encoding/json"
	"fmt"
)

type Pay interface {
	pay() string
}
type AliPay struct{}

type WeixinPay struct{}

type YinlianPay struct{}

func (a AliPay) pay() string {
	fmt.Println("支付宝pay")
	return "支付宝pay"
}

func (w WeixinPay) pay() string {
	fmt.Println("微信pay")
	return "微信pay"
}

func (y YinlianPay) pay() string {
	fmt.Println("银联pay")
	return "银联pay"
}

func main() {
	var p Pay
	p = AliPay{}
	p.pay()
	p = WeixinPay{}
	p.pay()
	p = YinlianPay{}
	p.pay()
	amap := map[string]interface{}{"name": "hehe", "age": 11}
	fmt.Println(amap)
	v, _ := json.Marshal(amap)
	fmt.Println(string(v))

}
