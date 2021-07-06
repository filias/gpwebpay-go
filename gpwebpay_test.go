package gpwebpay

import (
	"fmt"
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
	expected = map[string]interface{
		"a":[1,2],
		"b":3,
	}

	if callback != expected {
		t.Errorf("expected %q but got %q", expected, callback)
	}
}
