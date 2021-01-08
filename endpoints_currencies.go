package parmesan

import (
	"encoding/json"
)

func (b *BittrexClient) GetCurrencies() (result CurrencyList, err error) {
	response, err := b.apiGet([]string{"currencies"})
	if err == nil {
		err = json.Unmarshal(response, &result)
	}
	return result, err
}

func (b *BittrexClient) GetCurrency(symbol string) (result Currency, err error) {
	response, err := b.apiGet([]string{"currencies", symbol})
	if err == nil {
		err = json.Unmarshal(response, &result)
	}
	return result, err
}
