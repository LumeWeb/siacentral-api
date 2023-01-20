package sia

import (
	"time"

	"go.sia.tech/siad/modules"
	types "go.sia.tech/siad/types"
)

type (
	//Announcement a host announcement on the blockchain
	Announcement struct {
		TransactionID string    `json:"transaction_id"`
		BlockID       string    `json:"block_id"`
		PublicKey     string    `json:"public_key"`
		NetAddress    string    `json:"net_address"`
		Height        uint64    `json:"block_height"`
		Timestamp     time.Time `json:"timestamp,omitempty"`
	}

	// AvgHostBenchmark AvgHostBenchmark
	AvgHostBenchmark struct {
		ContractTime uint64 `json:"contract_time"`
		UploadTime   uint64 `json:"upload_time"`
		DownloadTime uint64 `json:"download_time"`
		DataSize     uint64 `json:"data_size"`
	}

	// HostBenchmark a benchmark from the host
	HostBenchmark struct {
		ContractTime   uint64    `json:"contract_time"`
		UploadTime     uint64    `json:"upload_time"`
		DownloadTime   uint64    `json:"download_time"`
		DataSize       uint64    `json:"data_size"`
		LastAttempt    time.Time `json:"last_attempt"`
		LastSuccessful time.Time `json:"last_successful"`
		ErrorMessage   *string   `json:"error"`
	}

	//HostDetails the latest details on the host
	HostDetails struct {
		NetAddress         string                 `json:"net_address"`
		PublicKey          string                 `json:"public_key"`
		Version            string                 `json:"version"`
		EstimatedUptime    float32                `json:"estimated_uptime"`
		Online             bool                   `json:"online"`
		FirstSeenHeight    uint64                 `json:"first_seen_height"`
		FirstSeenTimestamp time.Time              `json:"first_seen_timestamp"`
		LastScan           time.Time              `json:"last_scan"`
		LastSuccessScan    time.Time              `json:"last_success_scan"`
		Settings           *HostExternalSettings  `json:"settings"`
		PriceTable         *modules.RPCPriceTable `json:"price_table"`
		Benchmark          *HostBenchmark         `json:"benchmark"`
		BenchmarkRHP2      *HostBenchmark         `json:"benchmark_rhp2"`
	}

	//HostConfig the settings pulled from the host during a scan
	HostConfig struct {
		MaxDownloadBatchSize   uint64         `json:"max_download_batch_size"`
		MaxDuration            uint64         `json:"max_duration"`
		MaxReviseBatchSize     uint64         `json:"max_revise_batch_size"`
		RemainingStorage       uint64         `json:"remaining_storage"`
		SectorSize             uint64         `json:"sector_size"`
		TotalStorage           uint64         `json:"total_storage"`
		WindowSize             uint64         `json:"window_size"`
		RevisionNumber         uint64         `json:"revision_number"`
		BaseRPCPrice           types.Currency `json:"base_rpc_price"`
		Collateral             types.Currency `json:"collateral"`
		MaxCollateral          types.Currency `json:"max_collateral"`
		ContractPrice          types.Currency `json:"contract_price"`
		DownloadBandwidthPrice types.Currency `json:"download_price"`
		SectorAccessPrice      types.Currency `json:"sector_access_price"`
		StoragePrice           types.Currency `json:"storage_price"`
		UploadBandwidthPrice   types.Currency `json:"upload_price"`
	}

	//HostExternalSettings the settings pulled from the host during a scan
	HostExternalSettings struct {
		NetAddress             string         `json:"netaddress"`
		Version                string         `json:"version"`
		AcceptingContracts     bool           `json:"accepting_contracts"`
		MaxDownloadBatchSize   uint64         `json:"max_download_batch_size"`
		MaxDuration            uint64         `json:"max_duration"`
		MaxReviseBatchSize     uint64         `json:"max_revise_batch_size"`
		RemainingStorage       uint64         `json:"remaining_storage"`
		SectorSize             uint64         `json:"sector_size"`
		TotalStorage           uint64         `json:"total_storage"`
		WindowSize             uint64         `json:"window_size"`
		RevisionNumber         uint64         `json:"revision_number"`
		BaseRPCPrice           types.Currency `json:"base_rpc_price"`
		Collateral             types.Currency `json:"collateral"`
		MaxCollateral          types.Currency `json:"max_collateral"`
		ContractPrice          types.Currency `json:"contract_price"`
		DownloadBandwidthPrice types.Currency `json:"download_price"`
		SectorAccessPrice      types.Currency `json:"sector_access_price"`
		StoragePrice           types.Currency `json:"storage_price"`
		UploadBandwidthPrice   types.Currency `json:"upload_price"`
	}

	//SiacoinInput an input of siacoins for a transaction
	SiacoinInput struct {
		SiacoinOutput
		UnlockConditions UnlockCondition `json:"unlock_conditions"`
	}

	//SiacoinOutput an output of siacoins for a transaction
	SiacoinOutput struct {
		OutputID           string         `json:"output_id"`
		UnlockHash         string         `json:"unlock_hash"`
		Source             string         `json:"source"`
		SpentTransactionID string         `json:"spent_transaction_id"`
		MaturityHeight     uint64         `json:"maturity_height"`
		BlockHeight        uint64         `json:"block_height"`
		Value              types.Currency `json:"value"`
	}

	//SiafundOutput an output of siafunds for a transaction
	SiafundOutput struct {
		OutputID           string         `json:"output_id"`
		BlockID            string         `json:"block_id"`
		SpentTransactionID string         `json:"spent_transaction_id"`
		UnlockHash         string         `json:"unlock_hash"`
		BlockHeight        uint64         `json:"block_height"`
		Value              types.Currency `json:"value"`
		ClaimStart         types.Currency `json:"claim_start"`
		ClaimValue         types.Currency `json:"claim_value"`
	}

	//SiafundInput an input of siafunds for a transaction
	SiafundInput struct {
		SiafundOutput
		ClaimUnlockHash  string          `json:"claim_unlock_hash"`
		UnlockConditions UnlockCondition `json:"unlock_conditions"`
	}

	//StorageContract a storage contract on the blockchain
	StorageContract struct {
		ID                     string            `json:"id"`
		BlockID                string            `json:"block_id"`
		TransactionID          string            `json:"transaction_id"`
		MerkleRoot             string            `json:"merkle_root"`
		UnlockHash             string            `json:"unlock_hash"`
		Status                 string            `json:"status"`
		RevisionNumber         uint64            `json:"revision_number"`
		NegotiationHeight      uint64            `json:"negotiation_height"`
		ExpirationHeight       uint64            `json:"expiration_height"`
		ProofDeadline          uint64            `json:"proof_deadline"`
		ProofHeight            uint64            `json:"proof_height"`
		Payout                 types.Currency    `json:"payout"`
		FileSize               types.Currency    `json:"file_size"`
		ValidProofOutputs      []SiacoinOutput   `json:"valid_proof_outputs"`
		MissedProofOutputs     []SiacoinOutput   `json:"missed_proof_outputs"`
		NegotiationTimestamp   time.Time         `json:"negotiation_timestamp"`
		ExpirationTimestamp    time.Time         `json:"expiration_timestamp"`
		ProofDeadlineTimestamp time.Time         `json:"proof_deadline_timestamp"`
		ProofTimestamp         time.Time         `json:"proof_timestamp"`
		ProofConfirmed         bool              `json:"proof_confirmed"`
		Unused                 bool              `json:"unused"`
		PreviousRevisions      []StorageContract `json:"previous_revisions"`
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

	//Transaction a transaction on the blockchain
	Transaction struct {
		ID                    string                 `json:"id"`
		BlockID               string                 `json:"block_id"`
		BlockHeight           uint64                 `json:"block_height,omitempty"`
		Confirmations         uint64                 `json:"confirmations"`
		BlockIndex            int                    `json:"-"`
		Timestamp             time.Time              `json:"timestamp"`
		Fees                  types.Currency         `json:"fees"`
		SiacoinInputs         []SiacoinInput         `json:"siacoin_inputs"`
		SiacoinOutputs        []SiacoinOutput        `json:"siacoin_outputs"`
		SiafundInputs         []SiafundInput         `json:"siafund_inputs"`
		SiafundOutputs        []SiafundOutput        `json:"siafund_outputs"`
		StorageContracts      []StorageContract      `json:"storage_contracts"`
		ContractRevisions     []StorageContract      `json:"contract_revisions"`
		StorageProofs         []StorageProof         `json:"storage_proofs"`
		MinerFees             []types.Currency       `json:"miner_fees"`
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

	//Block a block on the Sia blockchain
	Block struct {
		ID                string          `json:"id"`
		ParentID          string          `json:"parent_id"`
		Height            uint64          `json:"height"`
		Nonce             [8]byte         `json:"nonce"`
		Transactions      []Transaction   `json:"transactions"`
		HostAnnouncements []Announcement  `json:"host_announcements"`
		SiacoinOutputs    []SiacoinOutput `json:"siacoin_outputs"`
		Timestamp         time.Time       `json:"timestamp"`
	}

	//ConnectionReport information about the connection
	ConnectionReport struct {
		NetAddress    string               `json:"netaddress"`
		PublicKey     string               `json:"public_key"`
		ConnectedIP   string               `json:"connected_ip"`
		Resolved      bool                 `json:"resolved"`
		Announced     bool                 `json:"announced"`
		Connected     bool                 `json:"connected"`
		Scanned       bool                 `json:"scanned"`
		Latency       uint64               `json:"latency"`
		ResolvedIPs   []string             `json:"resolved_ips"`
		Settings      HostExternalSettings `json:"external_settings"`
		Announcements []Announcement       `json:"announcements"`
		Errors        []ScanError          `json:"errors"`
	}

	//ScanError an error connecting to the host
	ScanError struct {
		Severity    string   `json:"severity"`
		Type        string   `json:"type"`
		Message     string   `json:"message"`
		Reasons     []string `json:"reasons"`
		Resolutions []string `json:"resolutions"`
	}

	//AddressUsage AddressUsage
	AddressUsage struct {
		Address   string `json:"address"`
		UsageType string `json:"usage_type"`
	}

	ChainIndex struct {
		ID       string `json:"id"`
		ParentID string `json:"parent_id"`
		Height   uint64 `json:"height"`
	}
)
