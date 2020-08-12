package parmesan

import (
	"encoding/json"
	"fmt"
)

func (b *BittrexClient) GetMarkets() (result MarketList, err error) {
	response, err := b.apiGet([]string{"markets"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

func (b *BittrexClient) GetMarketSummaryList() (result MarketSummaryList, err error) {
	response, err := b.apiGet([]string{"markets", "summaries"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

func (b *BittrexClient) GetMarketTickerList() (result MarketTickerList, err error) {
	response, err := b.apiGet([]string{"markets", "tickers"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

func (b *BittrexClient) GetMarketTicker(symbol string) (result MarketTicker, err error) {
	response, err := b.apiGet([]string{"markets", symbol, "ticker"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

func (b *BittrexClient) GetMarket(symbol string) (result Market, err error) {
	response, err := b.apiGet([]string{"markets", symbol})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

func (b *BittrexClient) GetMarketSummary(symbol string) (result MarketSummary, err error) {
	response, err := b.apiGet([]string{"markets", symbol, "summary"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

// TODO add depth parameter
func (b *BittrexClient) GetMarketOrderbook(symbol string) (result MarketOrderbook, err error) {
	response, err := b.apiGet([]string{"markets", symbol, "orderbook"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

func (b *BittrexClient) GetMarketTrades(symbol string) (result MarketTrades, err error) {
	response, err := b.apiGet([]string{"markets", symbol, "trades"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

func (b *BittrexClient) GetRecentCandles(symbol, period string) (result CandleList, err error) {
	response, err := b.apiGet([]string{"markets", symbol, "candles", period, "recent"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

// TODO validate date specifications & throw our own nice error here
// Date fields are expected to be 1-indexed.
func (b *BittrexClient) GetHistoricalCandles(symbol, period string, year, month, day int) (result CandleList, err error) {
	start := fmt.Sprintf("%d/%d/%d", year, month, day)
	response, err := b.apiGet([]string{"markets", symbol, "candles", period, "historical", start})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}
