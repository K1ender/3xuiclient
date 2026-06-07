package xuiclient

type Response[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"obj"`
}

type ClientResponse struct {
	Client     ClientRes `json:"client"`
	InboundIDS []int64   `json:"inboundIds"`
}

type ClientRes struct {
	ID         int64  `json:"id"`
	Email      string `json:"email"`
	SubID      string `json:"subId"`
	UUID       string `json:"uuid"`
	Password   string `json:"password"`
	Auth       string `json:"auth"`
	Flow       string `json:"flow"`
	Security   string `json:"security"`
	LimitIP    int64  `json:"limitIp"`
	TotalGB    int64  `json:"totalGB"`
	ExpiryTime int64  `json:"expiryTime"`
	Enable     bool   `json:"enable"`
	TgID       int64  `json:"tgId"`
	Group      string `json:"group"`
	Comment    string `json:"comment"`
	Reset      int64  `json:"reset"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
	Reverse    any    `json:"reverse"`
}
