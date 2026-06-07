package xuiclient

type Inbound struct {
	ID                   int64           `json:"id"`
	Up                   int64           `json:"up"`
	Down                 int64           `json:"down"`
	Total                int64           `json:"total"`
	Remark               string          `json:"remark"`
	Enable               bool            `json:"enable"`
	ExpiryTime           int64           `json:"expiryTime"`
	TrafficReset         string          `json:"trafficReset"`
	LastTrafficResetTime int64           `json:"lastTrafficResetTime"`
	ClientStats          []ClientStat    `json:"clientStats"`
	Listen               string          `json:"listen"`
	Port                 int64           `json:"port"`
	Protocol             string          `json:"protocol"`
	Tag                  string          `json:"tag"`
	Settings             InboundSettings `json:"settings"`
	StreamSettings       StreamSettings  `json:"streamSettings"`
	Sniffing             Sniffing        `json:"sniffing"`
}

type ClientStat struct {
	ID         int64  `json:"id"`
	InboundID  int64  `json:"inboundId"`
	Enable     bool   `json:"enable"`
	Email      string `json:"email"`
	UUID       string `json:"uuid"`
	SubID      string `json:"subId"`
	Up         int64  `json:"up"`
	Down       int64  `json:"down"`
	ExpiryTime int64  `json:"expiryTime"`
	Total      int64  `json:"total"`
	Reset      int64  `json:"reset"`
	LastOnline int64  `json:"lastOnline"`
}

type InboundSettings struct {
	Clients    []XrayClient `json:"clients"`
	Decryption *string      `json:"decryption,omitempty"`
	Encryption *string      `json:"encryption,omitempty"`
	Testseed   []int64      `json:"testseed,omitempty"`
	Version    *int64       `json:"version,omitempty"`
}
