package types

type Metadata []Scenario

type Port struct {
	Host      string `json:"host"`
	Container string `json:"container"`
}

type Volume struct {
	Host      string `json:"host"`
	Container string `json:"container"`
}

type Scenario struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Difficulty  string   `json:"difficulty"`
	Version     string   `json:"version"`
	Tags        []string `json:"tags"`
	Ports       []Port   `json:"ports"`
	Volumes     []Volume `json:"volumes"`
	Privileged  bool     `json:"privileged"`
}
