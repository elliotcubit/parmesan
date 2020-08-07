package parmesan

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
