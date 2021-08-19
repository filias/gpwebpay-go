package gpwebpay

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/url"

	pkcs8 "github.com/youmark/pkcs8"
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

func (client *GPWebpayClient) RequestPayment(orderNo string, amount int) (*http.Response, error) {
	// def request_payment(
	//     self, order_number: str = None, amount: int = 0, key: bytes = None
	// ) -> Type[requests.Response]:
	//     self._create_payment_data(order_number=order_number, amount=amount)
	//     message = self._create_message(self.data)
	//     self._sign_message(message, key=key)

	//     # Send the request
	//     # TODO: check if we need all these headers
	//     headers = {
	//         "accept-charset": "UTF-8",
	//         "accept-encoding": "UTF-8",
	//         "Content-Type": "application/x-www-form-urlencoded",
	//     }
	//     response = requests.post(
	//         configuration.GPWEBPAY_URL, data=self.data, headers=headers
	//     )

	//     return response

	// This makes an http request to gpwebpay
	url := client.config.GPWebpayUrl
	// url := "https://httpbin.rovio.com/redirect/1"
	// data := "url=http%3A%2F%2Fexample.com"

	pdata := client.createPaymentData(orderNo, amount)
	message := client.createMessage(pdata)
	digest, err := client.signMessage(message)
	if err != nil {
		return nil, err
	}
	bdata := bytes.NewBuffer([]byte(digest))

	resp, err := client.httpClient.Post(url, "application/x-www-form-urlencoded", bdata)

	if err != nil {
		return nil, fmt.Errorf("gpwebpay: Payment request failed. %s", err)
	}
	// defer resp.Body.Close()

	return resp, nil
}

// This signs the message with the key
func (c *GPWebpayClient) signMessage(message string) (string, error) {
	var emptyString string

	keyBase64 := c.config.MerchantPrivateKey
	keyPEM, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return emptyString, fmt.Errorf("gpwebpay: couldn't b64 decode the private key. %s", err)
	}

	block, rest := pem.Decode(keyPEM)
	if len(rest) != 0 && block == nil {
		return emptyString, fmt.Errorf("gpwebpay: couldn't decode PEM block. %s", string(rest))
	}

	keyPass := c.config.MerchantPrivateKeyPassphrase
	rsaPrivateKey, err := pkcs8.ParsePKCS8PrivateKeyRSA(block.Bytes, []byte(keyPass))
	if err != nil {
		return emptyString, fmt.Errorf("gpwebpay: couldn't parse the private key with password. %s", err)
	}

	hashed := sha1.Sum([]byte(message))
	signature, err := rsa.SignPKCS1v15(nil, rsaPrivateKey, crypto.SHA1, hashed[:])
	if err != nil {
		return emptyString, fmt.Errorf("gpwebpay: couldn't sign the message with a key. %s", err)
	}

	digest := base64.StdEncoding.EncodeToString(signature)
	return digest, nil
}

// This is a first shot at having some methods.
// TODO: fix declarations and return types.
func (c *GPWebpayClient) createPaymentData(orderNumber string, amount int) string {
	return ""
}
func (c *GPWebpayClient) createMessage(data string) string {
	return ""
}

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
