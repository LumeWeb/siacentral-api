package types

import (
	scprimetypes "gitlab.com/scpcorp/ScPrime/types"
)

type (
	//SiacoinInput an input of siacoins for a transaction
	SiacoinInput struct {
		SiacoinOutput
		UnlockConditions UnlockCondition `json:"unlock_conditions"`
	}

	//SiacoinOutput an output of siacoins for a transaction
	SiacoinOutput struct {
		OutputID       string                `json:"output_id"`
		UnlockHash     string                `json:"unlock_hash"`
		Source         string                `json:"source"`
		MaturityHeight uint64                `json:"maturity_height"`
		BlockHeight    uint64                `json:"block_height"`
		Value          scprimetypes.Currency `json:"value"`
	}

	//SiafundOutput an output of siafunds for a transaction
	SiafundOutput struct {
		OutputID     string                `json:"output_id"`
		UnlockHash   string                `json:"unlock_hash"`
		BlockHeight  uint64                `json:"block_height"`
		Value        scprimetypes.Currency `json:"value"`
		SiacoinClaim scprimetypes.Currency `json:"siacoin_claim"`
	}

	//SiafundInput an input of siafunds for a transaction
	SiafundInput struct {
		SiafundOutput
		ClaimUnlockHash  string          `json:"claim_unlock_hash"`
		UnlockConditions UnlockCondition `json:"unlock_conditions"`
	}
)
