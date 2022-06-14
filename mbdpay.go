package mbdpay

const API = "https://api.mianbaoduo.com"

func NewMbdPay() (*ClientWithResponses, error) {
	return NewClientWithResponses(API)
}
