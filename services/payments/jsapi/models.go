package jsapi

import (
	"encoding/json"
	"fmt"
)

// Amount
type Amount struct {
	// 订单总金额，单位为分
	Total    *int64  `json:"total"`
	Currency *string `json:"currency,omitempty"`
}

func (o Amount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Total == nil {
		return nil, fmt.Errorf("field `Total` is required and must be specified in Amount")
	}
	toSerialize["total"] = o.Total

	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

func (o Amount) String() string {
	var ret string
	if o.Total == nil {
		ret += "Total:<nil>, "
	} else {
		ret += fmt.Sprintf("Total:%v, ", *o.Total)
	}

	if o.Currency == nil {
		ret += "Currency:<nil>"
	} else {
		ret += fmt.Sprintf("Currency:%v", *o.Currency)
	}

	return fmt.Sprintf("Amount{%s}", ret)
}

func (o Amount) Clone() *Amount {
	ret := Amount{}

	if o.Total != nil {
		ret.Total = new(int64)
		*ret.Total = *o.Total
	}

	if o.Currency != nil {
		ret.Currency = new(string)
		*ret.Currency = *o.Currency
	}

	return &ret
}

// CloseOrderRequest
type CloseOrderRequest struct {
	OutTradeNo *string `json:"out_trade_no"`
	// 直连商户号
	Mchid *string `json:"mchid"`
}

// CloseRequest
type CloseRequest struct {
	// 直连商户号
	Mchid *string `json:"mchid"`
}

// Detail 优惠功能
type Detail struct {
	// 1.商户侧一张小票订单可能被分多次支付，订单原价用于记录整张小票的交易金额。 2.当订单原价与支付金额不相等，则不享受优惠。 3.该字段主要用于防止同一张小票分多次支付，以享受多次优惠的情况，正常支付订单不必上传此参数。
	CostPrice *int64 `json:"cost_price,omitempty"`
	// 商家小票ID。
	InvoiceId   *string       `json:"invoice_id,omitempty"`
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

// GoodsDetail
type GoodsDetail struct {
	// 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成。
	MerchantGoodsId *string `json:"merchant_goods_id"`
	// 微信支付定义的统一商品编号（没有可不传）。
	WechatpayGoodsId *string `json:"wechatpay_goods_id,omitempty"`
	// 商品的实际名称。
	GoodsName *string `json:"goods_name,omitempty"`
	// 用户购买的数量。
	Quantity *int64 `json:"quantity"`
	// 商品单价，单位为分。
	UnitPrice *int64 `json:"unit_price"`
}

// Payer
type Payer struct {
	// 用户在商户appid下的唯一标识。
	Openid *string `json:"openid,omitempty"`
}

// PrepayRequest
type PrepayRequest struct {
	// 公众号ID
	Appid *string `json:"appid"`
	// 直连商户号
	Mchid *string `json:"mchid"`
	// 商品描述
	Description *string `json:"description"`
	// 商户订单号
	OutTradeNo *string `json:"out_trade_no"`
	// 附加数据
	Attach *string `json:"attach,omitempty"`
	// 有效性：1. HTTPS；2. 不允许携带查询串。
	NotifyUrl *string `json:"notify_url"`
	// 商品标记，代金券或立减优惠功能的参数。
	GoodsTag *string `json:"goods_tag,omitempty"`

	// 交易类型, JSAPI：公众号支付, NATIVE：扫码支付, APP：App支付, MWEB：H5支付, MICROPAY：付款码支付
	TradeType *string `json:"trade_type"` // JSAPI

	// 订单生成时间，格式为rfc3339格式
	TimeStart *string `json:"time_start,omitempty"`
	// 订单失效时间，格式为rfc3339格式
	TimeExpire *string `json:"time_expire,omitempty"`

	// 商户行业编码，值列表详见商户行业编码
	MerchantCategoryCode *string `json:"merchant_category_code"`

	// 支付者
	Payer *Payer `json:"payer"`
	// 订单金额
	Amount *Amount `json:"amount"`
	// 场景信息
	SceneInfo *SceneInfo `json:"scene_info,omitempty"`
	// 优惠功能
	Detail *Detail `json:"detail,omitempty"`
}

// PrepayResponse
type PrepayResponse struct {
	// 预支付交易会话标识
	PrepayId *string `json:"prepay_id"`
}

// QueryOrderByIdRequest
type QueryOrderByIdRequest struct {
	TransactionId *string `json:"transaction_id"`
	// 直连商户号
	Mchid *string `json:"mchid"`
}

// QueryOrderByOutTradeNoRequest
type QueryOrderByOutTradeNoRequest struct {
	OutTradeNo *string `json:"out_trade_no"`
	// 直连商户号
	Mchid *string `json:"mchid"`
}

// SceneInfo 支付场景描述
type SceneInfo struct {
	// 商户端设备号
	DeviceId *string `json:"device_id,omitempty"`
	DeviceIp *string `json:"device_ip,omitempty"`
	// 用户终端IP
	PayerClientIp *string `json:"payer_client_ip"`
	// 操作员ID
	OperatorId *string `json:"operator_id"`
	// 商户门店信息
	StoreInfo *StoreInfo `json:"store_info,omitempty"`
}

// SettleInfo
type SettleInfo struct {
	// 是否指定分账
	ProfitSharing *bool `json:"profit_sharing,omitempty"`
}

// StoreInfo 商户门店信息
type StoreInfo struct {
	// 商户侧门店编号
	Id *string `json:"id"`
	// 商户侧门店名称
	Name *string `json:"name,omitempty"`
	// 详细的商户门店地址
	Address *string `json:"address,omitempty"`
}
