package models

type Coin struct {
	ID                   string  `json:"id"`
	Name                 string  `json:"name"`
	Symbol               string  `json:"symbol"`
	Price                float64 `json:"current_price"`
	OneHourChange        float64 `json:"price_change_percentage_1h_in_currency"`
	TwentyFourHourChange float64 `json:"price_change_percentage_24h_in_currency"`
	SevenDayChange       float64 `json:"price_change_percentage_7d_in_currency"`
	MarketCap            float64 `json:"market_cap"`
	Volume24h            float64 `json:"total_volume"`
}
