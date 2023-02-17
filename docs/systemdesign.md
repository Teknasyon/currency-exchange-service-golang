# Teknasyon Currency Exchange Service

## Motivation

Many brands in Teknasyon are need to exchange currencies and they are using different services for this purpose. We want to create a service that can be used by all brands in Teknasyon.

## Business Rules

- Exchange rates should be updated every 10 minutes.
- Exchange rates should be stored in a database.
- The base currency should be EUR.
- Exchange rates should be fetched from an external API.
- The service should be able to convert currencies.
- The service should be able to convert currencies with a specific date.

## Technical Requirements

- Golang 1.19 or higher
- Sqlite3 for development
- MySQL for production
- App should be able to run in a docker container
- The service should be able to handle high traffic
- The service should be able to handle high concurrency
- The service should be able to handle high availability
- The service should be able to handle high reliability
- The service should be able to handle high scalability
- The service should have a good test coverage
- The service should have a good documentation
- The service should have a good logging system
- The service should have a good monitoring system
- The service should have API documentation (Swagger)
- The service should have a good CI/CD pipeline
- The service should have a good error handling system
- The service should have a good rate limiting system
- The service should have a good caching system

## External API Endpoints

- `/latest` - Get the latest exchange rates
- `/exchange?from=USD&to=TRY&amount=100` - Convert currencies
- `/exchange?from=USD&to=TRY&amount=100&date=2021-01-01` - Convert currencies with a specific date
