package types

import (
	"time"
)

type (
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
)
