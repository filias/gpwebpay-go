package gpwebpay

import (
	"fmt"
	"net/http"
	"net/url"
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

// This is a first shot at having some methods.
// TODO: fix declarations and return types.
func (c *GPWebpayClient) createPaymentData(orderNumber string, amount int) {}
func (c *GPWebpayClient) createMessage(data interface{}, isDigest1 bool)   {}
func (c *GPWebpayClient) signMessage(message []byte, key []byte)           {}

func (c *GPWebpayClient) createCallbackData(urlToParse string) map[string]interface{} {
	parsed, err := url.Parse(urlToParse)
	if err != nil {
		panic(err)
	}
	queryString := parsed.Query()
	m := make(map[string]interface{})
	for key, value := range queryString {
		m[""+key+""] = value
		fmt.Println("Key:", key, "Value:", value)
	}
	return m
}

func (c *GPWebpayClient) isCallbackValid()                        {}
func (c *GPWebpayClient) GetPaymentResult(url string, key []byte) {}
