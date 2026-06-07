package xuiclient

type Response[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"obj"`
}

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

type XrayClient struct {
	Comment    string  `json:"comment"`
	CreatedAt  int64   `json:"created_at"`
	Email      string  `json:"email"`
	Enable     bool    `json:"enable"`
	ExpiryTime int64   `json:"expiryTime"`
	Flow       *Flow   `json:"flow,omitempty"`
	ID         *string `json:"id,omitempty"`
	LimitIP    int64   `json:"limitIp"`
	Reset      int64   `json:"reset"`
	SubID      string  `json:"subId"`
	TgID       int64   `json:"tgId"`
	TotalGB    int64   `json:"totalGB"`
	UpdatedAt  int64   `json:"updated_at"`
	Password   *string `json:"password,omitempty"`
	Security   *string `json:"security,omitempty"`
	Auth       *string `json:"auth,omitempty"`
}

type Sniffing struct {
	Enabled      bool     `json:"enabled"`
	DestOverride []string `json:"destOverride,omitempty"`
	MetadataOnly *bool    `json:"metadataOnly,omitempty"`
	RouteOnly    *bool    `json:"routeOnly,omitempty"`
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

type TCPSettings struct {
	AcceptProxyProtocol bool   `json:"acceptProxyProtocol"`
	Header              Header `json:"header"`
}

type Header struct {
	Type string `json:"type"`
}

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

type Certificate struct {
	CertificateFile string `json:"certificateFile"`
	KeyFile         string `json:"keyFile"`
	OneTimeLoading  bool   `json:"oneTimeLoading"`
	Usage           string `json:"usage"`
	BuildChain      bool   `json:"buildChain"`
}

type TLSSettingsSettings struct {
	Fingerprint          string `json:"fingerprint"`
	EchConfigList        string `json:"echConfigList"`
	PinnedPeerCERTSha256 []any  `json:"pinnedPeerCertSha256,omitempty"`
}

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

type Flow string

const (
	Empty          Flow = ""
	XtlsRprxVision Flow = "xtls-rprx-vision"
)

type CreateClientRequest struct {
	Client     CreateClient `json:"client"`
	InboundIDS []int64      `json:"inboundIds"`
}

type CreateClient struct {
	// Email login (id).
	Email string `json:"email"`
	// TotalGB in bytes.
	// 1024 * 1024 * 1024 = 1GB.
	// 0 = unlimited.
	TotalGB int64 `json:"totalGB"`
	// ExpiryTime timestamp in millis.
	// time.Now().Add(time.Hour).UnixMilli() - expire in 1 hour.
	// 0 = unlimited.
	ExpiryTime int64 `json:"expiryTime"`
	TgID       int64 `json:"tgId"`
	LimitIP    int64 `json:"limitIp"`
	// Only works when updating a client, not when creating one.
	Enable bool `json:"enable"`
}

type UpdateClientRequest struct {
	CreateClient
}
