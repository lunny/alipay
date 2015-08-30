package alipay

type Error struct {
	Desc string
	Reason string
	Resolved string
}

var (
	errorDescs = map[string]Error{
		"ACQ.SYSTEM_ERROR": {
			"接口返回错误",
			"系统超时",
			"请立即调用查询订单 API,查询当前订单的状态，并根据订单状态决定下一步的操作",
		},
		"ACQ.INVALID_PARAMETER": {
			"参数无效",
			"请求参数未按指引进行编写",
			"检查请求参数，修改后重新发起请求",
		},
		"ACQ.ACCESS_FORBIDDEN": {
			"无权限使用接口",
			"未签约条码支付或者合同已到期",
			"联系支付宝小二签约条码支付",
		},
		"ACQ.EXIST_FORBIDDEN_WORD": {
			"订单信息中包含违禁词",
			"订单信息中(标题，商品名称，描述等)包含了违禁词",
			"修改订单信息后，重新发起请求",
		},
		"ACQ.PARTNER_ERROR": {
			"应用 APP_ID 填写错误",
			"应用 APP_ID 填写错误或者对应的 APP_ID 状态无效",
			"联系支付宝小二，确认 APP_ID 的状态",
		},
		"ACQ.TOTAL_FEE_EXCEED": {
			"订单总金额超过限额",
			"输入的订单总金额超过限额",
			"修改订单金额再发起请求",
		},
		"ACQ.CONTEXT_INCONSISTENT": {
			"交易信息被篡改",
			"该笔交易已存在，但是 交易信息匹配不上",
			"更换商家订单号后，重新发起请求",
		},
		"ACQ.TRADE_HAS_SUCCESS": {
			"交易已被支付",
			"该笔交易已存在，并且 已经支付成功",
			"确认该笔交易信息是否为当前买家的，如果是则认为交易付款成功，如果不是则更换商家订单号后，重新发起请求",
		},
		"ACQ.TRADE_HAS_CLOSE": {
			"交易已经关闭",
			"该笔交易已存在，并且 该交易已经关闭",
			"更换商家订单号后，重新发起请求",
		},
		"ACQ.BUYER_SELLER_EQUAL": {
			"买卖家不能相同",
			"交易的买卖家为同一个人",
			"更换买家重新付款",
		},
		"ACQ.TRADE_BUYER_NOT_MATCH": {
			"交易买家不匹配",
			"该笔交易已存在，但是交易不属于当前付款的买家",
			"更换商家订单号后，重新发起请求",
		},
		"ACQ.BUYER_ENABLE_STATUS_FORBID": {
			"买家状态非法",
			"买家的状态不合法，不能进行交易",
			"用户联系支付宝小二，确认买家状态为什么非法",
		},
		"ACQ.BUYER_PAYMENT_AMOUNT_DAY_LIMIT_ERROR": {
			"买家付款日限额超限",
			"当前买家(用户)当日支付宝付款额度已用完",
			"更换买家进行支付",
		},
		"ACQ.BEYOND_PAY_RESTRICTION": {
			"商户收款额度超限",
			"商户收款额度超限",
			"联系支付宝小二提高限额",
		},
		"ACQ.BEYOND_PER_RECEIPT_RESTRICTION": {
			"商户收款金额超过月限额",
			"商户收款金额超过月限额",
			"联系支付宝小二提高限额",
		},
		"ACQ.BUYER_PAYMENT_AMOUNT_MONTH_LIMIT_ERROR": {
			"买家付款月额度超限",
			"买家本月付款额度已超限",
			"让买家更换账号后，重新付款或者更换其它付款方式",
		},
		"ACQ.SELLER_BEEN_BLOCKED": {
			"商家账号被冻结",
			"商家账号被冻结",
			"联系支付宝小二，解冻账号",
		},
		"ACQ.ERROR_BUYER_CERTIFY_LEVEL_LIMIT": {
			"买家未通过人行认证",
			"当前买家(用户)未通 过人行认证",
			"让用户联系支付宝小二并更 换其它付款方式",
		},
	}
)