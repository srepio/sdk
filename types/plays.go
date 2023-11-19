package types

type PlayStatus string

const (
	PlayRunning   PlayStatus = "RUNNING"
	PlayCompleted PlayStatus = "COMPLETED"
	PlayFailed    PlayStatus = "FAILED"
)

type Play struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	Scenario  string     `json:"scenario"`
	Driver    string     `json:"driver"`
	Status    PlayStatus `json:"status"`
	CreatedAt int64      `json:"created_at"`
	UpdatedAt int64      `json:"updated_at"`
}
