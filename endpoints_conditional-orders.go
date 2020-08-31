package parmesan

import (
	"encoding/json"
)

func (b *BittrexClient) GetConditionalOrder(id string) (result ConditionalOrder, err error) {
  response, err := b.authApiGet([]string{"conditional-orders", id})
  if err == nil {
    json.Unmarshal(response, &result)
  }
  return result, err
}

// TODO delete order
