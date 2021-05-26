package gpwebpay

import (
	"errors"
	"http"
)

type GPWebpayClient struct {
	config     Config
	httpClient *http.Client
}

func NewClient(config Config, httpClient *http.Client) (*GPWebpayClient, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// TODO: Fix checking if config has not been initialised yet.
	//       Right now it results in: "invalid operation: config == nil (mismatched types Config and nil)"
	// if config == nil {
	// 	config = InitConfigFromEnv()
	// }

	configErr := config.validate()
	if configErr != nil {
		errMsg := "Config is invalid."
		// TODO add configErr to errMsg.
		return nil, errors.New(errMsg)
	}

	gpWebpayClient := &GPWebpayClient{
		config:     config,
		httpClient: httpClient,
	}
	return gpWebpayClient, nil
}

// This is a first shot at having some methods.
// TODO: fix declarations and return types.
func (c *GPWebpayClient) createPaymentData(orderNumber string, amount int) {}
func (c *GPWebpayClient) createMessage(data interface{}, isDigest1 bool)   {}
func (c *GPWebpayClient) signMessage(message []byte, key []byte)           {}
func (c *GPWebpayClient) createCallbackData(url string)                    {}
func (c *GPWebpayClient) isCallbackValid()                                 {}
func (c *GPWebpayClient) RequestPayment()                                  {}
func (c *GPWebpayClient) GetPaymentResult(url string, key []byte)          {}
