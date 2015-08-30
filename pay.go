package alipay

var (
	payChannels = map[string]string{
		"COUPON":        "支付宝红包",
		"ALIPAYACCOUNT": "支付宝余额",
		"POINT":         "积分",
		"DISCOUNT":      "折扣券",
		"PCARD":         "预付卡",
		"MCARD":         "商户店铺卡",
		"MDISCOUNT":     "商户优惠券",
		"MCOUPON":       "商户红包",
	}

	tradeStatus = map[string]string{
		"WAIT_BUYER_PAY": "等待付款",
		"TRADE_SUCCESS":  "支持成功",
		"TRADE_FINISHED": "支持完成",
		"TRADE_CLOSED":   "交易已关闭",
	}
)

func IsTradeSuccess(status string) bool {
	return status == "TRADE_SUCCESS" || status == "TRADE_FINISHED"
}
