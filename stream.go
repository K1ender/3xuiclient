package xuiclient

type XhttpSettings struct {
	Path                 string  `json:"path"`
	Host                 string  `json:"host"`
	Headers              Headers `json:"headers"`
	ScMaxBufferedPosts   int64   `json:"scMaxBufferedPosts"`
	ScMaxEachPostBytes   string  `json:"scMaxEachPostBytes"`
	ScStreamUpServerSecs string  `json:"scStreamUpServerSecs"`
	NoSSEHeader          bool    `json:"noSSEHeader"`
	XPaddingBytes        string  `json:"xPaddingBytes"`
	Mode                 string  `json:"mode"`
	XPaddingObfsMode     bool    `json:"xPaddingObfsMode"`
	XPaddingKey          string  `json:"xPaddingKey"`
	XPaddingHeader       string  `json:"xPaddingHeader"`
	XPaddingPlacement    string  `json:"xPaddingPlacement"`
	XPaddingMethod       string  `json:"xPaddingMethod"`
	UplinkHTTPMethod     string  `json:"uplinkHTTPMethod"`
	SessionPlacement     string  `json:"sessionPlacement"`
	SessionKey           string  `json:"sessionKey"`
	SeqPlacement         string  `json:"seqPlacement"`
	SeqKey               string  `json:"seqKey"`
	UplinkDataPlacement  string  `json:"uplinkDataPlacement"`
	UplinkDataKey        string  `json:"uplinkDataKey"`
	UplinkChunkSize      int64   `json:"uplinkChunkSize"`
}

type TCPSettings struct {
	AcceptProxyProtocol bool   `json:"acceptProxyProtocol"`
	Header              Header `json:"header"`
}

type Header struct {
	Type string `json:"type"`
}

type StreamSettings struct {
	Network          string            `json:"network"`
	Security         string            `json:"security"`
	ExternalProxy    []any             `json:"externalProxy,omitempty"`
	TLSSettings      *TLSSettings      `json:"tlsSettings,omitempty"`
	XhttpSettings    *XhttpSettings    `json:"xhttpSettings,omitempty"`
	RealitySettings  *RealitySettings  `json:"realitySettings,omitempty"`
	TCPSettings      *TCPSettings      `json:"tcpSettings,omitempty"`
	HysteriaSettings *HysteriaSettings `json:"hysteriaSettings,omitempty"`
	Finalmask        *Finalmask        `json:"finalmask,omitempty"`
}
