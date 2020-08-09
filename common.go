package parmesan

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"strconv"
	"crypto/sha512"
	"encoding/hex"
)

const baseUrl = "https://api.bittrex.com/v3"

// TODO check for errored responses
// TODO parameters
func doApiGet(path []string) ([]byte, error) {
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

// Time in unix epoch-millsecond format as a string
func currentTime() string {
	return strconv.FormatInt(time.Now().UnixNano() / 1e6, 10)
}

// Hex representation of a sha512 hash of data
func hash(data []byte) string {
	hash := sha512.Sum512(data)
	return hex.EncodeToString(hash[:])
}
