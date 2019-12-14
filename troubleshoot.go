package apiclient

import (
	"errors"

	"github.com/siacentral/apiclient/types"
)

type (
	getConnectionResp struct {
		APIResponse
		Report types.ConnectionReport `json:"report"`
	}
)

//GetHostConnectivity checks that a host is running and connectable at the provided netaddress
func GetHostConnectivity(netaddress string) (report types.ConnectionReport, err error) {
	var resp getConnectionResp

	code, err := makeAPIRequest(HTTPGet, "/market/exchange-rate", nil, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 {
		err = errors.New(resp.Message)
		return
	}

	report = resp.Report

	return
}
