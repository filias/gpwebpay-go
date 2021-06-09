package gpwebpay

import (
	"fmt"
	"testing"
)

func TestValidateNoError(t *testing.T) {
	config := Config{
		Currency:                     "test",
		DepositFlag:                  "test",
		MerchantCallbackUrl:          "test",
		MerchantId:                   "test",
		MerchantPrivateKey:           "test",
		MerchantPrivateKeyPassphrase: "test",
		GPWebpayPublicKey:            "test",
		GPWebpayUrl:                  "test",
	}
	if validationError := config.validate(); validationError != nil {
		t.Error("Config should be valid. Expected nil, got", validationError)
	}
}

func TestValidateWithError(t *testing.T) {
	tests := []Config{
		{},
		{
			Currency: "test",
		},
		{
			Currency:    "test",
			DepositFlag: "test",
		},
		{
			Currency:            "test",
			DepositFlag:         "test",
			MerchantCallbackUrl: "test",
		},
		{
			Currency:            "test",
			DepositFlag:         "test",
			MerchantCallbackUrl: "test",
			MerchantId:          "test",
		},
		{
			Currency:            "test",
			DepositFlag:         "test",
			MerchantCallbackUrl: "test",
			MerchantId:          "test",
			MerchantPrivateKey:  "test",
		},
		{
			Currency:            "test",
			DepositFlag:         "test",
			MerchantCallbackUrl: "test",
			MerchantId:          "test",
			MerchantPrivateKey:  "test",
		},
		{
			Currency:                     "test",
			DepositFlag:                  "test",
			MerchantCallbackUrl:          "test",
			MerchantId:                   "test",
			MerchantPrivateKey:           "test",
			MerchantPrivateKeyPassphrase: "test",
		},
		{
			Currency:                     "test",
			DepositFlag:                  "test",
			MerchantCallbackUrl:          "test",
			MerchantId:                   "test",
			MerchantPrivateKey:           "test",
			MerchantPrivateKeyPassphrase: "test",
			GPWebpayPublicKey:            "test",
		},
	}

	for test_num, config := range tests {
		t.Run(fmt.Sprint(test_num), func(t *testing.T) {
			if validationError := config.validate(); validationError == nil {
				t.Error("Config should be invalid as a required field is missing. Expected an error, got nil.")
			}
		})
	}
}
