package sia

import (
	"errors"
	"net/http"
)

type getPriceResp struct {
	APIResponse
	Siacoin map[string]float64 `json:"siacoin"`
	Siafund map[string]float64 `json:"siafund"`
}

//GetExchangeRate gets the current market exchange rate for Siacoin and Siafund
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
