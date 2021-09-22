package types

import (
	"time"

	"go.sia.tech/siad/modules"
	siatypes "go.sia.tech/siad/types"
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
		MaxDownloadBatchSize   uint64            `json:"max_download_batch_size"`
		MaxDuration            uint64            `json:"max_duration"`
		MaxReviseBatchSize     uint64            `json:"max_revise_batch_size"`
		RemainingStorage       uint64            `json:"remaining_storage"`
		SectorSize             uint64            `json:"sector_size"`
		TotalStorage           uint64            `json:"total_storage"`
		WindowSize             uint64            `json:"window_size"`
		RevisionNumber         uint64            `json:"revision_number"`
		BaseRPCPrice           siatypes.Currency `json:"base_rpc_price"`
		Collateral             siatypes.Currency `json:"collateral"`
		MaxCollateral          siatypes.Currency `json:"max_collateral"`
		ContractPrice          siatypes.Currency `json:"contract_price"`
		DownloadBandwidthPrice siatypes.Currency `json:"download_price"`
		SectorAccessPrice      siatypes.Currency `json:"sector_access_price"`
		StoragePrice           siatypes.Currency `json:"storage_price"`
		UploadBandwidthPrice   siatypes.Currency `json:"upload_price"`
	}

	//HostExternalSettings the settings pulled from the host during a scan
	HostExternalSettings struct {
		NetAddress             string            `json:"netaddress"`
		Version                string            `json:"version"`
		AcceptingContracts     bool              `json:"accepting_contracts"`
		MaxDownloadBatchSize   uint64            `json:"max_download_batch_size"`
		MaxDuration            uint64            `json:"max_duration"`
		MaxReviseBatchSize     uint64            `json:"max_revise_batch_size"`
		RemainingStorage       uint64            `json:"remaining_storage"`
		SectorSize             uint64            `json:"sector_size"`
		TotalStorage           uint64            `json:"total_storage"`
		WindowSize             uint64            `json:"window_size"`
		RevisionNumber         uint64            `json:"revision_number"`
		BaseRPCPrice           siatypes.Currency `json:"base_rpc_price"`
		Collateral             siatypes.Currency `json:"collateral"`
		MaxCollateral          siatypes.Currency `json:"max_collateral"`
		ContractPrice          siatypes.Currency `json:"contract_price"`
		DownloadBandwidthPrice siatypes.Currency `json:"download_price"`
		SectorAccessPrice      siatypes.Currency `json:"sector_access_price"`
		StoragePrice           siatypes.Currency `json:"storage_price"`
		UploadBandwidthPrice   siatypes.Currency `json:"upload_price"`
	}
)
