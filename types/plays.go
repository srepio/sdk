package types

type PlayStatus string

const (
	PlayPending   PlayStatus = "PENDING"
	PlayRunning   PlayStatus = "RUNNING"
	PlayCompleted PlayStatus = "COMPLETED"
	PlayFailed    PlayStatus = "FAILED"
	PlayCancelled PlayStatus = "CANCELLED"
)

type Play struct {
	ID         string     `json:"id"`
	UserID     string     `json:"user_id"`
	Scenario   string     `json:"scenario"`
	Status     PlayStatus `json:"status"`
	CreatedAt  int64      `json:"created_at"`
	FinishedAt *int64     `json:"finished_at,omitempty"`
	UpdatedAt  int64      `json:"updated_at"`
}
