package xuiclient

type TLSSettings struct {
	ServerName              string              `json:"serverName"`
	MinVersion              string              `json:"minVersion"`
	MaxVersion              string              `json:"maxVersion"`
	CipherSuites            string              `json:"cipherSuites"`
	RejectUnknownSni        bool                `json:"rejectUnknownSni"`
	DisableSystemRoot       bool                `json:"disableSystemRoot"`
	EnableSessionResumption bool                `json:"enableSessionResumption"`
	Certificates            []Certificate       `json:"certificates"`
	Alpn                    []string            `json:"alpn"`
	EchServerKeys           string              `json:"echServerKeys"`
	EchForceQuery           *string             `json:"echForceQuery,omitempty"`
	Settings                TLSSettingsSettings `json:"settings"`
}

type TLSSettingsSettings struct {
	Fingerprint          string `json:"fingerprint"`
	EchConfigList        string `json:"echConfigList"`
	PinnedPeerCERTSha256 []any  `json:"pinnedPeerCertSha256,omitempty"`
}

type Certificate struct {
	CertificateFile string `json:"certificateFile"`
	KeyFile         string `json:"keyFile"`
	OneTimeLoading  bool   `json:"oneTimeLoading"`
	Usage           string `json:"usage"`
	BuildChain      bool   `json:"buildChain"`
}
