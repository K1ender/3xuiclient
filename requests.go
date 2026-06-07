package xuiclient

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
