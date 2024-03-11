package sia

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/shopspring/decimal"
)

type (
	getPriceResp struct {
		APIResponse
		Siacoin map[string]decimal.Decimal `json:"siacoin"`
		Siafund map[string]decimal.Decimal `json:"siafund"`
	}

	getHistoricalPriceResp struct {
		APIResponse
		Rates     map[string]map[string]decimal.Decimal `json:"rates"`
		Timestamp time.Time                             `json:"timestamp"`
	}

	ExchangeRate struct {
		Currency  string                     `json:"currency"`
		Rates     map[string]decimal.Decimal `json:"rates"`
		Timestamp time.Time                  `json:"timestamp"`
	}

	getYearHistoricalPriceResp struct {
		APIResponse
		Rates []ExchangeRate `json:"rates"`
	}
)

// GetExchangeRate gets the current market exchange rate for Siacoin and Siafund
func (a *APIClient) GetExchangeRate() (siacoin map[string]decimal.Decimal, siafund map[string]decimal.Decimal, err error) {
	var resp getPriceResp

	code, err := a.makeAPIRequest(http.MethodGet, "/market/exchange-rate", nil, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	siacoin = resp.Siacoin
	siafund = resp.Siafund

	return
}

// GetHistoricalExchangeRate gets the historical market exchange rate for
// Siacoins at the specified timestamp
func (a *APIClient) GetHistoricalExchangeRate(timestamp time.Time) (map[string]decimal.Decimal, error) {
	var resp getHistoricalPriceResp

	v := url.Values{
		"timestamp": []string{timestamp.Format(time.RFC3339)},
	}

	code, err := a.makeAPIRequest(http.MethodGet, "/market/exchange-rate/historical?"+v.Encode(), nil, &resp)
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		return nil, errors.New(resp.Message)
	}

	rates := make(map[string]decimal.Decimal)
	for k, v := range resp.Rates["sc"] {
		rates[k] = v
	}

	return rates, nil
}

// GetYearExchangeRate gets the rates for a full calendar year
func (a *APIClient) GetYearExchangeRate(timestamp time.Time) ([]ExchangeRate, error) {
	var resp getYearHistoricalPriceResp

	y, _, _ := timestamp.Date()
	timestamp = time.Date(y, 1, 1, 0, 0, 0, 0, timestamp.Location())

	v := url.Values{
		"timestamp": []string{timestamp.Format(time.RFC3339)},
	}

	code, err := a.makeAPIRequest(http.MethodGet, "/market/exchange-rate/historical/year?"+v.Encode(), nil, &resp)
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		return nil, errors.New(resp.Message)
	}

	return resp.Rates, nil
}
