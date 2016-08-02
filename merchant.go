package alipay

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Merchant struct {
	AppId  string
	logger *log.Logger

	privateKey   *rsa.PrivateKey
	aliPublicKey *rsa.PublicKey
}

func (m *Merchant) Sign(data []byte) (string, error) {
	return Sign(m.privateKey, data)
}

func (m *Merchant) Verify(data []byte, sig string) error {
	return Verify(m.aliPublicKey, data, sig)
}

func (m *Merchant) Error(args ...interface{}) {
	args = append([]interface{}{"[ERR]"}, args...)
	m.logger.Println(args...)
}

func (m *Merchant) Errorf(format string, args ...interface{}) {
	m.logger.Printf("[ERR]"+format, args...)
}

func (m *Merchant) Debug(args ...interface{}) {
	args = append([]interface{}{"[DBG]"}, args...)
	m.logger.Println(args...)
}

func (m *Merchant) Debugf(format string, args ...interface{}) {
	m.logger.Printf("[DBG]"+format, args...)
}

func NewMerchant(appid string, prikeyPath, aliPublicKeyPath string) (*Merchant, error) {
	priKey, err := LoadPrivateKey(prikeyPath)
	if err != nil {
		return nil, err
	}

	aliPublicKey, err := LoadPublicKey(aliPublicKeyPath)
	if err != nil {
		return nil, err
	}

	return &Merchant{
		logger:       log.New(os.Stdout, "["+appid+"]", log.LstdFlags),
		AppId:        appid,
		privateKey:   priKey,
		aliPublicKey: aliPublicKey,
	}, nil
}

func (m *Merchant) IsValid() bool {
	return m.AppId != "" && m.privateKey != nil && m.aliPublicKey != nil
}

func (m *Merchant) Notify(req *http.Request) (TradeResult, error) {
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	resp, err := m.ParseRequest(bs)
	if err != nil {
		return nil, err
	}
	return TradeResult(resp), nil
}

func (m *Merchant) ParseRequest(res []byte) (Params, error) {
	resp, err := ParseParams(string(res))
	if err != nil {
		return nil, err
	}

	sig := resp["sign"]
	delete(resp, "sign")
	delete(resp, "sign_type")

	err = Verify(m.aliPublicKey, []byte(resp.Encode(false)), sig)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *Merchant) BizRequest(url, method, notifyUrl string, bizData map[string]string) ([]byte, error) {
	bizContent, err := json.Marshal(bizData)
	if err != nil {
		return nil, err
	}

	var req = Params{
		"app_id":      m.AppId,
		"method":      method,
		"charset":     "utf-8",
		"timestamp":   time.Now().Format("2006-01-02 15:04:05"),
		"version":     "1.0",
		"biz_content": string(bizContent),
		"sign_type":   "RSA",
	}
	if len(notifyUrl) > 0 {
		req["notify_url"] = notifyUrl
	}

	sig, err := Sign(m.privateKey, []byte(req.Encode(false)))
	if err != nil {
		return nil, err
	}

	req["sign"] = sig

	return doHttpPost(gatewayUrl, []byte(req.Encode(true)))
}

// 预下单API, amount单位为分
// https://app.alipay.com/market/document.htm?name=saomazhifu#page-14
func (m *Merchant) PlaceOrder(orderId, goodsname, desc, clientIp, notifyUrl string, amount int64) ([]byte, PlaceOrderResult, error) {
	data, err := m.BizRequest(gatewayUrl, "alipay.trade.precreate", notifyUrl, map[string]string{
		"out_trade_no": orderId,                                  // 商户订单号
		"total_amount": fmt.Sprintf("%.2f", float64(amount)/100), // 总金额
		"subject":      goodsname,                                // 商品名称
	})
	if err != nil {
		return nil, nil, err
	}

	var resp PlaceOrderResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, nil, err
	}

	m.Debug("place order resp:", resp)

	enc := strings.TrimPrefix(string(data), `{"alipay_trade_precreate_response":{`)
	idx := strings.Index(enc, `},"sign":`)
	if idx == -1 {
		return nil, nil, errors.New("支付宝返回错误")
	}

	enc = "{" + enc[:idx] + "}"

	err = Verify(m.aliPublicKey, []byte(enc), resp.Sign)
	if err != nil {
		return nil, nil, err
	}

	return data, resp.PlaceOrderResult, nil
}

// 查询订单状态
// https://app.alipay.com/market/document.htm?name=saomazhifu#page-15
func (m *Merchant) QueryOrder(orderId string) ([]byte, *PlaceOrderQueryResult, error) {
	data, err := m.BizRequest(gatewayUrl, "alipay.trade.query", "", map[string]string{
		"out_trade_no": orderId, // 商户订单号
	})
	if err != nil {
		return nil, nil, err
	}

	var resp QueryOrderResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, nil, err
	}

	enc := strings.TrimPrefix(string(data), `{"alipay_trade_query_response":{`)
	idx := strings.Index(enc, `},"sign":`)
	if idx == -1 {
		return nil, nil, errors.New("支付宝返回错误")
	}

	enc = "{" + enc[:idx] + "}"

	//s := strings.Replace(string(enc), `/`, `\/`, -1)
	err = Verify(m.aliPublicKey, []byte(enc), resp.Sign)
	if err != nil {
		return nil, nil, err
	}

	return data, &resp.PlaceOrderQueryResult, nil
}

// 撤销订单
// https://app.alipay.com/market/document.htm?name=saomazhifu#page-16
func (m *Merchant) CloseOrder(orderId string) error {
	data, err := m.BizRequest(gatewayUrl, "alipay.trade.cancel", "", map[string]string{
		"out_trade_no": orderId, // 商户订单号
	})
	if err != nil {
		return err
	}

	var resp CloseOrderResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	enc := strings.TrimPrefix(string(data), `{"alipay_trade_cancel_response":{`)
	idx := strings.Index(enc, `},"sign":`)
	if idx == -1 {
		return errors.New("支付宝返回错误")
	}

	enc = "{" + enc[:idx] + "}"

	err = Verify(m.aliPublicKey, []byte(enc), resp.Sign)
	if err != nil {
		return err
	}

	return nil
}
