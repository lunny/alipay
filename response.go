package alipay

import (
	"strconv"
)

type PlaceOrderResponse struct {
	PlaceOrderResult `json:"alipay_trade_precreate_response"`
	Sign string `json:"sign"`
}

type CloseOrderResponse struct {
	PlaceOrderResult `json:"alipay_trade_precreate_response"`
	Sign string `json:"sign"`
}

type QueryOrderResponse struct {
	PlaceOrderQueryResult `json:"alipay_trade_query_response"`
	Sign string `json:"sign"`
}

type PlaceOrderQueryResult struct {
	Code string `json:"code"` // "10000",
    Msg string `json:"msg"` // "处理成功",
    TradeNO string `json:"trade_no"` //  "2013112011001004330000121536",
    OutTradeNO string `json:"out_trade_no"` //: "6823789339978248",
    TradeStatus string `json:"trade_status"` //: "TRADE_SUCCESS",
    OpenID string `json:"open_id,omitempty"` //: "2088102122524333",
    BuyerLogonID string `json:"buyer_logon_id,omitempty"` //: "159****5620",
    TotalAmount float64 `json:"total_amount,string"` //: "88.88",
    ReceiptAmount float64 `json:"receipt_amount,string"` //: "8.88",
    SendPayDate string `json:"send_pay_date"` //: "2014-11-27 15:45:57",
    StoreID string `json:"store_id,omitempty"` //:"NJ_S_001",
    TerminalID string `json:"terminal_id,omitempty"` //:"NJ_T_001",
    FundBillList []map[string]string `json:"fund_bill_list,omitempty"` //: [
    SubCode string `json:"sub_code,omitempty"` //: "ACQ.TRADE_NOT_EXIST", ACQ.SYSTEM_ERROR，ACQ.INVALID_PARAMETER
    SubMsg string `json:"sub_msg,omitempty"` //: "交易不存在"
}

func (p *PlaceOrderQueryResult) IsSuccess() bool {
	return p.Code == "10000"
}

// 预付单返回
type PlaceOrderResult map[string]string

func (p PlaceOrderResult) QrCode() string {
	return p["qr_code"]
}

func (p PlaceOrderResult) IsSuccess() bool {
	return p["code"] == "10000"
}

// 交易结果
type TradeResult map[string]string

func (p TradeResult) IsSuccess() bool {
	return p["code"] == "10000"
}

func (n TradeResult) IsTradeSuccess() bool {
	return IsTradeSuccess(n["trade_status"])
}

func (p TradeResult) TotalFee() int64 {
	m, _ := strconv.ParseFloat(p["total_amount"], 64)
	return int64(m * 100)
}