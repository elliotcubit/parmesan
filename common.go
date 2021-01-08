package parmesan

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const baseUrl = "https://api.bittrex.com/v3"

type BittrexClient struct {
	apiKey    string
	apiSecret string
	client    http.Client
}

func New(apiKey, apiSecret string) *BittrexClient {
	return &BittrexClient{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		client: http.Client{
			Timeout: time.Second * 30,
		},
	}
}

// TODO check for errored responses
func (b *BittrexClient) apiGet(path []string) ([]byte, error) {
	trail := strings.Join(path, "/")
	req := fmt.Sprintf("%s/%s", baseUrl, trail)
	fmt.Println(req)
	resp, err := http.Get(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Convenience methods
func (b *BittrexClient) authApiGet(path []string, getParameters map[string]interface{}) ([]byte, error) {
	return b.authApiGeneric("GET", path, getParameters)
}
func (b *BittrexClient) authApiDelete(path []string, getParameters map[string]interface{}) ([]byte, error) {
	return b.authApiGeneric("DELETE", path, getParameters)
}

func (b *BittrexClient) authApiGeneric(typ string, path []string, getParameters map[string]interface{}) ([]byte, error) {
	if b.apiSecret == "" || b.apiKey == "" {
		return nil, errors.New("You must provide credentials to use this endpoint.")
	}
	trail := strings.Join(path, "/")
	param := ""
	if len(getParameters) > 0 {
		if typ != "GET" {
			return nil, fmt.Errorf("GET parameters passed to not a HTTP GET request.")
		}
		// Trailing ampersands are supposedly legal
		param += "?="
		for key, value := range getParameters {
			var stringVal string
			switch argType := value.(type) {
			case string:
				stringVal = value.(string)
			case int:
				stringVal = fmt.Sprintf("%d", value)
			// TODO time representation varies sometimes, I think
			// Use ISO-8601
			case time.Time:
				stringVal = value.(time.Time).Format(time.RFC3339)
			case bool:
				stringVal = fmt.Sprintf("%v", value)
			default:
				return nil, fmt.Errorf("Unsupported query argument type (%v) for value (%v)", argType, value)
			}
			param += key + "=" + stringVal + "&"
		}
	}
	URL := fmt.Sprintf("%s/%s%s", baseUrl, trail, param)
	req, err := http.NewRequest(typ, URL, nil)
	if err != nil {
		return nil, err
	}

	nonce := b.currentTime()
	contentHash := b.hash([]byte{})

	preSigned := nonce + URL + typ + contentHash
	h := hmac.New(sha512.New, []byte(b.apiSecret))
	h.Write([]byte(preSigned))
	apiSignature := hex.EncodeToString(h.Sum(nil))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Api-Key", b.apiKey)
	req.Header.Set("Api-Timestamp", nonce)
	req.Header.Set("Api-Content-Hash", contentHash)
	req.Header.Set("Api-Signature", apiSignature)

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Time in unix epoch-millsecond format as a string
func (b *BittrexClient) currentTime() string {
	return strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
}

// Hex representation of a sha512 hash of data
func (b *BittrexClient) hash(data []byte) string {
	hash := sha512.Sum512(data)
	return hex.EncodeToString(hash[:])
}
