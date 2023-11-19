package types

type PlayStatus string

const (
	PlayRunning   PlayStatus = "RUNNING"
	PlayCompleted PlayStatus = "COMPLETED"
	PlayFailed    PlayStatus = "FAILED"
)

type Play struct{}
