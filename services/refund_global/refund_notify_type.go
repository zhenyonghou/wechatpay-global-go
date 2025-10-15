package refund_global

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

type RefundNotifyAmount struct {
	Total         *int64        `json:"total"`          // 订单金额
	Currency      *string       `json:"currency"`       // 订单标价币种
	Refund        *int64        `json:"refund"`         // 退款金额
	PayerTotal    *int64        `json:"payer_total"`    // 用户实际支付金额
	PayerRefund   *int64        `json:"payer_refund"`   // 退款给用户的金额，不包含优惠券金额
	PayerCurrency *string       `json:"payer_currency"` // 用户支付币种
	ExchangeRate  *ExchangeRate `json:"exchange_rate,omitempty"`
}
