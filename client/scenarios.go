package client

import (
	"context"
	"fmt"

	"github.com/srepio/sdk/types"
)

type GetscenariosResponse struct {
	Scenarios *types.Metadata `json:"scenarios"`
}

// Get all scenarios
func (c *Client) Getscenarios(ctx context.Context) (*GetscenariosResponse, error) {
	md := &GetscenariosResponse{}
	if _, err := c.get("/scenarios", map[string]string{}, md); err != nil {
		return nil, err
	}

	return md, nil
}

type FindScenarioResponse struct {
	Scenario *types.Scenario `json:"scenario"`
	History  []*types.Play   `json:"history,omitempty"`
}

// Get all scenarioa metdata
func (c *Client) FindScenario(ctx context.Context, name string) (*FindScenarioResponse, error) {
	s := &FindScenarioResponse{}
	if _, err := c.get(fmt.Sprintf("/scenarios/%s", name), map[string]string{}, s); err != nil {
		return nil, err
	}
	return s, nil
}
