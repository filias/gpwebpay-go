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
