package jsapi

import (
	"context"
	"fmt"
	"github.com/zhenyonghou/wechatpay-global-go/core"
	"github.com/zhenyonghou/wechatpay-global-go/utils"
	"strconv"
	"time"
)

type PrepayWithRequestPaymentResponse struct {
	// 预支付交易会话标识
	PrepayId *string `json:"prepay_id"` // revive:disable-line:var-naming
	// 应用ID
	Appid *string `json:"appid"`
	// 时间戳
	TimeStamp *string `json:"timeStamp"`
	// 随机字符串
	NonceStr *string `json:"nonceStr"`
	// 订单详情扩展字符串
	Package *string `json:"package"`
	// 签名方式
	SignType *string `json:"signType"`
	// 签名
	PaySign *string `json:"paySign"`
}

func (a *JsapiApiService) PrepayWithRequestPayment(
	ctx context.Context,
	req PrepayRequest,
	requestPaymentAppid string,
) (resp *PrepayWithRequestPaymentResponse, result *core.APIResult, err error) {
	prepayResp, result, err := a.Prepay(ctx, req)
	if err != nil {
		return nil, result, err
	}

	resp = new(PrepayWithRequestPaymentResponse)
	resp.PrepayId = prepayResp.PrepayId
	resp.SignType = core.String("RSA")
	resp.Appid = &requestPaymentAppid
	resp.TimeStamp = core.String(strconv.FormatInt(time.Now().Unix(), 10))
	nonce, err := utils.GenerateNonce()
	if err != nil {
		return nil, nil, fmt.Errorf("generate request for payment err:%s", err.Error())
	}
	resp.NonceStr = core.String(nonce)
	resp.Package = core.String("prepay_id=" + *prepayResp.PrepayId)
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n", *resp.Appid, *resp.TimeStamp, *resp.NonceStr, *resp.Package)
	signatureResult, err := a.Client.Sign(ctx, message)
	if err != nil {
		return nil, nil, fmt.Errorf("generate sign for payment err:%s", err.Error())
	}
	resp.PaySign = core.String(signatureResult.Signature)
	return resp, result, nil
}
