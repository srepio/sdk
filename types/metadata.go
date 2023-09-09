package types

import "github.com/srepio/sdk/types/docker"

type Metadata []Scenario

type Scenario struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Difficulty  string          `json:"difficulty"`
	Version     string          `json:"version"`
	Tags        []string        `json:"tags"`
	Ports       []docker.Port   `json:"ports"`
	Volumes     []docker.Volume `json:"volumes"`
	Privileged  bool            `json:"privileged"`
}
