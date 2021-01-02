package main

import (
	"fmt"
	"github.com/elliotcubit/parmesan"
)

func main() {
	Test()
}

func Test() {
	b := &parmesan.BittrexClient{}
	// res, err := b.GetMarkets()
	// res, err := b.GetMarketSummaries()
	// res, err := b.GetMarketTickerList()
	// res, err := b.GetMarketTicker("BTC-USD")
	// res, err := b.GetMarket("BTC-USD")
	// res, err := b.GetMarketSummary("BTC-USD")
	// res, err := b.GetMarketOrderbook("BTC-USD")
	// res, err := b.GetMarketTrades("BTC-USD")
	// res, err := b.GetRecentCandles("BTC-USD", parmesan.MinuteOne)
	// res, err := b.GetHistoricalCandles("BTC-USD", parmesan.HistoricalDayOne, 2020, 1, 1)
	// res, err := b.GetCurrencies()
	// res, err := b.GetCurrency("XMR")
	res, err := b.Ping()
	if err != nil {
		fmt.Println("fuck")
		return
	}
	fmt.Printf("%+v", res)
}
