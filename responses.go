package parmesan

import (
	"github.com/shopspring/decimal"
)

// ===== Markets =====
type Market struct {
	Symbol              string          `json:"symbol"`
	BaseCurrencySymbol  string          `json:"baseCurrencySymbol"`
	QuoteCurrencySymbol string          `json:"quoteCurrencySymbol"`
	MinTradeSize        decimal.Decimal `json:"minTradeSize"`
	Precision           int32           `json:"precision"`
	Status              string          `json:"status"`
	CreatedAt           string          `json:"createdAt"`
	Notice              string          `json:"notice"`
	ProhibitedIn        []string        `json:"prohibitedIn"`
}

type MarketList []Market

type MarketTicker struct {
	Symbol        string          `json:"symbol"`
	LastTradeRate decimal.Decimal `json:"lastTradeRate"`
	BidRate       decimal.Decimal `json:"bidRate"`
	AskRate       decimal.Decimal `json:"askRate"`
}

type MarketTickerList []MarketTicker

type MarketSummary struct {
	Symbol        string          `json:"symbol"`
	High          decimal.Decimal `json:"high"`
	Low           decimal.Decimal `json:"low"`
	Volume        decimal.Decimal `json:"volume"`
	QuoteVolume   decimal.Decimal `json:"quoteVolume"`
	PercentChange decimal.Decimal `json:"percentChange"`
	UpdatedAt     string          `json:"updatedAt"`
}

type MarketSummaryList []MarketSummary

type OrderBookEntry struct {
	Quantity decimal.Decimal `json:"quantity"`
	Rate     decimal.Decimal `json:"rate"`
}

type MarketOrderbook struct {
	Bid []OrderBookEntry `json:"bid"`
	Ask []OrderBookEntry `json:"ask"`
}

type Trade struct {
	Id         string          `json:"id"`
	ExecutedAt string          `json:"executedAt"`
	Quantity   decimal.Decimal `json:"quantity"`
	Rate       decimal.Decimal `json:"rate"`
	TakerSide  string          `json:"takerSide"`
}

type MarketTrades []Trade

// By weird design, recent intervals and historical intervals
// use the same string enum, but have different effects.. Bittrex?
const (
	MinuteOne  = "MINUTE_1"
	MinuteFive = "MINUTE_5"
	HourOne    = "HOUR_1"
	DayOne     = "DAY_1"

	HistoricalDayOne   = "MINUTE_1"
	HistoricalMonthOne = "HOUR_1"
	HistoricalYearOne  = "DAY_1"
)

type Candle struct {
	StartsAt    string          `json:"startsAt"`
	Open        decimal.Decimal `json:"open"`
	High        decimal.Decimal `json:"high"`
	Low         decimal.Decimal `json:"low"`
	Close       decimal.Decimal `json:"close"`
	Volume      decimal.Decimal `json:"volume"`
	QuoteVolume decimal.Decimal `json:"quoteVolume"`
}

type CandleList []Candle

// ===== Currencies =====

type Currency struct {
	Symbol string `json:"symbol"`
	Name string `json:"name"`
	CoinType string `json:"coinType"`
	Status string `json:"status"`
	MinConfirmations int32 `json:"minConfirmations"`
	Notice string `json:"notice"`
	TxFee decimal.Decimal `json:"txFee"`
	LogoUrl string `json:"logoUrl"`
	ProhibitedIn []string `json:"prohibitedIn"`
}

type CurrencyList []Currency


// ===== Ping =====
type PingResponse struct {
	ServerTime int64 `json:"serverTime"`
}
