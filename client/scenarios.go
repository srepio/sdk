package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/srepio/sdk/types"
)

type GetscenariosResponse struct {
	Scenarios *types.Metadata `json:"scenarios"`
}

// Get all scenarios
func (c *Client) Getscenarios(ctx context.Context) (*GetscenariosResponse, error) {
	resp, err := c.get("/scenarios", map[string]string{})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error getting resource")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	md := &GetscenariosResponse{}
	if err := json.Unmarshal(body, md); err != nil {
		return nil, err
	}

	return md, nil
}

type FindScenarioResponse struct {
	Scenario *types.Scenario `json:"scenario"`
}

// Get all scenarioa metdata
func (c *Client) FindScenario(ctx context.Context, name string) (*FindScenarioResponse, error) {
	resp, err := c.get(fmt.Sprintf("/scenarios/%s", name), map[string]string{})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error getting resource")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	s := &FindScenarioResponse{}
	if err := json.Unmarshal(body, s); err != nil {
		return nil, err
	}

	return s, nil
}
