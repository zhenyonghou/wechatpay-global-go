package refund_global

import "fmt"

// 退款通知数据结构
type RefundNotifyData struct {
	Mchid         *string             `json:"mchid"`                  // 微信支付分配的商户号,仅适用于直连模式
	OutTradeNo    *string             `json:"out_trade_no"`           // 返回的商户订单号
	TransactionId *string             `json:"transaction_id"`         // 微信支付订单号
	OutRefundNo   *string             `json:"out_refund_no"`          // 商户退款单号
	RefundId      *string             `json:"refund_id"`              // 微信退款单号
	RefundStatus  *string             `json:"refund_status"`          // 退款状态, SUCCESS：退款成功, CLOSED：退款关闭, ABNORMAL：退款异常
	SuccessTime   *string             `json:"success_time,omitempty"` // 退款成功时间，当退款状态为退款成功时有返回，格式为rfc3339格式
	RecvAccount   *string             `json:"recv_account,omitempty"` // 取当前退款单的退款入账方, 示例值：招商银行信用卡0403
	FundSource    *string             `json:"fund_source,omitempty"`  // 退款资金来源
	Amount        *RefundNotifyAmount `json:"amount"`                 // 金额信息
}

func (o RefundNotifyData) String() string {
	var ret string
	if o.Mchid == nil {
		ret += "Mchid:<nil>, "
	} else {
		ret += fmt.Sprintf("Mchid:%v, ", *o.Mchid)
	}

	if o.OutTradeNo == nil {
		ret += "OutTradeNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutTradeNo:%v, ", *o.OutTradeNo)
	}

	if o.TransactionId == nil {
		ret += "TransactionId:<nil>, "
	} else {
		ret += fmt.Sprintf("TransactionId:%v, ", *o.TransactionId)
	}

	if o.OutRefundNo == nil {
		ret += "OutRefundNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutRefundNo:%v, ", *o.OutRefundNo)
	}

	if o.RefundId == nil {
		ret += "RefundId:<nil>, "
	} else {
		ret += fmt.Sprintf("RefundId:%v, ", *o.RefundId)
	}

	if o.RefundStatus == nil {
		ret += "RefundStatus:<nil>, "
	} else {
		ret += fmt.Sprintf("RefundStatus:%v, ", *o.RefundStatus)
	}

	// Optional fields (omitempty): print only when non-nil
	if o.SuccessTime != nil {
		ret += fmt.Sprintf("SuccessTime:%v, ", *o.SuccessTime)
	}

	if o.RecvAccount != nil {
		ret += fmt.Sprintf("RecvAccount:%v, ", *o.RecvAccount)
	}

	if o.FundSource != nil {
		ret += fmt.Sprintf("FundSource:%v, ", *o.FundSource)
	}

	if o.Amount == nil {
		ret += "Amount:<nil>"
	} else {
		ret += fmt.Sprintf("Amount:%v", o.Amount)
	}

	return fmt.Sprintf("RefundNotifyData{%s}", ret)
}

type RefundNotifyAmount struct {
	Total         *int64        `json:"total"`          // 订单金额
	Currency      *string       `json:"currency"`       // 订单标价币种
	Refund        *int64        `json:"refund"`         // 退款金额
	PayerTotal    *int64        `json:"payer_total"`    // 用户实际支付金额
	PayerRefund   *int64        `json:"payer_refund"`   // 退款给用户的金额，不包含优惠券金额
	PayerCurrency *string       `json:"payer_currency"` // 用户支付币种
	ExchangeRate  *ExchangeRate `json:"exchange_rate,omitempty"`
}

func (o RefundNotifyAmount) String() string {
	var ret string
	if o.Total == nil {
		ret += "Total:<nil>, "
	} else {
		ret += fmt.Sprintf("Total:%v, ", *o.Total)
	}

	if o.Currency == nil {
		ret += "Currency:<nil>, "
	} else {
		ret += fmt.Sprintf("Currency:%v, ", *o.Currency)
	}

	if o.Refund == nil {
		ret += "Refund:<nil>, "
	} else {
		ret += fmt.Sprintf("Refund:%v, ", *o.Refund)
	}

	if o.PayerTotal == nil {
		ret += "PayerTotal:<nil>, "
	} else {
		ret += fmt.Sprintf("PayerTotal:%v, ", *o.PayerTotal)
	}

	if o.PayerRefund == nil {
		ret += "PayerRefund:<nil>, "
	} else {
		ret += fmt.Sprintf("PayerRefund:%v, ", *o.PayerRefund)
	}

	if o.PayerCurrency == nil {
		ret += "PayerCurrency:<nil>, "
	} else {
		ret += fmt.Sprintf("PayerCurrency:%v, ", *o.PayerCurrency)
	}

	// Optional (omitempty): only print when non-nil
	if o.ExchangeRate != nil {
		ret += fmt.Sprintf("ExchangeRate:%v", o.ExchangeRate)
	}

	return fmt.Sprintf("RefundNotifyAmount{%s}", ret)
}
