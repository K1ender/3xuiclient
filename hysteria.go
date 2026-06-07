package xuiclient

type HysteriaSettings struct {
	Version        int64      `json:"version"`
	UDPIdleTimeout int64      `json:"udpIdleTimeout"`
	Masquerade     Masquerade `json:"masquerade"`
}

type Masquerade struct {
	Type        string  `json:"type"`
	Dir         string  `json:"dir"`
	URL         string  `json:"url"`
	RewriteHost bool    `json:"rewriteHost"`
	Insecure    bool    `json:"insecure"`
	Content     string  `json:"content"`
	Headers     Headers `json:"headers"`
	StatusCode  int64   `json:"statusCode"`
}

type Headers struct {
}
