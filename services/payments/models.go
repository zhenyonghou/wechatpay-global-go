package payments

import "fmt"

type Transaction struct {
	Mchid      *string `json:"mchid,omitempty"`
	Appid      *string `json:"appid,omitempty"`
	OutTradeNo *string `json:"out_trade_no,omitempty"`

	Id *string `json:"id"` // 微信支付订单号
	//TransactionId   *string            `json:"transaction_id,omitempty"`
	Attach          *string            `json:"attach,omitempty"`
	TradeType       *string            `json:"trade_type"`          // 交易类型
	BankType        *string            `json:"bank_type,omitempty"` // 付款银行
	SuccessTime     *string            `json:"success_time"`        // 支付完成时间
	TradeState      *string            `json:"trade_state"`         // 交易状态
	TradeStateDesc  *string            `json:"trade_state_desc"`    // 交易状态描述
	Payer           *TransactionPayer  `json:"payer"`
	Amount          *TransactionAmount `json:"amount"`
	ExchangeRate    *ExchangeRate      `json:"exchange_rate,omitempty"`
	PromotionDetail []PromotionDetail  `json:"promotion_detail,omitempty"`
}

type TransactionPayer struct {
	Openid *string `json:"openid,omitempty"`
}

type TransactionAmount struct {
	Currency      *string `json:"currency,omitempty"`
	PayerCurrency *string `json:"payer_currency,omitempty"`
	PayerTotal    *int64  `json:"payer_total,omitempty"`
	Total         *int64  `json:"total,omitempty"`
}

type ExchangeRate struct {
	Type *string `json:"type,omitempty"` // SETTLEMENT_RATE,即标价币种和结算币种的汇率
	Rate *int    `json:"rate,omitempty"` // rate值是兑换比例乘以10的8次方
}

func (o ExchangeRate) String() string {
	var ret string
	if o.Type != nil {
		ret += fmt.Sprintf("Type:%s, ", *o.Type)
	}
	if o.Rate != nil {
		ret += fmt.Sprintf("Rate:%d, ", *o.Rate)
	}
	return fmt.Sprintf("ExchangeRate{%s}", ret)
}

type PromotionDetail struct {
	// 券ID
	PromotionId *string `json:"promotion_id"`
	// 优惠名称
	Name *string `json:"name,omitempty"`
	// GLOBAL：全场代金券；SINGLE：单品优惠
	Scope *string `json:"scope,omitempty"`
	// CASH：充值；NOCASH：预充值。
	Type *string `json:"type,omitempty"`
	// 优惠券面额
	Amount *int64 `json:"amount,omitempty"`
	// 活动ID，批次ID
	StockId *string `json:"stock_id,omitempty"`
	// 单位为分
	WechatpayContribute *int64 `json:"wechatpay_contribute,omitempty"`
	// 单位为分
	MerchantContribute *int64 `json:"merchant_contribute,omitempty"`
	// 单位为分
	OtherContribute *int64 `json:"other_contribute,omitempty"`
	// CNY：人民币，境内商户号仅支持人民币。
	Currency    *string                `json:"currency,omitempty"`
	GoodsDetail []PromotionGoodsDetail `json:"goods_detail,omitempty"`
}

type PromotionGoodsDetail struct {
	// 商品编码
	GoodsId *string `json:"goods_id"`
	// 商品数量
	Quantity *int64 `json:"quantity"`
	// 商品价格
	UnitPrice *int64 `json:"unit_price"`
	// 商品优惠金额
	DiscountAmount *int64 `json:"discount_amount"`
	// 商品备注
	GoodsRemark *string `json:"goods_remark,omitempty"`
}

func (o Transaction) String() string {
	var ret string
	if o.Mchid != nil {
		ret += fmt.Sprintf("Mchid:%s, ", *o.Mchid)
	}
	if o.Appid != nil {
		ret += fmt.Sprintf("Appid:%s, ", *o.Appid)
	}
	if o.OutTradeNo != nil {
		ret += fmt.Sprintf("OutTradeNo:%s, ", *o.OutTradeNo)
	}
	if o.Id != nil {
		ret += fmt.Sprintf("Id:%s, ", *o.Id)
	}
	if o.Attach != nil {
		ret += fmt.Sprintf("Attach:%s, ", *o.Attach)
	}
	if o.TradeType != nil {
		ret += fmt.Sprintf("TradeType:%s, ", *o.TradeType)
	}
	if o.BankType != nil {
		ret += fmt.Sprintf("BankType:%s, ", *o.BankType)
	}
	if o.SuccessTime != nil {
		ret += fmt.Sprintf("SuccessTime:%s, ", *o.SuccessTime)
	}
	if o.TradeState != nil {
		ret += fmt.Sprintf("TradeState:%s, ", *o.TradeState)
	}
	if o.TradeStateDesc != nil {
		ret += fmt.Sprintf("TradeStateDesc:%s, ", *o.TradeStateDesc)
	}
	if o.Payer != nil {
		ret += fmt.Sprintf("Payer:%v, ", o.Payer)
	}
	if o.Amount != nil {
		ret += fmt.Sprintf("Amount:%v, ", o.Amount)
	}
	if o.ExchangeRate != nil {
		ret += fmt.Sprintf("ExchangeRate:%v, ", o.ExchangeRate)
	}
	if len(o.PromotionDetail) > 0 {
		ret += fmt.Sprintf("PromotionDetail:%v, ", o.PromotionDetail)
	}
	return fmt.Sprintf("Transaction{%s}", ret)
}

func (o TransactionPayer) String() string {
	var ret string
	if o.Openid != nil {
		ret += fmt.Sprintf("Openid:%s, ", *o.Openid)
	}
	return fmt.Sprintf("TransactionPayer{%s}", ret)
}

func (o TransactionAmount) String() string {
	var ret string
	if o.Currency != nil {
		ret += fmt.Sprintf("Currency:%s, ", *o.Currency)
	}
	if o.PayerCurrency != nil {
		ret += fmt.Sprintf("PayerCurrency:%s, ", *o.PayerCurrency)
	}
	if o.PayerTotal != nil {
		ret += fmt.Sprintf("PayerTotal:%d, ", *o.PayerTotal)
	}
	if o.Total != nil {
		ret += fmt.Sprintf("Total:%d, ", *o.Total)
	}
	return fmt.Sprintf("TransactionAmount{%s}", ret)
}

func (o PromotionDetail) String() string {
	var ret string
	if o.PromotionId != nil {
		ret += fmt.Sprintf("PromotionId:%s, ", *o.PromotionId)
	}
	if o.Name != nil {
		ret += fmt.Sprintf("Name:%s, ", *o.Name)
	}
	if o.Scope != nil {
		ret += fmt.Sprintf("Scope:%s, ", *o.Scope)
	}
	if o.Type != nil {
		ret += fmt.Sprintf("Type:%s, ", *o.Type)
	}
	if o.Amount != nil {
		ret += fmt.Sprintf("Amount:%d, ", *o.Amount)
	}
	if o.StockId != nil {
		ret += fmt.Sprintf("StockId:%s, ", *o.StockId)
	}
	if o.WechatpayContribute != nil {
		ret += fmt.Sprintf("WechatpayContribute:%d, ", *o.WechatpayContribute)
	}
	if o.MerchantContribute != nil {
		ret += fmt.Sprintf("MerchantContribute:%d, ", *o.MerchantContribute)
	}
	if o.OtherContribute != nil {
		ret += fmt.Sprintf("OtherContribute:%d, ", *o.OtherContribute)
	}
	if o.Currency != nil {
		ret += fmt.Sprintf("Currency:%s, ", *o.Currency)
	}
	if len(o.GoodsDetail) > 0 {
		ret += fmt.Sprintf("GoodsDetail:%v, ", o.GoodsDetail)
	}
	return fmt.Sprintf("PromotionDetail{%s}", ret)
}

func (o PromotionGoodsDetail) String() string {
	var ret string
	if o.GoodsId != nil {
		ret += fmt.Sprintf("GoodsId:%s, ", *o.GoodsId)
	}
	if o.Quantity != nil {
		ret += fmt.Sprintf("Quantity:%d, ", *o.Quantity)
	}
	if o.UnitPrice != nil {
		ret += fmt.Sprintf("UnitPrice:%d, ", *o.UnitPrice)
	}
	if o.DiscountAmount != nil {
		ret += fmt.Sprintf("DiscountAmount:%d, ", *o.DiscountAmount)
	}
	if o.GoodsRemark != nil {
		ret += fmt.Sprintf("GoodsRemark:%s, ", *o.GoodsRemark)
	}
	return fmt.Sprintf("PromotionGoodsDetail{%s}", ret)
}
