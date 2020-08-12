package parmesan

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"strconv"
	"crypto/sha512"
	"crypto/hmac"
	"encoding/hex"
)

const baseUrl = "https://api.bittrex.com/v3"

type BittrexClient struct {
	apiKey    string
	apiSecret string
	client    http.Client
}

func New(apiKey, apiSecret string) (*BittrexClient) {
	return &BittrexClient{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		client:    http.Client{
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

func (b *BittrexClient) authApiGet(path []string) ([]byte, error) {
	trail := strings.Join(path, "/")
	URL := fmt.Sprintf("%s/%s", baseUrl, trail)
	reqType := "GET"
	req, err := http.NewRequest(reqType, URL, nil)
	if err != nil {
		return nil, err
	}

	nonce := b.currentTime()
	contentHash := b.hash([]byte{})

	preSigned := nonce+URL+reqType+contentHash
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
	return strconv.FormatInt(time.Now().UnixNano() / 1e6, 10)
}

// Hex representation of a sha512 hash of data
func (b *BittrexClient) hash(data []byte) string {
	hash := sha512.Sum512(data)
	return hex.EncodeToString(hash[:])
}
