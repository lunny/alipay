package alipay

import (
	"strconv"
)

type PlaceOrderResponse struct {
	PlaceOrderResult `json:"alipay_trade_precreate_response"`
	Sign             string `json:"sign"`
}

type CloseOrderResponse struct {
	PlaceOrderResult `json:"alipay_trade_precreate_response"`
	Sign             string `json:"sign"`
}

type QueryOrderResponse struct {
	PlaceOrderQueryResult `json:"alipay_trade_query_response"`
	Sign                  string `json:"sign"`
}

type PlaceOrderQueryResult struct {
	Code          string              `json:"code"`                     // "10000",
	Msg           string              `json:"msg"`                      // "处理成功",
	TradeNO       string              `json:"trade_no"`                 //  "2013112011001004330000121536",
	OutTradeNO    string              `json:"out_trade_no"`             //: "6823789339978248",
	TradeStatus   string              `json:"trade_status"`             //: "TRADE_SUCCESS",
	OpenID        string              `json:"open_id,omitempty"`        //: "2088102122524333",
	BuyerLogonID  string              `json:"buyer_logon_id,omitempty"` //: "159****5620",
	TotalAmount   float64             `json:"total_amount,string"`      //: "88.88",
	ReceiptAmount float64             `json:"receipt_amount,string"`    //: "8.88",
	SendPayDate   string              `json:"send_pay_date"`            //: "2014-11-27 15:45:57",
	StoreID       string              `json:"store_id,omitempty"`       //:"NJ_S_001",
	TerminalID    string              `json:"terminal_id,omitempty"`    //:"NJ_T_001",
	FundBillList  []map[string]string `json:"fund_bill_list,omitempty"` //: [
	SubCode       string              `json:"sub_code,omitempty"`       //: "ACQ.TRADE_NOT_EXIST", ACQ.SYSTEM_ERROR，ACQ.INVALID_PARAMETER
	SubMsg        string              `json:"sub_msg,omitempty"`        //: "交易不存在"
}

func (p *PlaceOrderQueryResult) IsSuccess() bool {
	return p.TradeStatus == "TRADE_SUCCESS" || p.Code == "10000"
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
// map[seller_id:2088021244059960 trade_status:TRADE_SUCCESS gmt_payment:2016-01-10 14:15:18 point_amount:0.00 trade_no:2016011021001004500071808018 invoice_amount:1.00 notify_type:trade_status_sync open_id:20881036880046219982469381819650 receipt_amount:1.00 buyer_logon_id:xia***@gmail.com buyer_pay_amount:1.00 subject:汽车洗车-好车店 gmt_create:2016-01-10 14:14:57 seller_email:op@lovechebang.com notify_id:c5be78fb74a7d7f2336582597678a5djuw fund_bill_list:[{"amount":"1.00","fundChannel":"ALIPAYACCOUNT"}] notify_time:2016-01-10 14:15:18 buyer_id:2088002359340503 app_id:2015081700218350 total_amount:1.00 out_trade_no:M14524064880000000000003]
type TradeResult map[string]string

func (p TradeResult) IsSuccess() bool {
	return p["trade_status"] == "TRADE_SUCCESS"
}

func (n TradeResult) IsTradeSuccess() bool {
	return IsTradeSuccess(n["trade_status"])
}

func (p TradeResult) TotalFee() int64 {
	m, _ := strconv.ParseFloat(p["total_amount"], 64)
	return int64(m * 100)
}
