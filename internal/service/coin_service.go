package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "go-crypto-tracker/pkg/models"
)

type CoinService struct{}

const CoinGeckoAPIURL = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=10&page=1&sparkline=false"

func (s *CoinService) GetCoins() ([]*models.Coin, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=10&page=1&sparkline=false&price_change_percentage=1h,24h,7d")
	if err != nil {
		return nil, fmt.Errorf("error fetching data from CoinGecko API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	var coinsData []struct {
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

	if err := json.NewDecoder(resp.Body).Decode(&coinsData); err != nil {
		return nil, fmt.Errorf("error decoding CoinGecko API response: %v", err)
	}

	// Gelen veriyi incelemek için print edelim
	fmt.Println(coinsData)

	coins := make([]*models.Coin, len(coinsData))
	for i, coinData := range coinsData {
		coins[i] = &models.Coin{
			ID:                   coinData.ID,
			Name:                 coinData.Name,
			Symbol:               coinData.Symbol,
			Price:                coinData.Price,
			OneHourChange:        coinData.OneHourChange,
			TwentyFourHourChange: coinData.TwentyFourHourChange,
			SevenDayChange:       coinData.SevenDayChange,
			MarketCap:            coinData.MarketCap,
			Volume24h:            coinData.Volume24h,
		}
	}

	return coins, nil
}

func (s *CoinService) GetCoinByName(name string) (*models.Coin, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", name)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data from CoinGecko API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	var result struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Symbol     string `json:"symbol"`
		MarketData struct {
			CurrentPrice struct {
				USD float64 `json:"usd"`
			} `json:"current_price"`
		} `json:"market_data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding CoinGecko API response: %v", err)
	}

	return &models.Coin{
		ID:   result.ID,
		Name: result.Name,
	}, nil
}
