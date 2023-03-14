package payment

import (
	"context"
	"github.com/greebear/Golang-Plugins/plugin/payment/schema"
)

type Payer interface {
	Pay(ctx context.Context, req *schema.PaymentReq) (*schema.PaymentRsp, error)
}
