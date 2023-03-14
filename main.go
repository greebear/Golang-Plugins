package main

import (
	"context"
	"fmt"
	"github.com/greebear/Golang-Plugins/plugin/payment"
	"github.com/greebear/Golang-Plugins/plugin/payment/schema"
	"plugin"
)

func main() {
	ctx := context.Background()
	req := &schema.PaymentReq{
		OrderId:  "my_order_001",
		UserId:   10000001,
		ItemList: []uint64{101, 102, 103},
		Price:    100,
	}

	// wechat
	modWechat := "./plugin/payment/so/wechat.so"
	plugWechat, err := plugin.Open(modWechat)
	if err != nil {
		panic(err)
	}

	// common
	payer, err := plugWechat.Lookup("Payer")
	if err != nil {
		panic(err)
	}
	payImpl, ok := payer.(payment.Payer)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		panic(10086)
	}
	wechatRsp, err := payImpl.Pay(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("SN:%v|discount:%v|refund:%v", wechatRsp.SerialNumber, wechatRsp.Discount, wechatRsp.Refund)
}
