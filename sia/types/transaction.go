package types

import (
	"time"

	siatypes "go.sia.tech/siad/types"
)

type (
	//Transaction a transaction on the blockchain
	Transaction struct {
		ID                    string                 `json:"id"`
		BlockID               string                 `json:"block_id"`
		BlockHeight           uint64                 `json:"block_height,omitempty"`
		Confirmations         uint64                 `json:"confirmations"`
		BlockIndex            int                    `json:"-"`
		Timestamp             time.Time              `json:"timestamp"`
		Fees                  siatypes.Currency      `json:"fees"`
		SiacoinInputs         []SiacoinInput         `json:"siacoin_inputs"`
		SiacoinOutputs        []SiacoinOutput        `json:"siacoin_outputs"`
		SiafundInputs         []SiafundInput         `json:"siafund_inputs"`
		SiafundOutputs        []SiafundOutput        `json:"siafund_outputs"`
		StorageContracts      []StorageContract      `json:"storage_contracts"`
		ContractRevisions     []StorageContract      `json:"contract_revisions"`
		StorageProofs         []StorageProof         `json:"storage_proofs"`
		MinerFees             []siatypes.Currency    `json:"miner_fees"`
		HostAnnouncements     []Announcement         `json:"host_announcements"`
		ArbitraryData         [][]byte               `json:"arbitrary_data"`
		TransactionSignatures []TransactionSignature `json:"transaction_signatures"`
	}

	//UnlockCondition unlock conditions of a transaction input
	UnlockCondition struct {
		PublicKeys         []string `json:"public_keys"`
		Timelock           uint64   `json:"timelock"`
		RequiredSignatures uint64   `json:"required_signatures"`
	}

	//CoveredFields the covered fields of a transaction signature and their indexes
	CoveredFields struct {
		WholeTransaction         bool     `json:"whole_transaction"`
		SiacoinInputs            []uint64 `json:"siacoin_inputs"`
		SiacoinOutputs           []uint64 `json:"siacoin_outputs"`
		StorageContracts         []uint64 `json:"storage_contracts"`
		StorageContractRevisions []uint64 `json:"storage_contract_revisions"`
		StorageProofs            []uint64 `json:"storage_proofs"`
		MinerFees                []uint64 `json:"miner_fees"`
		ArbitraryData            []uint64 `json:"arbitrary_data"`
		TransactionSignatures    []uint64 `json:"transaction_signatures"`
	}

	//TransactionSignature a signature verifying a part of the transaction
	TransactionSignature struct {
		ParentID       string        `json:"parent_id"`
		TransactionID  string        `json:"transaction_id"`
		BlockID        string        `json:"block_id"`
		Signature      string        `json:"signature"`
		PublicKeyIndex uint64        `json:"public_key_index"`
		CoveredFields  CoveredFields `json:"covered_fields"`
	}
)
