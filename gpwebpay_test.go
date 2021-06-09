package gpwebpay

import (
	"fmt"
	"reflect"
	"testing"
)

func TestClient(t *testing.T) {
	config := Config{}
	client, _ := NewClient(config, nil)

	response, respError := client.RequestPayment()

	if respError != nil {
		fmt.Println(respError)
	} else {
		fmt.Println(response)
	}

}

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

	config := Config{}
	client, _ := NewClient(config, nil)

	signedMessage, _ := client.signMessage(message)
	if signedMessage != expectedDigest {
		t.Error("Signing message failed", signedMessage)
	}

}
