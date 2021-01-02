package parmesan

import (
	"encoding/json"
	"time"
)

/*
	Returns the ConditionalOrder that was created
*/
func (b *BittrexClient) GetConditionalOrder(id string) (result ConditionalOrder, err error) {
	response, err := b.authApiGet([]string{"conditional-orders", id}, nil)
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return
}

/*
	Returns the ConditionalOrder that was deleted
*/
func (b *BittrexClient) DeleteConditionalOrder(id string) (result ConditionalOrder, err error) {
	response, err := b.authApiDelete([]string{"conditional-orders", id}, nil)
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return
}

/*
	Get closed orders
	Zero values or nil pointers are not passed along
	All arguments are optional
*/
func (b *BittrexClient) GetClosedConditionalOrders(
	marketSymbol string,
	nextPageToken string,
	previousPageToken string,
	pageSize int,
	startDate *time.Time,
	endDate *time.Time,
) (result ConditionalOrderList, err error) {
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
	response, err := b.authApiGet([]string{"conditional-orders", "closed"}, queryArgs)
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return
}

func (b *BittrexClient) GetOpenConditionalOrders(marketSymbol string) (result ConditionalOrderList, err error) {
	queryArgs := make(map[string]interface{}, 0)
	if marketSymbol != "" {
		queryArgs["marketSymbol"] = marketSymbol
	}
	response, err := b.authApiGet([]string{"conditional-orders", "open"}, queryArgs)
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return
}
