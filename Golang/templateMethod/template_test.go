package templateMethod

import (
	"fmt"
	"testing"
)

type PayMethod interface {
	pay()
}

type DistributionMethod interface {
	express()
}

type Template struct {
	PayMethod
	DistributionMethod
}

func (t *Template) process() {
	t.pay()
	t.express()
}

type AliPay struct {
}

func (a *AliPay) pay() {
	fmt.Println("这是支付宝支付")
}

type WechatPay struct {
}

func (a *WechatPay) pay() {
	fmt.Println("这是微信支付")
}

type Online struct {
}

func (o *Online) express() {
	fmt.Println("这是线上发货方式")
}

type Offline struct {
}

func (o *Offline) express() {
	fmt.Println("这是线下发货方式")
}

func NewTemplate(pay PayMethod, express DistributionMethod) *Template {
	return &Template{
		PayMethod:          pay,
		DistributionMethod: express,
	}
}

func TestNewAliPayOffline(t *testing.T) {
	pay := &AliPay{}
	express := &Offline{}
	a := NewTemplate(pay, express)
	a.process()
}

func TestNewWechatPayOnline(t *testing.T) {
	pay := &WechatPay{}
	express := &Online{}
	a := NewTemplate(pay, express)
	a.process()
}
