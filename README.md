```bash
go build -buildmode=plugin -o plugin/payment/so/wechat.so plugin/payment/impl/wechat/wechat_payment.go
go build -buildmode=plugin -o plugin/payment/so/alipay.so plugin/payment/impl/alipay/alipay_payment.go

go build -gcflags "all=-N -l" -x -v -buildmode=plugin -o plugin/payment/so/wechat.so plugin/payment/impl/wechat/wechat_payment.go
```