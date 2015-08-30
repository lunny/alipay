package alipay

import (
	"fmt"
	"testing"
)

var (
	toSigned = "test"
)

func TestPrivateKeySign(t *testing.T) {
	pk, err := LoadPrivateKey("./rsa_private_key.pem")
	if err != nil {
		t.Fatal(err)
		return
	}

	signed, err := Sign(pk, []byte(toSigned))
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println("signed:", string(signed))

	pubKey, err := LoadPublicKey("./rsa_public_key.pem")
	if err != nil {
		t.Fatal(err)
		return
	}

	err = Verify(pubKey, []byte(toSigned), signed)
	if err != nil {
		t.Fatal(err)
		return
	}
}
