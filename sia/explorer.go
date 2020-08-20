package sia

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/siacentral/apisdkgo/sia/types"
)

type (
	getBlockResp struct {
		APIResponse
		Block types.Block `json:"block"`
	}

	batchBlocksResp struct {
		APIResponse
		Blocks []types.Block `json:"blocks"`
	}

	getTransactionResp struct {
		APIResponse
		Transaction types.Transaction `json:"transaction"`
	}

	batchTransactionsResp struct {
		APIResponse
		Transactions []types.Transaction `json:"transactions"`
	}

	getContractResp struct {
		APIResponse
		Contract types.StorageContract `json:"contract"`
	}

	batchContractsResp struct {
		APIResponse
		Contracts []types.StorageContract `json:"contracts"`
	}
)

//GetLatestBlock returns the latest block in the Sia Central explorer
func (a *APIClient) GetLatestBlock() (block types.Block, err error) {
	var resp getBlockResp

	code, err := a.makeAPIRequest(http.MethodGet, "/explorer/blocks", nil, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	block = resp.Block
	return
}

//GetBlockByID returns the block with the matching id in the Sia Central explorer
func (a *APIClient) GetBlockByID(id string) (block types.Block, err error) {
	var resp getBlockResp

	code, err := a.makeAPIRequest(http.MethodGet, fmt.Sprintf("/explorer/blocks/%s", id), nil, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	block = resp.Block
	return
}

//GetBlockByHeight returns the block at the specified height in the Sia Central explorer
func (a *APIClient) GetBlockByHeight(height uint64) (block types.Block, err error) {
	var resp getBlockResp

	code, err := a.makeAPIRequest(http.MethodGet, fmt.Sprintf("/explorer/blocks/%d", height), nil, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	block = resp.Block
	return
}

//FindBlocksByID returns all blocks with the specified ids from the Sia Central explorer
func (a *APIClient) FindBlocksByID(ids ...string) (blocks []types.Block, err error) {
	var resp batchBlocksResp

	if len(ids) > 10000 {
		err = errors.New("maximum of 10000 ids")
		return
	}

	code, err := a.makeAPIRequest(http.MethodPost, "/explorer/blocks", map[string]interface{}{
		"block_ids": ids,
	}, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	blocks = resp.Blocks
	return
}

//FindBlocksByHeight returns all blocks with the specified heights from the Sia Central explorer
func (a *APIClient) FindBlocksByHeight(heights ...uint64) (blocks []types.Block, err error) {
	var resp batchBlocksResp

	if len(heights) > 10000 {
		err = errors.New("maximum of 10000 heights")
		return
	}

	code, err := a.makeAPIRequest(http.MethodPost, "/explorer/blocks", map[string]interface{}{
		"heights": heights,
	}, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	blocks = resp.Blocks
	return
}

//GetTransactionByID returns the transaction at the specified height in the Sia Central explorer
func (a *APIClient) GetTransactionByID(id string) (transaction types.Transaction, err error) {
	var resp getTransactionResp

	code, err := a.makeAPIRequest(http.MethodGet, fmt.Sprintf("/explorer/transactions/%s", id), nil, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	transaction = resp.Transaction
	return
}

//FindTransactionsByID returns all transactions with the specified ids from the Sia Central explorer
func (a *APIClient) FindTransactionsByID(ids ...string) (transactions []types.Transaction, err error) {
	var resp batchTransactionsResp

	if len(ids) > 10000 {
		err = errors.New("maximum of 10000 ids")
		return
	}

	code, err := a.makeAPIRequest(http.MethodPost, "/explorer/transactions", map[string]interface{}{
		"transaction_ids": ids,
	}, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	transactions = resp.Transactions
	return
}

//GetContractByID returns the contract at the specified height in the Sia Central explorer
func (a *APIClient) GetContractByID(id string) (contract types.StorageContract, err error) {
	var resp getContractResp

	code, err := a.makeAPIRequest(http.MethodGet, fmt.Sprintf("/explorer/contracts/%s", id), nil, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	contract = resp.Contract
	return
}

//FindContractsByID returns all contracts with the specified ids from the Sia Central explorer
func (a *APIClient) FindContractsByID(ids ...string) (contracts []types.StorageContract, err error) {
	var resp batchContractsResp

	if len(ids) > 10000 {
		err = errors.New("maximum of 10000 ids")
		return
	}

	code, err := a.makeAPIRequest(http.MethodPost, "/explorer/contracts", map[string]interface{}{
		"contracts": ids,
	}, &resp)

	if err != nil {
		return
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		err = errors.New(resp.Message)
		return
	}

	contracts = resp.Contracts
	return
}
