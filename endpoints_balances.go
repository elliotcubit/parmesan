package parmesan

import (
	"encoding/json"
)

func (b *BittrexClient) GetBalances() (result BalanceList, err error) {
	response, err := b.authApiGet([]string{"balances"}, nil)
	if err == nil {
		err = json.Unmarshal(response, &result)
	}
	return result, err
}

func (b *BittrexClient) GetBalance(symbol string) (result Balance, err error) {
	response, err := b.authApiGet([]string{"balances", symbol}, nil)
	if err == nil {
		err = json.Unmarshal(response, &result)
	}
	return result, err
}
