package sia

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/siacentral/apisdkgo/sia/types"
	siamods "go.sia.tech/siad/modules"
	siatypes "go.sia.tech/siad/types"
)

type (
	getHostsResp struct {
		APIResponse
		Hosts []types.HostDetails `json:"hosts"`
	}

	getHostDetailResp struct {
		APIResponse
		Host types.HostDetails `json:"host"`
	}

	getAveragesResp struct {
		APIResponse
		Settings       types.HostConfig       `json:"settings"`
		PriceTable     siamods.RPCPriceTable  `json:"price_table"`
		Benchmarks     types.AvgHostBenchmark `json:"benchmarks"`
		BenchmarksRHP2 types.AvgHostBenchmark `json:"benchmarks_rhp2"`
	}

	HostFilter struct {
		Page                 int
		Limit                int
		AcceptingContracts   *bool
		Online               *bool
		Benchmarked          *bool
		MinUptime            *float32
		MinDuration          *uint64
		MinStorage           *uint64
		MinUploadSpeed       *uint64
		MinDownloadSpeed     *uint64
		MaxStoragePrice      *siatypes.Currency
		MaxUploadPrice       *siatypes.Currency
		MaxDownloadPrice     *siatypes.Currency
		MaxContractPrice     *siatypes.Currency
		MaxBaseRPCPrice      *siatypes.Currency
		MaxSectorAccessPrice *siatypes.Currency
	}
)

func buildFilter(filter HostFilter) url.Values {
	vals := url.Values{}

	if filter.Page > 0 {
		vals.Add("page", strconv.Itoa(filter.Page))
	}

	if filter.Limit > 0 {
		vals.Add("limit", strconv.Itoa(filter.Limit))
	}

	if filter.AcceptingContracts != nil {
		vals.Add("acceptcontracts", fmt.Sprintf("%t", *filter.AcceptingContracts))
	}

	if filter.Online != nil {
		vals.Add("online", fmt.Sprintf("%t", *filter.Online))
	}

	if filter.Benchmarked != nil {
		vals.Add("benchmarked", fmt.Sprintf("%t", *filter.Benchmarked))
	}

	if filter.MinUptime != nil {
		vals.Add("minuptime", fmt.Sprintf("%f", *filter.MinUptime))
	}

	if filter.MinDuration != nil {
		vals.Add("minduration", strconv.FormatUint(*filter.MinDuration, 10))
	}

	if filter.MinStorage != nil {
		vals.Add("minstorage", strconv.FormatUint(*filter.MinStorage, 10))
	}

	if filter.MinUploadSpeed != nil {
		vals.Add("minuploadspeed", strconv.FormatUint(*filter.MinUploadSpeed, 10))
	}

	if filter.MinDownloadSpeed != nil {
		vals.Add("mindownloadspeed", strconv.FormatUint(*filter.MinDownloadSpeed, 10))
	}

	if filter.MaxStoragePrice != nil {
		vals.Add("maxstorageprice", filter.MaxStoragePrice.String())
	}

	if filter.MaxUploadPrice != nil {
		vals.Add("maxuploadprice", filter.MaxUploadPrice.String())
	}

	if filter.MaxDownloadPrice != nil {
		vals.Add("maxdownloadprice", filter.MaxDownloadPrice.String())
	}

	if filter.MaxContractPrice != nil {
		vals.Add("maxcontractprice", filter.MaxContractPrice.String())
	}

	if filter.MaxBaseRPCPrice != nil {
		vals.Add("maxbaserpcprice", filter.MaxBaseRPCPrice.String())
	}

	if filter.MaxSectorAccessPrice != nil {
		vals.Add("maxsectoraccessprice", filter.MaxSectorAccessPrice.String())
	}

	return vals
}

//GetNetworkAverages gets the average settings and benchmarks of all active hosts on the network
func (a *APIClient) GetNetworkAverages() (settings types.HostConfig, rhp3Bench types.AvgHostBenchmark, rhp2Bench types.AvgHostBenchmark, err error) {
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

//GetActiveHosts gets all Sia hosts that have been successfully scanned in the last 24 hours
func (a *APIClient) GetActiveHosts(filter HostFilter) (hosts []types.HostDetails, err error) {
	var resp getHostsResp

	url, _ := url.Parse("https://api.siacentral.com/v2/hosts")
	url.RawQuery = buildFilter(filter).Encode()

	code, err := a.makeAPIRequest(http.MethodGet, url.String(), nil, &resp)

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
