package xuiclient

type RealitySettings struct {
	Show         bool                    `json:"show"`
	Xver         int64                   `json:"xver"`
	Target       string                  `json:"target"`
	ServerNames  []string                `json:"serverNames"`
	PrivateKey   string                  `json:"privateKey"`
	MinClientVer string                  `json:"minClientVer"`
	MaxClientVer string                  `json:"maxClientVer"`
	MaxTimediff  int64                   `json:"maxTimediff"`
	ShortIDS     []string                `json:"shortIds"`
	Mldsa65Seed  string                  `json:"mldsa65Seed"`
	Settings     RealitySettingsSettings `json:"settings"`
}

type RealitySettingsSettings struct {
	PublicKey     string `json:"publicKey"`
	Fingerprint   string `json:"fingerprint"`
	ServerName    string `json:"serverName"`
	SpiderX       string `json:"spiderX"`
	Mldsa65Verify string `json:"mldsa65Verify"`
}
