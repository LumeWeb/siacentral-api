package types

import (
	siatypes "gitlab.com/NebulousLabs/Sia/types"
)

type (
	//SiacoinInput an input of siacoins for a transaction
	SiacoinInput struct {
		SiacoinOutput
		UnlockConditions UnlockCondition `json:"unlock_conditions"`
	}

	//SiacoinOutput an output of siacoins for a transaction
	SiacoinOutput struct {
		OutputID       string            `json:"output_id"`
		UnlockHash     string            `json:"unlock_hash"`
		Source         string            `json:"source"`
		MaturityHeight uint64            `json:"maturity_height"`
		BlockHeight    uint64            `json:"block_height"`
		Value          siatypes.Currency `json:"value"`
	}

	//SiafundOutput an output of siafunds for a transaction
	SiafundOutput struct {
		OutputID           string            `json:"output_id"`
		BlockID            string            `json:"block_id"`
		SpentTransactionID string            `json:"spent_transaction_id"`
		UnlockHash         string            `json:"unlock_hash"`
		BlockHeight        uint64            `json:"block_height"`
		Value              siatypes.Currency `json:"value"`
		ClaimStart         siatypes.Currency `json:"claim_start"`
		ClaimValue         siatypes.Currency `json:"claim_value"`
	}

	//SiafundInput an input of siafunds for a transaction
	SiafundInput struct {
		SiafundOutput
		ClaimUnlockHash  string          `json:"claim_unlock_hash"`
		UnlockConditions UnlockCondition `json:"unlock_conditions"`
	}
)
