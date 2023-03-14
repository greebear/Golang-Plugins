package main

import (
	"context"
	"fmt"
	"github.com/greebear/Golang-Plugins/plugin/payment/schema"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type paymentImpl string

const discount = 0.5

func (p paymentImpl) Pay(ctx context.Context, req *schema.PaymentReq) (*schema.PaymentRsp, error) {
	// Edge Case
	if req.Price == 0 {
		return nil, errors.Errorf("wrong price=%v", req.Price)
	}

	// Wechat with 50% discounts
	refund := req.Price * discount

	// Set Rsp
	rsp := &schema.PaymentRsp{
		SerialNumber: fmt.Sprintf("alipay-%v", uuid.NewV4()),
		Discount:     discount,
		Refund:       refund,
	}
	return rsp, nil
}

var Payer paymentImpl
