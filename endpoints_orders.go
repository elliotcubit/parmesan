package parmesan

import (
  "encoding/json"
  "time"
)

func (b *BittrexClient) GetClosedOrders(
  marketSymbol,
  nextPageToken,
  previousPageToken string,
  pageSize int,
  startDate,
  endDate *time.Time,
  ) (result OrderList, err error) {
  queryArgs := make(map[string]interface{}, 0)
  if marketSymbol != "" {
		queryArgs["marketSymbol"] = marketSymbol
	}
	if nextPageToken != "" {
		queryArgs["nextPageToken"] = nextPageToken
	}
	if previousPageToken != "" {
		queryArgs["previousPageToken"] = previousPageToken
	}
	if pageSize != 0 {
		queryArgs["pageSize"] = pageSize
	}
	if startDate != nil {
		queryArgs["startDate"] = startDate
	}
	if endDate != nil {
		queryArgs["endDate"] = endDate
	}
  response, err := b.authApiGet([]string{"orders", "closed"}, queryArgs)
  if err == nil {
    err = json.Unmarshal(response, &result)
  }
  return
}

func (b *BittrexClient) GetOpenOrders() (result OrderList, err error) {
  response, err := b.authApiGet([]string{"orders", "open"}, nil)
  if err == nil {
    err = json.Unmarshal(response, &result)
  }
  return
}

func (b *BittrexClient) DeleteOpenOrders(marketSymbol string) (result BulkCancelResult, err error){
  response, err := b.authApiDelete([]string{"orders", "open"}, nil)
  if err == nil {
    err = json.Unmarshal(response, &result)
  }
  return
}

func (b *BittrexClient) GetOrder(id string) (result Order, err error) {
  response, err := b.authApiGet([]string{"orders", id}, nil)
  if err == nil {
    err = json.Unmarshal(response, &result)
  }
  return
}

func (b *BittrexClient) DeleteOrder(id string) (result Order, err error) {
  response, err := b.authApiDelete([]string{"orders", id}, nil)
  if err == nil {
    err = json.Unmarshal(response, &result)
  }
  return
}

func (b *BittrexClient) GetExecutionsForOrder(id string) (result ExecutionList, err error) {
  response, err := b.authApiGet([]string{"orders", id, "executions"}, nil)
  if err == nil {
    err = json.Unmarshal(response, &result)
  }
  return
}