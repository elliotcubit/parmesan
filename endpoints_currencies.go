package parmesan

import (
	"encoding/json"
)

func GetCurrencies() (result CurrencyList, err error) {
  response, err := doApiGet([]string{"currencies"})
  if err == nil {
    json.Unmarshal(response, &result)
  }
  return result, err
}

func GetCurrency(symbol string) (result Currency, err error) {
  response, err := doApiGet([]string{"currencies", symbol})
  if err == nil {
    json.Unmarshal(response, &result)
  }
  return result, err
}
