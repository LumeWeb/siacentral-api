package types

import (
	"time"

	siatypes "gitlab.com/NebulousLabs/Sia/types"
)

type (
	//AddressUsage AddressUsage
	AddressUsage struct {
		Address   string `json:"address"`
		UsageType string `json:"usage_type"`
	}

	//WalletTransaction SiacoinTransaction
	WalletTransaction struct {
		TransactionID  string            `json:"transaction_id"`
		BlockHeight    uint64            `json:"block_height"`
		Confirmations  uint64            `json:"confirmations"`
		Fees           siatypes.Currency `json:"fees"`
		SiacoinOutputs []SiacoinOutput   `json:"siacoin_outputs"`
		SiacoinInputs  []SiacoinInput    `json:"siacoin_inputs"`
		SiafundInputs  []SiafundInput    `json:"siafund_inputs"`
		SiafundOutputs []SiafundOutput   `json:"siafund_outputs"`
		Tags           []string          `json:"tags"`
		Timestamp      time.Time         `json:"timestamp"`
	}
)
