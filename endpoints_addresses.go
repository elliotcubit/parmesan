package parmesan

import (
	"encoding/json"
)

func (b *BittrexClient) GetAddresses() (result AddressList, err error) {
	response, err := b.authApiGet([]string{"addresses"}, nil)
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

func (b *BittrexClient) GetAddress(symbol string) (result Address, err error) {
	response, err := b.authApiGet([]string{"addresses", symbol}, nil)
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}
