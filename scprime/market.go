package scprime

import (
	"errors"
	"net/http"
)

type getPriceResp struct {
	APIResponse
	SCPrimeCoin map[string]float64 `json:"scprimecoin"`
	SCPrimeFund map[string]float64 `json:"scprimefund"`
}

//GetExchangeRate gets the current market exchange rate for SCPrimeCoin and SCPrimeFunds
func (a *APIClient) GetExchangeRate() (scp map[string]float64, spf map[string]float64, err error) {
	var resp getPriceResp

	code, err := a.makeAPIRequest(http.MethodGet, "/market/exchange-rate", nil, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	scp = resp.SCPrimeCoin
	spf = resp.SCPrimeFund

	return
}
