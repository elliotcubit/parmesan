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
	Symbol           string          `json:"symbol"`
	Name             string          `json:"name"`
	CoinType         string          `json:"coinType"`
	Status           string          `json:"status"`
	MinConfirmations int32           `json:"minConfirmations"`
	Notice           string          `json:"notice"`
	TxFee            decimal.Decimal `json:"txFee"`
	LogoUrl          string          `json:"logoUrl"`
	ProhibitedIn     []string        `json:"prohibitedIn"`
}

type CurrencyList []Currency

// ===== Ping =====
type PingResponse struct {
	ServerTime int64 `json:"serverTime"`
}

// ===== Account =====
type Account struct {
	SubaccountId string `json:"subaccountId"`
	AccountId    string `json:"accountId"`
}

type AccountVolume struct {
	Updated      string          `json:"updated"`
	Volume30Days decimal.Decimal `json:"volume30days"`
}

// ===== Addresses =====
type Address struct {
	Status           string `json:"status"`
	CurrencySymbol   string `json:"currencySymbol"`
	CryptoAddress    string `json:"cryptoAddress"`
	CryptoAddressTag string `json:"cryptoAddressTag"`
}

type AddressList []Address

// ===== Balances =====
type Balance struct {
	CurrencySymbol string          `json:"currencySymbol"`
	Total          decimal.Decimal `json:"total"`
	Available      decimal.Decimal `json:"available"`
	UpdatedAt      string          `json:"updatedAt"`
}

type BalanceList []Balance

// ===== Order Enums =====
const (
	OPERAND_LTE                         = "LTE"
	OPERAND_GTE                         = "GTE"
	STATUS_OPEN                         = "OPEN"
	STATUS_COMPLETED                    = "COMPLETED"
	STATUS_CANCELLED                    = "CANCELLED"
	STATUS_FAILED                       = "FAILED"
	DIRECTION_BUY                       = "BUY"
	DIRECTION_SELL                      = "SELL"
	TYPE_LIMIT                          = "LIMIT"
	TYPE_MARKET                         = "MARKET"
	TYPE_CEILING_LIMIT                  = "CEILING_LIMIT"
	TYPE_CEILING_MARKET                 = "CEILING_MARKET"
	PERIOD_GOOD_TILL_CANCELLED          = "GOOD_TILL_CANCELLED"
	PERIOD_IMMEDIATE_OR_CANCEL          = "IMMEDIATE_OR_CANCEL"
	PERIOD_FILL_OR_KILL                 = "FILL_OR_KILL"
	PERIOD_POST_ONLY_GOOD_TIL_CANCELLED = "POST_ONLY_GOOD_TIL_CANCELLED"
	PERIOD_BUY_NOW                      = "BUY_NOW"
)

// ===== Conditional Orders =====
type ConditionalOrder struct {
	Id                       string                    `json:"id"`
	MarketSymbol             string                    `json:"marketSymbol"`
	Operand                  string                    `json:"operand"`
	TriggerPrice             decimal.Decimal           `json:"triggerPrice"`
	TrailingStopPercent      decimal.Decimal           `json:"trailingStopPercent"`
	CreatedOrderId           string                    `json:"createdOrderId"`
	OrderToCreate            NewOrder                  `json:"orderToCreate"`
	OrderToCancel            NewCancelConditionalOrder `json:"newCancelConditionalOrder"`
	ClientConditionalOrderId string                    `json:"clientConditionalOrderId"`
	Status                   string                    `json:"status"`
	OrderCreationErrorCode   string                    `json:"orderCreationErrorCode"`
	CreatedAt                string                    `json:"createdAt"`
	UpdatedAt                string                    `json:"updatedAt"`
	ClosedAt                 string                    `json:"closedAt"`
}

type ConditionalOrderList []ConditionalOrder

type NewCancelConditionalOrder struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// ===== Orders =====
type Order struct {
	Id            string                    `json:"id"`
	MarketSymbol  string                    `json:"marketSymbol"`
	Direction     string                    `json:"direction"`
	Type          string                    `json:"type"`
	Quantity      decimal.Decimal           `json:"quantity"`
	Limit         decimal.Decimal           `json:"limit"`
	Ceiling       decimal.Decimal           `json:"ceiling"`
	TimeInForce   string                    `json:"timeInForce"`
	ClientOrderId string                    `json:"clientOrderId"`
	FillQuantity  decimal.Decimal           `json:"fillQuantity"`
	Commission    decimal.Decimal           `json:"commission"`
	Proceeds      decimal.Decimal           `json:"proceeds"`
	Status        string                    `json:"status"`
	CreatedAt     string                    `json:"createdAt"`
	UpdatedAt     string                    `json:"updatedAt"`
	ClosedAt      string                    `json:"closedAt"`
	OrderToCancel NewCancelConditionalOrder `json:"orderToCancel"`
}

type OrderList []Order

type BulkCancelResult []BulkCancelItem

type BulkCancelItem struct {
	Id         string `json:"id"`
	StatusCode string `json:"statusCode"`
	Result     Order
}

type ExecutionList []Execution

type Execution struct {
	Id           string          `json:"id"`
	MarketSymbol string          `json:"marketSymbol"`
	ExecutedAt   string          `json:"executedAt"`
	Quantity     decimal.Decimal `json:"quantity"`
	Rate         decimal.Decimal `json:"rate"`
	OrderId      string          `json:"orderId"`
	Commission   decimal.Decimal `json:"commission"`
	IsTaker      bool            `json:"isTaker"`
}

type NewOrder struct {
	MarketSymbol  string          `json:"marketSymbol"`
	Direction     string          `json:"direction"`
	Type          string          `json:"type"`
	Quantity      decimal.Decimal `json:"quantity"`
	Ceiling       decimal.Decimal `json:"ceiling"`
	Limit         decimal.Decimal `json:"limit"`
	TimeInForce   string          `json:"timeInForce"`
	ClientOrderId string          `json:"clientOrderId"`
	UseAwards     bool            `json:"useAwards"`
}

