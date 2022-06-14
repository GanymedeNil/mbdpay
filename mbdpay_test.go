package mbdpay

import (
	"context"
	"os"
	"testing"

	"github.com/GanymedeNil/mbdpay/util"
)

var appId, appKey, openId string

func init() {
	appId = os.Getenv("APPID")
	appKey = os.Getenv("APPKEY")
	openId = os.Getenv("OPENID")
}

func TestWxH5(t *testing.T) {
	mm := WxPrepayJSONRequestBody{
		AmountTotal: 1,
		AppId:       appId,
		Description: "测试商品",
	}
	channel := "h5"
	mm.Channel = &channel
	mm.Sign, _ = util.Sign(mm, appKey)
	c, _ := NewMbdPay()
	res, _ := c.WxPrepayWithResponse(context.Background(), mm)
	if res.JSON200.Error != nil {
		t.Fatal(res.JSON200.Error)
	}
}

func TestWxJS(t *testing.T) {
	mm := WxPrepayJSONRequestBody{
		AmountTotal: 1,
		AppId:       appId,
		Description: "测试商品",
	}
	mm.Openid = &openId
	callbackUrl := "http://www.baidu.com"
	mm.CallbackUrl = &callbackUrl
	mm.Sign, _ = util.Sign(mm, appKey)
	c, _ := NewMbdPay()
	res, _ := c.WxPrepayWithResponse(context.Background(), mm)
	if res.JSON200.AppId == nil || res.JSON200.NonceStr == nil || res.JSON200.Package == nil || res.JSON200.PaySign == nil {
		t.Fail()
	}
}

func TestAlipayPay(t *testing.T) {
	mm := AlipayPayJSONRequestBody{
		AmountTotal: 1,
		AppId:       appId,
		Description: "测试商品",
		Url:         "http://www.example.com",
	}
	mm.Sign, _ = util.Sign(mm, appKey)
	c, _ := NewMbdPay()
	res, _ := c.AlipayPayWithResponse(context.Background(), mm)
	if res.JSON200.Error != nil || res.JSON200.Body == nil {
		t.Fatal(res.JSON200.Error)
	}
}

func TestRefund(t *testing.T) {
	mm := RefundJSONRequestBody{
		OrderId: "36b4272290745783269904952a40d6d1",
		AppId:   appId,
	}
	mm.Sign, _ = util.Sign(mm, appKey)
	c, _ := NewMbdPay()
	res, _ := c.RefundWithResponse(context.Background(), mm)
	if res.JSON200.Message == nil {
		t.Fail()
	}
}

func TestSearchOrder(t *testing.T) {
	mm := SearchOrderJSONRequestBody{
		OutTradeNo: "36b4272290745783269904952a40d6d1",
		AppId:      appId,
	}
	mm.Sign, _ = util.Sign(mm, appKey)
	c, _ := NewMbdPay()
	res, _ := c.SearchOrderWithResponse(context.Background(), mm)
	if res.JSON200.Error == nil {
		t.Fail()
	}
}
