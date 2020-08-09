package main

import (
	"fmt"
	"github.com/elliotcubit/parmesan"
)

func main() {
	Test()
}

func Test() {
	// res, err := parmesan.GetMarkets()
	// res, err := parmesan.GetMarketSummaries()
	// res, err := parmesan.GetMarketTickerList()
	// res, err := parmesan.GetMarketTicker("BTC-USD")
	// res, err := parmesan.GetMarket("BTC-USD")
	// res, err := parmesan.GetMarketSummary("BTC-USD")
	// res, err := parmesan.GetMarketOrderbook("BTC-USD")
	// res, err := parmesan.GetMarketTrades("BTC-USD")
	// res, err := parmesan.GetRecentCandles("BTC-USD", parmesan.MinuteOne)
	// res, err := parmesan.GetHistoricalCandles("BTC-USD", parmesan.HistoricalDayOne, 2020, 1, 1)
	// res, err := parmesan.GetCurrencies()
	// res, err := parmesan.GetCurrency("XMR")
	res, err := parmesan.Ping()
	if err != nil {
		fmt.Println("fuck")
		return
	}
	fmt.Printf("%+v", res)
}
