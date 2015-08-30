package alipay

import "testing"

func TestParams(t *testing.T) {
	params := Params{
		"nonce_str":    "35dcf9064d9b84f971d6120f6c652ff7",
		"out_trade_no": "0123456",
		"subject":      "一元洗车",
		"bod":          "hehe",
	}

	correctParams := Params{
		"bod":          "hehe",
		"nonce_str":    "35dcf9064d9b84f971d6120f6c652ff7",
		"out_trade_no": "0123456",
		"subject":      "一元洗车",
	}

	p1 := params.Encode(false)
	p2 := correctParams.Encode(false)

	if p1 != p2 {
		t.Fatal(p1, "is not equal to", p2)
		return
	}
}
