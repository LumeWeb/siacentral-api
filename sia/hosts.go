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
	HostSort   string
)

var (
	HostSortDateCreated        HostSort = "date_created"
	HostSortNetAddress         HostSort = "net_address"
	HostSortPublicKey          HostSort = "public_key"
	HostSortAcceptingContracts HostSort = "accepting_contracts"
	HostSortUptime             HostSort = "uptime"
	HostSortUploadSpeed        HostSort = "upload_speed"
	HostSortDownloadSpeed      HostSort = "download_speed"
	HostSortRemainingStorage   HostSort = "remaining_storage"
	HostSortTotalStorage       HostSort = "total_storage"
	HostSortUsedStorage        HostSort = "used_storage"
	HostSortAge                HostSort = "age"
	HostSortUtilization        HostSort = "utilization"
	HostSortContractPrice      HostSort = "contract_price"
	HostSortStoragePrice       HostSort = "storage_price"
	HostSortDownloadPrice      HostSort = "download_price"
	HostSortUploadPrice        HostSort = "upload_price"
)

// WithAcceptingContracts sets the accepting contracts parameter for the host's query
func (hf HostFilter) WithAcceptingContracts(accepting bool) {
	hf["acceptcontracts"] = []string{strconv.FormatBool(accepting)}
}

// WithOnline sets the online parameter for the host's query
func (hf HostFilter) WithOnline(online bool) {
	hf["online"] = []string{strconv.FormatBool(online)}
}

// WithLimit sets the benchmark parameter for the host's query
func (hf HostFilter) WithBenchmarked(benchmarked bool) {
	hf["benchmarked"] = []string{strconv.FormatBool(benchmarked)}
}

// WithMinAge sets the min age for the host query
func (hf HostFilter) WithMinAge(age uint64) {
	hf["minage"] = []string{strconv.FormatUint(age, 10)}
}

// WithMinUptime sets the min uptime for the host query
func (hf HostFilter) WithMinUptime(minUptime float64) {
	hf["minuptime"] = []string{strconv.FormatFloat(minUptime, 'f', -1, 64)}
}

// WithMinDuration sets the min contract duration for the host query
func (hf HostFilter) WithMinDuration(minDuration uint64) {
	hf["minduration"] = []string{strconv.FormatUint(minDuration, 10)}
}

// WithMinStorage sets the min available storage for the host query
func (hf HostFilter) WithMinStorage(minStorage uint64) {
	hf["minstorage"] = []string{strconv.FormatUint(minStorage, 10)}
}

// WithMinUploadSpeed sets the min upload speed for the host query
func (hf HostFilter) WithMinUploadSpeed(minUploadSpeed uint64) {
	hf["minuploadspeed"] = []string{strconv.FormatUint(minUploadSpeed, 10)}
}

// WithMinDownloadSpeed sets the min download speed for the host query
func (hf HostFilter) WithMinDownloadSpeed(minDownloadSpeed uint64) {
	hf["mindownloadspeed"] = []string{strconv.FormatUint(minDownloadSpeed, 10)}
}

// WithMaxStoragePrice sets the max storage price for the host query
func (hf HostFilter) WithMaxStoragePrice(price types.Currency) {
	hf["maxstorageprice"] = []string{price.String()}
}

// WithMaxUploadPrice sets the max upload price for the host query
func (hf HostFilter) WithMaxUploadPrice(price types.Currency) {
	hf["maxuploadprice"] = []string{price.String()}
}

// WithMaxDownloadPrice sets the max download price for the host query
func (hf HostFilter) WithMaxDownloadPrice(price types.Currency) {
	hf["maxdownloadprice"] = []string{price.String()}
}

// WithMaxContractPrice sets the max contract price for the host query
func (hf HostFilter) WithMaxContractPrice(price types.Currency) {
	hf["maxcontractprice"] = []string{price.String()}
}

// WithMaxBaseRPCPrice sets the max base RPC price for the host query
func (hf HostFilter) WithMaxBaseRPCPrice(price types.Currency) {
	hf["maxbaserpcprice"] = []string{price.String()}
}

// WithSectorAccessPrice sets the sector access price for the host query
func (hf HostFilter) WithSectorAccessPrice(price types.Currency) {
	hf["maxsectoraccessprice"] = []string{price.String()}
}

// WithSort sets the sort order for the host's query
func (hf HostFilter) WithSort(field HostSort, desc bool) {
	hf["sort"] = []string{string(field)}
	if desc {
		hf["desc"] = []string{"desc"}
	} else {
		hf["desc"] = []string{"asc"}
	}
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

	values := make(url.Values)
	for k, v := range filter {
		values[k] = v
	}

	values.Add("page", strconv.Itoa(page))
	values.Add("limit", strconv.Itoa(limit))

	endpoint, _ := url.Parse("https://api.siacentral.com/v2/hosts")
	endpoint.RawQuery = values.Encode()

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
