package types

import (
	"time"

	siatypes "gitlab.com/NebulousLabs/Sia/types"
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

	//HostInfo information about a host in the database
	HostInfo struct {
		PublicKey   string                `json:"public_key"`
		NetAddress  string                `json:"net_address"`
		ConnectedIP string                `json:"connected_ip"`
		Latency     time.Duration         `json:"latency"`
		LastScan    time.Time             `json:"last_scan"`
		Settings    *HostExternalSettings `json:"settings"`
	}

	//HostDetails additional details about a host from the database
	HostDetails struct {
		HostInfo
		ResolvedIPs   []string       `json:"resolved_ips,omitempty"`
		Announcements []Announcement `json:"announcements,omitempty"`
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
