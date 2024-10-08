// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

type Coin struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	Symbol               string   `json:"symbol"`
	Price                float64  `json:"price"`
	OneHourChange        *float64 `json:"oneHourChange,omitempty"`
	TwentyFourHourChange *float64 `json:"twentyFourHourChange,omitempty"`
	SevenDayChange       *float64 `json:"sevenDayChange,omitempty"`
	MarketCap            *float64 `json:"marketCap,omitempty"`
	Volume24h            *float64 `json:"volume24h,omitempty"`
}

type Query struct {
}
