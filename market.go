package main

import (
	"errors"
)

type getPriceResp struct {
	APIResponse
	Rates map[string]float64 `json:"price"`
}

//GetSiacoinExchangeRate gets the current market exchange rate for Siacoin
func GetSiacoinExchangeRate() (rates map[string]float64, err error) {
	var resp getPriceResp

	code, err := makeAPIRequest(HTTPGet, "/market/exchange-rate", nil, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	rates = resp.Rates

	return
}
