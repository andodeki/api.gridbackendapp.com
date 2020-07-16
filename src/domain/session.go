package domain

type DeviceID string

var NilDeviceID DeviceID

// Session represent the users Session
type Session struct {
	UserID       UserID   `db:"user_id,omitempty" `
	DeviceID     DeviceID `db:"device_id,omitempty"`
	RefreshToken string   `db:"refresh_token,omitempty"`
	ExpiresAt    int64    `db:"expires_at"`
}

// Session is used to represents data sent in json body with request
type SessionData struct {
	DeviceID DeviceID `db:"deviceID,omitempty"`
}
