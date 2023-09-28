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
		Siacoin map[string]float64 `json:"siacoin"`
		Siafund map[string]float64 `json:"siafund"`
	}

	getHistoricalPriceResp struct {
		APIResponse
		Rates     map[string]map[string]decimal.Decimal `json:"rates"`
		Timestamp time.Time                             `json:"timestamp"`
	}
)

// GetExchangeRate gets the current market exchange rate for Siacoin and Siafund
func (a *APIClient) GetExchangeRate() (siacoin map[string]float64, siafund map[string]float64, err error) {
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
func (a *APIClient) GetHistoricalExchangeRate(timestamp time.Time) (map[string]float64, error) {
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

	rates := make(map[string]float64)
	for k, v := range resp.Rates["sc"] {
		rates[k], _ = v.Float64()
	}

	return rates, nil
}
