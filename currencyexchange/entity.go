package currencyexchange

type CurrencyExchange struct {
	From string `json:"from"`
	To   string `json:"to"`
	Rate string `json:"rate"`
}

type Rate struct {
	CurrencyCode string `json:"currency_code"`
	Rate         string `json:"rate"`
}
