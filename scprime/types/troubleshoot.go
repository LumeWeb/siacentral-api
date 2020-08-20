package types

type (
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
)
