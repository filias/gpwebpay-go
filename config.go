package gpwebpay

import (
	"errors"
	"os"
)

type Config struct {
	Currency                     string
	DepositFlag                  string
	MerchantCallbackUrl          string
	MerchantId                   string
	MerchantPrivateKey           string
	MerchantPrivateKeyPassphrase string
	GPWebpayPublicKey            string
	GPWebpayUrl                  string
}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func (config *Config) InitConfigFromEnv() {
	config.Currency = getEnv("GPWEBPAY_CURRENCY", "978") // EUR
	config.DepositFlag = getEnv("GPWEBPAY_DEPOSIT_FLAG", "1")
	config.MerchantCallbackUrl = getEnv("GPWEBPAY_MERCHANT_CALLBACK_URL", "https://localhost:5000/payment_callback")
	config.MerchantId = getEnv("GPWEBPAY_MERCHANT_ID", "")
	config.MerchantPrivateKey = getEnv("GPWEBPAY_MERCHANT_PRIVATE_KEY", "")
	config.MerchantPrivateKeyPassphrase = getEnv("GPWEBPAY_MERCHANT_PRIVATE_KEY_PASSPHRASE", "")
	config.GPWebpayPublicKey = getEnv("GPWEBPAY_PUBLIC_KEY", "")
	config.GPWebpayUrl = getEnv("GPWEBPAY_URL", "https://test.3dsecure.gpwebpay.com/pgw/order.do") // Default to test env
}

func (c *Config) validate() error {
	if c.Currency == "" {
		return errors.New("Currency is required.")
	}
	if c.DepositFlag == "" {
		return errors.New("DepositFlag is required.")
	}
	if c.MerchantCallbackUrl == "" {
		return errors.New("MerchantCallbackUrl is required.")
	}
	if c.MerchantId == "" {
		return errors.New("MerchantId is required.")
	}
	if c.MerchantPrivateKey == "" {
		return errors.New("MerchantPrivateKey is required.")
	}
	if c.MerchantPrivateKeyPassphrase == "" {
		return errors.New("MerchantPrivateKeyPassphrase is required.")
	}
	if c.GPWebpayPublicKey == "" {
		return errors.New("GPWebpayPublicKey is required.")
	}
	if c.GPWebpayUrl == "" {
		return errors.New("GPWebpayUrl is required.")
	}
	return nil
}
