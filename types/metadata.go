package types

type Metadata []Scenario

type Scenario struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Difficulty  string   `json:"difficulty"`
	Version     string   `json:"version"`
	Tags        []string `json:"tags"`
}
