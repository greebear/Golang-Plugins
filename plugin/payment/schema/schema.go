package schema

type PaymentReq struct {
	OrderId  string
	UserId   uint64
	ItemList []uint64
	Price    float64
}

type PaymentRsp struct {
	SerialNumber string
	Discount     float64
	Refund       float64
}
