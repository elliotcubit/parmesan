package parmesan

import (
	"encoding/json"
)

func (b *BittrexClient) Ping() (result PingResponse, err error) {
	response, err := b.apiGet([]string{"ping"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}
