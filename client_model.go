package xuiclient

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

type Flow string

const (
	Empty          Flow = ""
	XtlsRprxVision Flow = "xtls-rprx-vision"
)
