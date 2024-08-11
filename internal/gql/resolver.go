package gql

import (
	"context"
	"fmt"
	"go-crypto-tracker/internal/service"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Coins(ctx context.Context) ([]*Coin, error) {
	coinService := service.CoinService{}
	coins, err := coinService.GetCoins()
	if err != nil {
		return nil, err
	}

	result := make([]*Coin, len(coins))
	for i, coin := range coins {
		result[i] = &Coin{
			ID:                   coin.ID,
			Name:                 coin.Name,
			Symbol:               coin.Symbol,
			Price:                coin.Price,
			OneHourChange:        &coin.OneHourChange,        // Pointer olarak geçiriliyor
			TwentyFourHourChange: &coin.TwentyFourHourChange, // Pointer olarak geçiriliyor
			SevenDayChange:       &coin.SevenDayChange,       // Pointer olarak geçiriliyor
			MarketCap:            &coin.MarketCap,            // Pointer olarak geçiriliyor
			Volume24h:            &coin.Volume24h,            // Pointer olarak geçiriliyor
		}
	}
	return result, nil
}

func (r *queryResolver) Coin(ctx context.Context, name string) (*Coin, error) {
	coinService := service.CoinService{}
	coin, err := coinService.GetCoinByName(name)
	if err != nil {
		return nil, fmt.Errorf("error fetching coin data: %v", err)
	}

	// *models.Coin yerine *Coin döndürüyoruz
	return &Coin{
		ID:                   coin.ID,
		Name:                 coin.Name,
		Symbol:               coin.Symbol,
		Price:                coin.Price,
		OneHourChange:        &coin.OneHourChange,
		TwentyFourHourChange: &coin.TwentyFourHourChange,
		SevenDayChange:       &coin.SevenDayChange,
		MarketCap:            &coin.MarketCap,
		Volume24h:            &coin.Volume24h,
	}, nil
}
