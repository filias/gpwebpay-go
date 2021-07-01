package gpwebpay

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"net/http"
)

type GPWebpayClient struct {
	config     Config
	httpClient *http.Client
}

func NewClient(config Config, httpClient *http.Client) (*GPWebpayClient, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// Init the config
	config.InitConfigFromEnv()

	// Create the client
	gpWebpayClient := &GPWebpayClient{
		config:     config,
		httpClient: httpClient,
	}

	return gpWebpayClient, nil
}

func (client *GPWebpayClient) RequestPayment() (*http.Response, error) {
	// This makes an http request to gpwebpay
	resp, err := http.Post(client.config.GPWebpayUrl, "application/x-www-form-urlencoded", nil)

	if err != nil {
		// handle error
	}

	return resp, nil
}

func (c *GPWebpayClient) signMessage(message string) (string, error) {
	// This signs the message with the key
	keyBase64 := "secret"
	keyPEM, _ := base64.StdEncoding.DecodeString(keyBase64)

	block, _ := pem.Decode(keyPEM)

	x509.DecryptPEMBlock(block, "secret")
	rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}

	hashed := sha1.Sum([]byte(message))

	signature, err := rsa.SignPKCS1v15(nil, rsaPrivateKey, crypto.SHA1, hashed[:])
	if err != nil {
		fmt.Println(err)
	}

	digest := base64.StdEncoding.EncodeToString(signature)
	return digest, nil
}

// This is a first shot at having some methods.
// TODO: fix declarations and return types.
func (c *GPWebpayClient) createPaymentData(orderNumber string, amount int) {}
func (c *GPWebpayClient) createMessage(data interface{}, isDigest1 bool)   {}
func (c *GPWebpayClient) createCallbackData(url string)                    {}
func (c *GPWebpayClient) isCallbackValid()                                 {}
func (c *GPWebpayClient) GetPaymentResult(url string, key []byte)          {}
