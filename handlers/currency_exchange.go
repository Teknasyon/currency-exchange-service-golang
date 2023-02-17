package handlers

import (
	"context"
	"net/http"

	"github.com/Teknasyon/currency-exchange-service-golang/currencyexchange"
)

type CurrencyExchangeService interface {
	Latest(context.Context) ([]currencyexchange.Rate, error)
	Exchange(context.Context, *currencyexchange.CurrencyExchangeParams) (currencyexchange.CurrencyExchange, error)
}

type CurrencyExchange struct {
	CurrencyExchange CurrencyExchangeService
}

// Latest returns the latest exchange rate for a given currency
func (c *CurrencyExchange) Latest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	_, err := c.CurrencyExchange.Latest(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Latest"))
}

// Exchange returns the exchange rate for a given currency on a given date
// If no date is provided, it returns the latest exchange rate
// Example: /exchange?from=USD&to=EUR&date=2018-01-01
// Example: /exchange?from=USD&to=EUR
func (c *CurrencyExchange) Exchange(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Exchange"))
}
