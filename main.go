package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Teknasyon/currency-exchange-service-golang/currencyexchange"
	"github.com/Teknasyon/currency-exchange-service-golang/handlers"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {

	r := mux.NewRouter()

	currencyExchangeHandler := handlers.CurrencyExchange{
		CurrencyExchange: currencyexchange.New(),
	}

	r.HandleFunc("/latest", currencyExchangeHandler.Latest).Methods("GET")
	r.HandleFunc("/exchange", currencyExchangeHandler.Exchange).Methods("GET")

	log.Println("Listening on :8080...")
	if err := http.ListenAndServe(":"+getEnv("PORT", "8080"), r); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
