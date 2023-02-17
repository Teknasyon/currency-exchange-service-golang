package currencyexchange

import (
	"context"
	"log"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Latest(ctx context.Context) ([]Rate, error) {
	log.Println("Latest called")

	return []Rate{}, nil
}

type CurrencyExchangeParams struct {
	From   string
	To     string
	Amount string
	Date   string
}

func (s *Service) Exchange(ctx context.Context, params *CurrencyExchangeParams) (CurrencyExchange, error) {
	return CurrencyExchange{}, nil
}
