package util

import (
	"testing"

	"github.com/GanymedeNil/mbdpay"
)

func TestSign(t *testing.T) {
	testData := mbdpay.WxPrepayReq{
		AmountTotal: 1,
		AppId:       "12345",
	}
	outTradeNo := "123123123123"
	testData.OutTradeNo = &outTradeNo
	testKey := "1234567890"
	want := "2f370f09c2331caff69a86911f124323"
	if got, _ := Sign(testData, testKey); got != want {
		t.Errorf("Sign() = %v, want %v", got, want)
	}
}
