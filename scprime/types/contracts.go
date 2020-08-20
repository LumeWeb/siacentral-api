package types

import (
	"time"

	scprimetypes "gitlab.com/scpcorp/ScPrime/types"
)

type (
	//StorageContract a storage contract on the blockchain
	StorageContract struct {
		ID                     string                `json:"id"`
		BlockID                string                `json:"block_id"`
		TransactionID          string                `json:"transaction_id"`
		MerkleRoot             string                `json:"merkle_root"`
		UnlockHash             string                `json:"unlock_hash"`
		Status                 string                `json:"status"`
		RevisionNumber         uint64                `json:"revision_number"`
		NegotiationHeight      uint64                `json:"negotiation_height"`
		ExpirationHeight       uint64                `json:"expiration_height"`
		ProofDeadline          uint64                `json:"proof_deadline"`
		ProofHeight            uint64                `json:"proof_height"`
		Payout                 scprimetypes.Currency `json:"payout"`
		FileSize               scprimetypes.Currency `json:"file_size"`
		ValidProofOutputs      []SiacoinOutput       `json:"valid_proof_outputs"`
		MissedProofOutputs     []SiacoinOutput       `json:"missed_proof_outputs"`
		NegotiationTimestamp   time.Time             `json:"negotiation_timestamp"`
		ExpirationTimestamp    time.Time             `json:"expiration_timestamp"`
		ProofDeadlineTimestamp time.Time             `json:"proof_deadline_timestamp"`
		ProofTimestamp         time.Time             `json:"proof_timestamp"`
		ProofConfirmed         bool                  `json:"proof_confirmed"`
		Unused                 bool                  `json:"unused"`
	}

	//StorageProof a storage proof on the blockchain
	StorageProof struct {
		ContractID    string    `json:"contract_id"`
		TransactionID string    `json:"transaction_id"`
		BlockID       string    `json:"block_id"`
		BlockHeight   uint64    `json:"block_height"`
		Segment       [64]byte  `json:"segment"`
		Hashset       []string  `json:"hashset"`
		Timestamp     time.Time `json:"timestamp"`
	}
)
