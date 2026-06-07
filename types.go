package xuiclient

type Sniffing struct {
	Enabled      bool     `json:"enabled"`
	DestOverride []string `json:"destOverride,omitempty"`
	MetadataOnly *bool    `json:"metadataOnly,omitempty"`
	RouteOnly    *bool    `json:"routeOnly,omitempty"`
}

type Finalmask struct {
	QuicParams QuicParams `json:"quicParams"`
}

type QuicParams struct {
	Congestion                  string `json:"congestion"`
	InitStreamReceiveWindow     int64  `json:"initStreamReceiveWindow"`
	MaxStreamReceiveWindow      int64  `json:"maxStreamReceiveWindow"`
	InitConnectionReceiveWindow int64  `json:"initConnectionReceiveWindow"`
	MaxConnectionReceiveWindow  int64  `json:"maxConnectionReceiveWindow"`
	KeepAlivePeriod             int64  `json:"keepAlivePeriod"`
	MaxIncomingStreams          int64  `json:"maxIncomingStreams"`
}
