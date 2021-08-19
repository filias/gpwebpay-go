package gpwebpay

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestCreateCallbackData(t *testing.T) {
	config := Config{}
	client, _ := NewClient(config, nil)

	callback := client.createCallbackData("https://example.org/?a=1&a=2&b=3")
	fmt.Println(callback)
	expected := map[string]string{
		"a": "1",
		"b": "3",
	}

	if !reflect.DeepEqual(callback, expected) {
		t.Errorf("expected %q but got %q", expected, callback)
	}
}

func TestSignature(t *testing.T) {
	message := "1234567890|CREATE_ORDER|123456|10|978|1|https://localhost:5000/payment_callback"
	expectedDigest := "DWarvfXJP5CFFvn8zNEtImumad7Cmj/M5qQrbcFd66bjhFR4NxkEj4WSR4sCG/6YBWQAgJ3H/n7XPCRnTu670GaivWQ0dg7DevzyZKcCJwFs4olcA2mb4vfM0yAFW0jkqD3G3eCpHylWogxCVCXrMso8WIpc5nliwq1Sp/53Q3weXAYXIwvgOe/qtVqhdeOa+r5RNaYcgKzAWafSf9bAfweoedq1yMGfXRPTyLIQfwAhsk8DTN9ybohw4mQeZ2/LFcJklMdUuLKqJ/5MLwyV9/0jmxf2bZvymr4aj3S/wpLCJnZV5HDXqYXaVPokOwvnvGXwSMNw45h1zIwIXpQhig=="

	key := os.Getenv("GPWEBPAY_MERCHANT_PRIVATE_KEY")
	pass := os.Getenv("GPWEBPAY_MERCHANT_PRIVATE_KEY_PASSPHRASE")
	if key == "" {
		t.Fatalf("No private key in the environment. Set GPWEBPAY_MERCHANT_PRIVATE_KEY")
	}
	if pass == "" {
		t.Fatalf("No private key passphrase in the environment. Set GPWEBPAY_MERCHANT_PRIVATE_KEY_PASSPHRASE")
	}
	config := Config{
		MerchantPrivateKey:           key,
		MerchantPrivateKeyPassphrase: pass,
	}
	client, _ := NewClient(config, nil)

	signedMessage, _ := client.signMessage(message)
	if signedMessage != expectedDigest {
		t.Error("Signing message failed", signedMessage)
	}
}

func TestRequestPayment(t *testing.T) {

	key := os.Getenv("GPWEBPAY_MERCHANT_PRIVATE_KEY")
	pass := os.Getenv("GPWEBPAY_MERCHANT_PRIVATE_KEY_PASSPHRASE")
	if key == "" {
		t.Fatalf("No private key in the environment. Set GPWEBPAY_MERCHANT_PRIVATE_KEY")
	}
	if pass == "" {
		t.Fatalf("No private key passphrase in the environment. Set GPWEBPAY_MERCHANT_PRIVATE_KEY_PASSPHRASE")
	}
	config := Config{
		MerchantPrivateKey:           key,
		MerchantPrivateKeyPassphrase: pass,
		// TODO add URL
	}

	client, _ := NewClient(config, nil)

	resp, err := client.RequestPayment("foobar", 300)
	if err != nil {
		t.Error("paymnet request test failed", err)
	}
	fmt.Println("======= URL ======= ", resp.Request.URL)
	loc, err := resp.Location()
	fmt.Println("===== Location ==== ", loc, err)
	fmt.Println("=== status code ===", resp.StatusCode)
	t.FailNow()
}
