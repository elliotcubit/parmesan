package parmesan

import (
	"encoding/json"
)

func (b *BittrexClient) GetAccount() (result Account, err error) {
	response, err := b.authApiGet([]string{"account"}, nil)
	if err == nil {
		err = json.Unmarshal(response, &result)
	}
	return result, err
}

func (b *BittrexClient) GetAccountVolume() (result AccountVolume, err error) {
	response, err := b.authApiGet([]string{"account", "volume"}, nil)
	if err == nil {
		err = json.Unmarshal(response, &result)
	}
	return result, err
}
