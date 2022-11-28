package sia

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	siamods "go.sia.tech/siad/modules"
	"go.sia.tech/siad/types"
)

type (
	getHostsResp struct {
		APIResponse
		Hosts []HostDetails `json:"hosts"`
	}

	getHostDetailResp struct {
		APIResponse
		Host HostDetails `json:"host"`
	}

	getAveragesResp struct {
		APIResponse
		Settings       HostConfig            `json:"settings"`
		PriceTable     siamods.RPCPriceTable `json:"price_table"`
		Benchmarks     AvgHostBenchmark      `json:"benchmarks"`
		BenchmarksRHP2 AvgHostBenchmark      `json:"benchmarks_rhp2"`
	}

	HostFilter url.Values
)

func (hf HostFilter) WithAcceptingContracts(accepting bool) {
	hf["acceptcontracts"] = []string{strconv.FormatBool(accepting)}
}

func (hf HostFilter) WithOnline(online bool) {
	hf["online"] = []string{strconv.FormatBool(online)}
}

func (hf HostFilter) WithBenchmarked(benchmarked bool) {
	hf["benchmarked"] = []string{strconv.FormatBool(benchmarked)}
}

func (hf HostFilter) WithMinUptime(minUptime float64) {
	hf["minuptime"] = []string{strconv.FormatFloat(minUptime, 'f', -1, 64)}
}

func (hf HostFilter) WithMinDuration(minDuration uint64) {
	hf["minduration"] = []string{strconv.FormatUint(minDuration, 10)}
}

func (hf HostFilter) WithMinStorage(minStorage uint64) {
	hf["minstorage"] = []string{strconv.FormatUint(minStorage, 10)}
}

func (hf HostFilter) WithMinUploadSpeed(minUploadSpeed uint64) {
	hf["minuploadspeed"] = []string{strconv.FormatUint(minUploadSpeed, 10)}
}

func (hf HostFilter) WithMinDownloadSpeed(minDownloadSpeed uint64) {
	hf["mindownloadspeed"] = []string{strconv.FormatUint(minDownloadSpeed, 10)}
}

func (hf HostFilter) WithMaxStoragePrice(price types.Currency) {
	hf["maxstorageprice"] = []string{price.String()}
}

func (hf HostFilter) WithMaxUploadPrice(price types.Currency) {
	hf["maxuploadprice"] = []string{price.String()}
}

func (hf HostFilter) WithMaxDownloadPrice(price types.Currency) {
	hf["maxdownloadprice"] = []string{price.String()}
}

func (hf HostFilter) WithMaxContractPrice(price types.Currency) {
	hf["maxcontractprice"] = []string{price.String()}
}

func (hf HostFilter) WithMaxBaseRPCPrice(price types.Currency) {
	hf["maxbaserpcprice"] = []string{price.String()}
}

func (hf HostFilter) WithSectorAccessPrice(price types.Currency) {
	hf["maxsectoraccessprice"] = []string{price.String()}
}

// GetNetworkAverages gets the average settings and benchmarks of all active hosts on the network
func (a *APIClient) GetNetworkAverages() (settings HostConfig, rhp3Bench AvgHostBenchmark, rhp2Bench AvgHostBenchmark, err error) {
	var resp getAveragesResp

	code, err := a.makeAPIRequest(http.MethodGet, "/hosts/network/averages", nil, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	settings = resp.Settings
	rhp3Bench = resp.Benchmarks
	rhp2Bench = resp.BenchmarksRHP2

	return
}

// GetActiveHosts gets all Sia hosts that have been successfully scanned in the last 24 hours
func (a *APIClient) GetActiveHosts(filter HostFilter, page, limit int) (hosts []HostDetails, err error) {
	var resp getHostsResp

	if page < 0 {
		page = 0
	}

	if limit < 0 || limit > 500 {
		limit = 500
	}

	url.Values(filter).Add("page", strconv.Itoa(page))
	url.Values(filter).Add("limit", strconv.Itoa(limit))

	endpoint, _ := url.Parse("https://api.siacentral.com/v2/hosts")
	endpoint.RawQuery = url.Values(filter).Encode()

	code, err := a.makeAPIRequest(http.MethodGet, endpoint.String(), nil, &resp)

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

// GetHost finds a host matching the public key or netaddress
func (a *APIClient) GetHost(id string) (host HostDetails, err error) {
	var resp getHostDetailResp

	code, err := a.makeAPIRequest(http.MethodGet, fmt.Sprintf("/hosts/%s", url.PathEscape(id)), nil, &resp)

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
