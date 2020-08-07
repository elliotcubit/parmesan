package parmesan

import (
	"encoding/json"
	"fmt"
)

func GetMarkets() (MarketList, error) {
	var result MarketList
	response, err := doApiGet([]string{"markets"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

func GetMarketSummaryList() (MarketSummaryList, error) {
	var result MarketSummaryList
	response, err := doApiGet([]string{"markets", "summaries"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

func GetMarketTickerList() (MarketTickerList, error) {
	var result MarketTickerList
	response, err := doApiGet([]string{"markets", "tickers"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

func GetMarketTicker(symbol string) (MarketTicker, error) {
	var result MarketTicker
	response, err := doApiGet([]string{"markets", symbol, "ticker"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

func GetMarket(symbol string) (Market, error) {
	var result Market
	response, err := doApiGet([]string{"markets", symbol})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

func GetMarketSummary(symbol string) (MarketSummary, error) {
	var result MarketSummary
	response, err := doApiGet([]string{"markets", symbol, "summary"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

// TODO add depth parameter
func GetMarketOrderbook(symbol string) (MarketOrderbook, error) {
	var result MarketOrderbook
	response, err := doApiGet([]string{"markets", symbol, "orderbook"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

func GetMarketTrades(symbol string) (MarketTrades, error) {
	var result MarketTrades
	response, err := doApiGet([]string{"markets", symbol, "trades"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

func GetRecentCandles(symbol, period string) (CandleList, error) {
	var result CandleList
	response, err := doApiGet([]string{"markets", symbol, "candles", period, "recent"})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}

// TODO validate date specifications & throw our own nice error here
// Date fields are expected to be 1-indexed.
func GetHistoricalCandles(symbol, period string, year, month, day int) (CandleList, error) {
	var result CandleList
	start := fmt.Sprintf("%d/%d/%d", year, month, day)
	response, err := doApiGet([]string{"markets", symbol, "candles", period, "historical", start})
	if err == nil {
		json.Unmarshal(response, &result)
	}
	return result, err
}
