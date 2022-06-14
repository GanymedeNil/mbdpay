// Package mbdpay provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package mbdpay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// AlipayPayReq defines model for AlipayPayReq.
type AlipayPayReq struct {
	// 金额，单位为分
	AmountTotal int `json:"amount_total"`

	// 你的 app_id，可在控制台查看
	AppId string `json:"app_id"`

	// 支付后跳转地址，如不填会只显示「支付成功」
	CallbackUrl *string `json:"callback_url,omitempty"`

	// 支付描述，一般为商品名称
	Description string `json:"description"`

	// 订单号，如不填，面包多将随机生成订单号
	OutTradeNo *string `json:"out_trade_no,omitempty"`

	// 分账参数，需先开通分账权限
	ShareCode *string `json:"share_code,omitempty"`

	// 请求签名，参照签名算法
	Sign string `json:"sign"`

	// 支付后跳转的URL地址
	Url string `json:"url"`
}

// AlipayPayRes defines model for AlipayPayRes.
type AlipayPayRes struct {
	Body *string `json:"body,omitempty"`
	Code *string `json:"code,omitempty"`

	// 错误说明
	Error   *string `json:"error,omitempty"`
	Message *string `json:"message,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Code *string `json:"code,omitempty"`

	// 错误说明
	Error   *string `json:"error,omitempty"`
	Message *string `json:"message,omitempty"`
}

// RefundReq defines model for RefundReq.
type RefundReq struct {
	// 你的 app_id，可在控制台查看
	AppId string `json:"app_id"`

	// 订单号
	OrderId string `json:"order_id"`

	// 请求签名，参照签名算法
	Sign string `json:"sign"`
}

// RefundRes defines model for RefundRes.
type RefundRes struct {
	Code *string `json:"code,omitempty"`

	// 错误说明
	Error   *string `json:"error,omitempty"`
	Info    *string `json:"info,omitempty"`
	Message *string `json:"message,omitempty"`
}

// SearchOrderReq defines model for SearchOrderReq.
type SearchOrderReq struct {
	// 你的 app_id，可在控制台查看
	AppId string `json:"app_id"`

	// 订单号（也支持微信/支付宝流水号）
	OutTradeNo string `json:"out_trade_no"`

	// 请求签名，参照签名算法
	Sign string `json:"sign"`
}

// SearchOrderRes defines model for SearchOrderRes.
type SearchOrderRes struct {
	// 支付金额，单位为分
	Amount *int `json:"amount,omitempty"`

	// 支付渠道流水号
	ChargeId *string `json:"charge_id,omitempty"`
	Code     *string `json:"code,omitempty"`

	// 支付时间（时间戳）
	CreateTime *int `json:"create_time,omitempty"`

	// 商品描述
	Description *string `json:"description,omitempty"`

	// 错误说明
	Error   *string `json:"error,omitempty"`
	Message *string `json:"message,omitempty"`

	// 订单号
	OrderId *string `json:"order_id,omitempty"`

	// 支付渠道，1为微信支付，2为支付宝
	Payway *int `json:"payway,omitempty"`

	// 附加参数（json格式）
	Plusinfo *string `json:"plusinfo,omitempty"`

	// 已退款金额，单位为分
	RefundAmount *int `json:"refund_amount,omitempty"`

	// 退款状态，0为无退款，1为部分退款，2为全部退款
	RefundState *int `json:"refund_state,omitempty"`

	// 结算ID
	ShareId *string `json:"share_id,omitempty"`

	// 结算状态
	ShareState *int `json:"share_state,omitempty"`

	// 支付状态，0-未支付，1-已支付，2-已结算，3-投诉中，4-投诉完结，5-投诉超时，6-投诉中(买家处理中)
	State *int `json:"state,omitempty"`
}

// WxPrepayReq defines model for WxPrepayReq.
type WxPrepayReq struct {
	// 金额，单位为分
	AmountTotal int `json:"amount_total"`

	// 你的 app_id，可在控制台查看
	AppId string `json:"app_id"`

	// 支付后跳转地址(JSAPI支付该字段为必填)
	CallbackUrl *string `json:"callback_url,omitempty"`

	// 固定为 h5(h5支付该字段为必填)
	Channel *string `json:"channel,omitempty"`

	// 支付描述，一般为商品名称
	Description string `json:"description"`

	// 用户的openid (JSAPI支付该字段为必填)
	Openid *string `json:"openid,omitempty"`

	// 订单号，如不填，面包多将随机生成订单号
	OutTradeNo *string `json:"out_trade_no,omitempty"`

	// 分账参数，需先开通分账权限
	ShareCode *string `json:"share_code,omitempty"`

	// 请求签名
	Sign string `json:"sign"`
}

// WxPrepayRes defines model for WxPrepayRes.
type WxPrepayRes struct {
	AppId *string `json:"appId,omitempty"`
	Code  *string `json:"code,omitempty"`

	// 错误说明
	Error     *string `json:"error,omitempty"`
	H5Url     *string `json:"h5_url,omitempty"`
	Message   *string `json:"message,omitempty"`
	NonceStr  *string `json:"nonceStr,omitempty"`
	Package   *string `json:"package,omitempty"`
	PaySign   *string `json:"paySign,omitempty"`
	SignType  *string `json:"signType,omitempty"`
	TimeStamp *string `json:"timeStamp,omitempty"`
}

// AlipayPayJSONBody defines parameters for AlipayPay.
type AlipayPayJSONBody = AlipayPayReq

// RefundJSONBody defines parameters for Refund.
type RefundJSONBody = RefundReq

// SearchOrderJSONBody defines parameters for SearchOrder.
type SearchOrderJSONBody = SearchOrderReq

// WxPrepayJSONBody defines parameters for WxPrepay.
type WxPrepayJSONBody = WxPrepayReq

// AlipayPayJSONRequestBody defines body for AlipayPay for application/json ContentType.
type AlipayPayJSONRequestBody = AlipayPayJSONBody

// RefundJSONRequestBody defines body for Refund for application/json ContentType.
type RefundJSONRequestBody = RefundJSONBody

// SearchOrderJSONRequestBody defines body for SearchOrder for application/json ContentType.
type SearchOrderJSONRequestBody = SearchOrderJSONBody

// WxPrepayJSONRequestBody defines body for WxPrepay for application/json ContentType.
type WxPrepayJSONRequestBody = WxPrepayJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// AlipayPay request with any body
	AlipayPayWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AlipayPay(ctx context.Context, body AlipayPayJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Refund request with any body
	RefundWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Refund(ctx context.Context, body RefundJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// SearchOrder request with any body
	SearchOrderWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	SearchOrder(ctx context.Context, body SearchOrderJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// WxPrepay request with any body
	WxPrepayWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	WxPrepay(ctx context.Context, body WxPrepayJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) AlipayPayWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAlipayPayRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AlipayPay(ctx context.Context, body AlipayPayJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAlipayPayRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RefundWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRefundRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Refund(ctx context.Context, body RefundJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRefundRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SearchOrderWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSearchOrderRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SearchOrder(ctx context.Context, body SearchOrderJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSearchOrderRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) WxPrepayWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewWxPrepayRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) WxPrepay(ctx context.Context, body WxPrepayJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewWxPrepayRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewAlipayPayRequest calls the generic AlipayPay builder with application/json body
func NewAlipayPayRequest(server string, body AlipayPayJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAlipayPayRequestWithBody(server, "application/json", bodyReader)
}

// NewAlipayPayRequestWithBody generates requests for AlipayPay with any type of body
func NewAlipayPayRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/release/alipay/pay")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewRefundRequest calls the generic Refund builder with application/json body
func NewRefundRequest(server string, body RefundJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewRefundRequestWithBody(server, "application/json", bodyReader)
}

// NewRefundRequestWithBody generates requests for Refund with any type of body
func NewRefundRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/release/main/refund")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewSearchOrderRequest calls the generic SearchOrder builder with application/json body
func NewSearchOrderRequest(server string, body SearchOrderJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewSearchOrderRequestWithBody(server, "application/json", bodyReader)
}

// NewSearchOrderRequestWithBody generates requests for SearchOrder with any type of body
func NewSearchOrderRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/release/main/search_order")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewWxPrepayRequest calls the generic WxPrepay builder with application/json body
func NewWxPrepayRequest(server string, body WxPrepayJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewWxPrepayRequestWithBody(server, "application/json", bodyReader)
}

// NewWxPrepayRequestWithBody generates requests for WxPrepay with any type of body
func NewWxPrepayRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/release/wx/prepay")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// AlipayPay request with any body
	AlipayPayWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AlipayPayResponse, error)

	AlipayPayWithResponse(ctx context.Context, body AlipayPayJSONRequestBody, reqEditors ...RequestEditorFn) (*AlipayPayResponse, error)

	// Refund request with any body
	RefundWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*RefundResponse, error)

	RefundWithResponse(ctx context.Context, body RefundJSONRequestBody, reqEditors ...RequestEditorFn) (*RefundResponse, error)

	// SearchOrder request with any body
	SearchOrderWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SearchOrderResponse, error)

	SearchOrderWithResponse(ctx context.Context, body SearchOrderJSONRequestBody, reqEditors ...RequestEditorFn) (*SearchOrderResponse, error)

	// WxPrepay request with any body
	WxPrepayWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*WxPrepayResponse, error)

	WxPrepayWithResponse(ctx context.Context, body WxPrepayJSONRequestBody, reqEditors ...RequestEditorFn) (*WxPrepayResponse, error)
}

type AlipayPayResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AlipayPayRes
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r AlipayPayResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AlipayPayResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RefundResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RefundRes
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r RefundResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RefundResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type SearchOrderResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SearchOrderRes
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r SearchOrderResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SearchOrderResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type WxPrepayResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *WxPrepayRes
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r WxPrepayResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r WxPrepayResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// AlipayPayWithBodyWithResponse request with arbitrary body returning *AlipayPayResponse
func (c *ClientWithResponses) AlipayPayWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AlipayPayResponse, error) {
	rsp, err := c.AlipayPayWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAlipayPayResponse(rsp)
}

func (c *ClientWithResponses) AlipayPayWithResponse(ctx context.Context, body AlipayPayJSONRequestBody, reqEditors ...RequestEditorFn) (*AlipayPayResponse, error) {
	rsp, err := c.AlipayPay(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAlipayPayResponse(rsp)
}

// RefundWithBodyWithResponse request with arbitrary body returning *RefundResponse
func (c *ClientWithResponses) RefundWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*RefundResponse, error) {
	rsp, err := c.RefundWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRefundResponse(rsp)
}

func (c *ClientWithResponses) RefundWithResponse(ctx context.Context, body RefundJSONRequestBody, reqEditors ...RequestEditorFn) (*RefundResponse, error) {
	rsp, err := c.Refund(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRefundResponse(rsp)
}

// SearchOrderWithBodyWithResponse request with arbitrary body returning *SearchOrderResponse
func (c *ClientWithResponses) SearchOrderWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SearchOrderResponse, error) {
	rsp, err := c.SearchOrderWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSearchOrderResponse(rsp)
}

func (c *ClientWithResponses) SearchOrderWithResponse(ctx context.Context, body SearchOrderJSONRequestBody, reqEditors ...RequestEditorFn) (*SearchOrderResponse, error) {
	rsp, err := c.SearchOrder(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSearchOrderResponse(rsp)
}

// WxPrepayWithBodyWithResponse request with arbitrary body returning *WxPrepayResponse
func (c *ClientWithResponses) WxPrepayWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*WxPrepayResponse, error) {
	rsp, err := c.WxPrepayWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseWxPrepayResponse(rsp)
}

func (c *ClientWithResponses) WxPrepayWithResponse(ctx context.Context, body WxPrepayJSONRequestBody, reqEditors ...RequestEditorFn) (*WxPrepayResponse, error) {
	rsp, err := c.WxPrepay(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseWxPrepayResponse(rsp)
}

// ParseAlipayPayResponse parses an HTTP response from a AlipayPayWithResponse call
func ParseAlipayPayResponse(rsp *http.Response) (*AlipayPayResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AlipayPayResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AlipayPayRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseRefundResponse parses an HTTP response from a RefundWithResponse call
func ParseRefundResponse(rsp *http.Response) (*RefundResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RefundResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RefundRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseSearchOrderResponse parses an HTTP response from a SearchOrderWithResponse call
func ParseSearchOrderResponse(rsp *http.Response) (*SearchOrderResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &SearchOrderResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SearchOrderRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseWxPrepayResponse parses an HTTP response from a WxPrepayWithResponse call
func ParseWxPrepayResponse(rsp *http.Response) (*WxPrepayResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &WxPrepayResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest WxPrepayRes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
