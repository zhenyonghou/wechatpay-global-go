package refund_global

import (
	"encoding/json"
	"fmt"
	"time"
)

// Account * `AVAILABLE` - 可用余额, 多账户资金准备退款可用余额出资账户类型 * `UNAVAILABLE` - 不可用余额, 多账户资金准备退款不可用余额出资账户类型
type FundSource string

const (
	FundsRefundableBalance FundSource = "FUNDS_REFUNDABLE_BALANCE" // 可垫付退款余额
	OrderRefundableBalance FundSource = "ORDER_REFUNDABLE_BALANCE" // 订单未分可退余额
)

// Amount
type Amount struct {
	// 退款金额，币种的最小单位，只能为整数，不能超过原订单支付金额，如果有使用券，后台会按比例退
	Refund *int64 `json:"refund"`
	// 退款币种，符合ISO 4217标准的三位字母代码
	Currency *string `json:"currency"`

	// 退款给用户的金额，不包含所有优惠券金额
	PayerRefund *int64 `json:"payer_refund"`
	// 符合ISO 4217标准的三位字母代码
	PayerCurrency *string `json:"payer_currency"`

	// 去掉非充值代金券退款金额后的退款金额，单位为分，退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
	SettlementRefund *int64 `json:"settlement_refund"`
	// 结算币种，符合ISO 4217标准的三位字母代码
	SettlementCurrency *string `json:"settlement_currency"`

	ExchangeRate *ExchangeRate `json:"exchange_rate"`

	// 退款出资的账户类型及金额信息
	From []FundsFromItem `json:"from,omitempty"`
}

type ExchangeRate struct {
	Type string `json:"type"` // 汇率类型
	Rate *int   `json:"rate"` // 汇率值 rate值是兑换比例乘以10的8次方
}

// AmountReq
type AmountReq struct {
	// 退款金额，币种的最小单位，只能为整数，不能超过原订单支付金额。
	Refund *int64 `json:"refund"`
	// 退款需要从指定账户出资时，传递此参数指定出资金额（币种的最小单位，只能为整数）。 同时指定多个账户出资退款的使用场景需要满足以下条件：1、未开通退款支出分离产品功能；2、订单属于分账订单，且分账处于待分账或分账中状态。 参数传递需要满足条件：1、基本账户可用余额出资金额与基本账户不可用余额出资金额之和等于退款金额；2、账户类型不能重复。 上述任一条件不满足将返回错误
	From []FundsFromItem `json:"from,omitempty"`
	// 原支付交易的订单总金额，币种的最小单位，只能为整数。
	Total *int64 `json:"total"`
	// 符合ISO 4217标准的三位字母代码，退款币种必须和标价币种一致，币种列表详见币种类型。
	Currency *string `json:"currency"`
}

func (o AmountReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Refund == nil {
		return nil, fmt.Errorf("field `Refund` is required and must be specified in AmountReq")
	}
	toSerialize["refund"] = o.Refund

	if o.From != nil {
		toSerialize["from"] = o.From
	}

	if o.Total == nil {
		return nil, fmt.Errorf("field `Total` is required and must be specified in AmountReq")
	}
	toSerialize["total"] = o.Total

	if o.Currency == nil {
		return nil, fmt.Errorf("field `Currency` is required and must be specified in AmountReq")
	}
	toSerialize["currency"] = o.Currency
	return json.Marshal(toSerialize)
}

func (o AmountReq) String() string {
	var ret string
	if o.Refund == nil {
		ret += "Refund:<nil>, "
	} else {
		ret += fmt.Sprintf("Refund:%v, ", *o.Refund)
	}

	ret += fmt.Sprintf("From:%v, ", o.From)

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

	return fmt.Sprintf("AmountReq{%s}", ret)
}

func (o AmountReq) Clone() *AmountReq {
	ret := AmountReq{}

	if o.Refund != nil {
		ret.Refund = new(int64)
		*ret.Refund = *o.Refund
	}

	if o.From != nil {
		ret.From = make([]FundsFromItem, len(o.From))
		for i, item := range o.From {
			ret.From[i] = *item.Clone()
		}
	}

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

// Channel * `ORIGINAL` - 原路退款, 退款渠道 * `BALANCE` - 退回到余额, 退款渠道 * `OTHER_BALANCE` - 原账户异常退到其他余额账户, 退款渠道 * `OTHER_BANKCARD` - 原银行卡异常退到其他银行卡, 退款渠道
type Channel string

func (e Channel) Ptr() *Channel {
	return &e
}

// Enums of Channel
const (
	CHANNEL_ORIGINAL       Channel = "ORIGINAL"
	CHANNEL_BALANCE        Channel = "BALANCE"
	CHANNEL_OTHER_BALANCE  Channel = "OTHER_BALANCE"
	CHANNEL_OTHER_BANKCARD Channel = "OTHER_BANKCARD"
)

// CreateRequest
type CreateRequest struct {
	// 微信支付分配的商户号, 仅适用于直连模式
	Mchid *string `json:"mchid,omitempty"`
	// 商户在微信开放平台申请移动应用对应的APPID, 仅适用于直连模式
	Appid *string `json:"appid,omitempty"`

	// 微信支付分配给机构的商户号, 仅适用于机构模式
	SpMchid *string `json:"sp_mchid,omitempty"`
	// 微信支付分配子商户的商户号, 仅适用于机构模式
	SubMchid *string `json:"sub_mchid,omitempty"`
	// 商户在微信公众平台申请服务号对应的APPID, 仅适用于机构模式
	SpAppid *string `json:"sp_appid,omitempty"`
	// 子商户APPID, 仅适用于机构模式
	SubAppid *string `json:"sub_appid,omitempty"`

	// 原支付交易对应的微信订单号
	TransactionId *string `json:"transaction_id,omitempty"`
	// 原支付交易对应的商户订单号
	OutTradeNo *string `json:"out_trade_no,omitempty"`

	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo *string `json:"out_refund_no"`

	// 若商户传入，会在下发给用户的退款消息中体现退款原因
	Reason *string `json:"reason,omitempty"`

	// 异步接收微信支付退款结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。 如果参数中传了notify_url，则商户平台上配置的回调地址将不会生效，优先回调当前传的这个地址。
	NotifyUrl *string `json:"notify_url,omitempty"`
	// 订单金额信息
	Amount *AmountReq `json:"amount"`
}

// FundsAccount * `UNSETTLED` - 未结算资金, 退款所使用资金对应的资金账户类型 * `AVAILABLE` - 可用余额, 退款所使用资金对应的资金账户类型 * `UNAVAILABLE` - 不可用余额, 退款所使用资金对应的资金账户类型 * `OPERATION` - 运营户, 退款所使用资金对应的资金账户类型 * `BASIC` - 基本账户（含可用余额和不可用余额）, 退款所使用资金对应的资金账户类型
type FundsAccount string

func (e FundsAccount) Ptr() *FundsAccount {
	return &e
}

// Enums of FundsAccount
const (
	FUNDSACCOUNT_UNSETTLED   FundsAccount = "UNSETTLED"
	FUNDSACCOUNT_AVAILABLE   FundsAccount = "AVAILABLE"
	FUNDSACCOUNT_UNAVAILABLE FundsAccount = "UNAVAILABLE"
	FUNDSACCOUNT_OPERATION   FundsAccount = "OPERATION"
	FUNDSACCOUNT_BASIC       FundsAccount = "BASIC"
)

// FundsFromItem
type FundsFromItem struct {
	// 下面枚举值多选一。 枚举值： AVAILABLE : 可用余额 UNAVAILABLE : 不可用余额 * `AVAILABLE` - 可用余额 * `UNAVAILABLE` - 不可用余额
	FundSource *FundSource `json:"fund_source"`
	// 对应账户出资金额
	Amount *int64 `json:"amount"`
}

func (o FundsFromItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.FundSource == nil {
		return nil, fmt.Errorf("field `Account` is required and must be specified in FundsFromItem")
	}
	toSerialize["fund_source"] = o.FundSource

	if o.Amount == nil {
		return nil, fmt.Errorf("field `Amount` is required and must be specified in FundsFromItem")
	}
	toSerialize["amount"] = o.Amount
	return json.Marshal(toSerialize)
}

func (o FundsFromItem) String() string {
	var ret string
	if o.FundSource == nil {
		ret += "FundSource:<nil>, "
	} else {
		ret += fmt.Sprintf("FundSource:%v, ", *o.FundSource)
	}

	if o.Amount == nil {
		ret += "Amount:<nil>"
	} else {
		ret += fmt.Sprintf("Amount:%v", *o.Amount)
	}

	return fmt.Sprintf("FundsFromItem{%s}", ret)
}

func (o FundsFromItem) Clone() *FundsFromItem {
	ret := FundsFromItem{}

	if o.FundSource != nil {
		ret.FundSource = new(FundSource)
		*ret.FundSource = *o.FundSource
	}

	if o.Amount != nil {
		ret.Amount = new(int64)
		*ret.Amount = *o.Amount
	}

	return &ret
}

// GoodsDetail
type GoodsDetail struct {
	// 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成
	MerchantGoodsId *string `json:"merchant_goods_id"`
	// 微信支付定义的统一商品编号（没有可不传）
	WechatpayGoodsId *string `json:"wechatpay_goods_id,omitempty"`
	// 商品的实际名称
	GoodsName *string `json:"goods_name,omitempty"`
	// 商品单价金额，单位为分
	UnitPrice *int64 `json:"unit_price"`
	// 商品退款金额，单位为分
	RefundAmount *int64 `json:"refund_amount"`
	// 对应商品的退货数量
	RefundQuantity *int64 `json:"refund_quantity"`
}

func (o GoodsDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.MerchantGoodsId == nil {
		return nil, fmt.Errorf("field `MerchantGoodsId` is required and must be specified in GoodsDetail")
	}
	toSerialize["merchant_goods_id"] = o.MerchantGoodsId

	if o.WechatpayGoodsId != nil {
		toSerialize["wechatpay_goods_id"] = o.WechatpayGoodsId
	}

	if o.GoodsName != nil {
		toSerialize["goods_name"] = o.GoodsName
	}

	if o.UnitPrice == nil {
		return nil, fmt.Errorf("field `UnitPrice` is required and must be specified in GoodsDetail")
	}
	toSerialize["unit_price"] = o.UnitPrice

	if o.RefundAmount == nil {
		return nil, fmt.Errorf("field `RefundAmount` is required and must be specified in GoodsDetail")
	}
	toSerialize["refund_amount"] = o.RefundAmount

	if o.RefundQuantity == nil {
		return nil, fmt.Errorf("field `RefundQuantity` is required and must be specified in GoodsDetail")
	}
	toSerialize["refund_quantity"] = o.RefundQuantity
	return json.Marshal(toSerialize)
}

func (o GoodsDetail) String() string {
	var ret string
	if o.MerchantGoodsId == nil {
		ret += "MerchantGoodsId:<nil>, "
	} else {
		ret += fmt.Sprintf("MerchantGoodsId:%v, ", *o.MerchantGoodsId)
	}

	if o.WechatpayGoodsId == nil {
		ret += "WechatpayGoodsId:<nil>, "
	} else {
		ret += fmt.Sprintf("WechatpayGoodsId:%v, ", *o.WechatpayGoodsId)
	}

	if o.GoodsName == nil {
		ret += "GoodsName:<nil>, "
	} else {
		ret += fmt.Sprintf("GoodsName:%v, ", *o.GoodsName)
	}

	if o.UnitPrice == nil {
		ret += "UnitPrice:<nil>, "
	} else {
		ret += fmt.Sprintf("UnitPrice:%v, ", *o.UnitPrice)
	}

	if o.RefundAmount == nil {
		ret += "RefundAmount:<nil>, "
	} else {
		ret += fmt.Sprintf("RefundAmount:%v, ", *o.RefundAmount)
	}

	if o.RefundQuantity == nil {
		ret += "RefundQuantity:<nil>"
	} else {
		ret += fmt.Sprintf("RefundQuantity:%v", *o.RefundQuantity)
	}

	return fmt.Sprintf("GoodsDetail{%s}", ret)
}

func (o GoodsDetail) Clone() *GoodsDetail {
	ret := GoodsDetail{}

	if o.MerchantGoodsId != nil {
		ret.MerchantGoodsId = new(string)
		*ret.MerchantGoodsId = *o.MerchantGoodsId
	}

	if o.WechatpayGoodsId != nil {
		ret.WechatpayGoodsId = new(string)
		*ret.WechatpayGoodsId = *o.WechatpayGoodsId
	}

	if o.GoodsName != nil {
		ret.GoodsName = new(string)
		*ret.GoodsName = *o.GoodsName
	}

	if o.UnitPrice != nil {
		ret.UnitPrice = new(int64)
		*ret.UnitPrice = *o.UnitPrice
	}

	if o.RefundAmount != nil {
		ret.RefundAmount = new(int64)
		*ret.RefundAmount = *o.RefundAmount
	}

	if o.RefundQuantity != nil {
		ret.RefundQuantity = new(int64)
		*ret.RefundQuantity = *o.RefundQuantity
	}

	return &ret
}

// QueryByOutRefundNoRequest
type QueryByOutRefundNoRequest struct {
	Mchid       *string `json:"mchid"`
	OutRefundNo *string `json:"out_refund_no"`
}

// Refund
type Refund struct {
	// 微信支付退款号
	Id *string `json:"id"`
	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo *string `json:"out_refund_no"`

	// 退款受理时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。
	CreateTime *time.Time `json:"create_time"`

	// 退款金额
	Amount *Amount `json:"amount"`

	// 优惠退款详情
	Detail *PromotionDetail `json:"detail"`
}

type PromotionDetail struct {
}

// ReqFundsAccount * `AVAILABLE` - 可用余额, 仅对老资金流商户适用，指定从可用余额账户出资
type ReqFundsAccount string

func (e ReqFundsAccount) Ptr() *ReqFundsAccount {
	return &e
}

// Enums of ReqFundsAccount
const (
	REQFUNDSACCOUNT_AVAILABLE ReqFundsAccount = "AVAILABLE"
)

// Scope * `GLOBAL` - 全场代金券, 全场优惠类型 * `SINGLE` - 单品优惠, 单品优惠类型
type Scope string

func (e Scope) Ptr() *Scope {
	return &e
}

// Enums of Scope
const (
	SCOPE_GLOBAL Scope = "GLOBAL"
	SCOPE_SINGLE Scope = "SINGLE"
)

// Status * `SUCCESS` - 退款成功, 退款状态 * `CLOSED` - 退款关闭, 退款状态 * `PROCESSING` - 退款处理中, 退款状态 * `ABNORMAL` - 退款异常, 退款状态
type Status string

func (e Status) Ptr() *Status {
	return &e
}

// Enums of Status
const (
	STATUS_SUCCESS    Status = "SUCCESS"
	STATUS_CLOSED     Status = "CLOSED"
	STATUS_PROCESSING Status = "PROCESSING"
	STATUS_ABNORMAL   Status = "ABNORMAL"
)

// Type * `COUPON` - 代金券, 代金券类型，需要走结算资金的充值型代金券 * `DISCOUNT` - 优惠券, 优惠券类型，不走结算资金的免充值型优惠券
type Type string

func (e Type) Ptr() *Type {
	return &e
}

// Enums of Type
const (
	TYPE_COUPON   Type = "COUPON"
	TYPE_DISCOUNT Type = "DISCOUNT"
)
