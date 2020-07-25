package apisdkgo

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/siacentral/apisdkgo/types"
)

type (
	getHostsResp struct {
		APIResponse
		Hosts []types.HostInfo `json:"hosts"`
	}

	getHostDetailResp struct {
		APIResponse
		Host types.HostDetails `json:"host"`
	}

	getNetworkSettingsResp struct {
		APIResponse
		Settings types.HostConfig `json:"settings"`
	}
)

//GetAverageSettings gets the average settings of all active hosts on the network
func (a *APIClient) GetAverageSettings() (settings types.HostConfig, err error) {
	var resp getNetworkSettingsResp

	code, err := a.makeAPIRequest(HTTPGet, "/hosts/settings/average", nil, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	settings = resp.Settings

	return
}

//GetActiveHosts gets all Sia hosts that have been successfully scanned in the last 24 hours
func (a *APIClient) GetActiveHosts() (hosts []types.HostInfo, err error) {
	var resp getHostsResp

	code, err := a.makeAPIRequest(HTTPGet, "/hosts", nil, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	hosts = resp.Hosts

	return
}

//GetHost finds a host matching the public key or netaddress
func (a *APIClient) GetHost(id string) (host types.HostDetails, err error) {
	var resp getHostDetailResp

	code, err := a.makeAPIRequest(HTTPGet, fmt.Sprintf("/hosts/%s", url.PathEscape(id)), nil, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	host = resp.Host

	return
}
