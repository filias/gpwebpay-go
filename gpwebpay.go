package gpwebpay

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
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

func (c *GPWebpayClient) signMessage(message string) (string, error) {
	// This signs the message with the key
	hasher := hmac.New(sha1.New, []byte(c.config.MerchantPrivateKey))
	hasher.Write([]byte(message))
	signature := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	return signature, nil
}

// This is a first shot at having some methods.
// TODO: fix declarations and return types.
func (c *GPWebpayClient) createPaymentData(orderNumber string, amount int) {}
func (c *GPWebpayClient) createMessage(data interface{}, isDigest1 bool)   {}

// def _create_callback_data(self, url: str) -> dict:
// 		# All the data is in the querystring
// 		parsed = urlparse.urlparse(url)
// 		query_string = parse_qs(parsed.query)
// 		data = {key: value[0] for key, value in query_string.items()}
// 		return data
func (c *GPWebpayClient) createCallbackData(urlToParse string) map[string]string {
	parsed, err := url.Parse(urlToParse)
	if err != nil {
		panic(err)
	}
	queryString := parsed.Query()
	m := make(map[string]string)
	for key, value := range queryString {
		m[key] = value[0]
		fmt.Println("Key:", key, "Value:", value)
	}
	return m
}

func (c *GPWebpayClient) isCallbackValid()                        {}
func (c *GPWebpayClient) GetPaymentResult(url string, key []byte) {}
