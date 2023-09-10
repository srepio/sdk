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

type GetMetadataResponse struct {
	Scenarios *types.Metadata `json:"scenarios"`
}

// Get all scenarioa metdata
func (c *Client) GetMetadata(ctx context.Context) (*GetMetadataResponse, error) {
	resp, err := c.get("/metadata", map[string]string{})
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

	md := &types.Metadata{}
	if err := json.Unmarshal(body, md); err != nil {
		return nil, err
	}

	return &GetMetadataResponse{
		Scenarios: md,
	}, nil
}

type GetScenarioResponse struct {
	Scenario *types.Scenario `json:"scenario"`
}

// Get all scenarioa metdata
func (c *Client) FindScenario(ctx context.Context, name string) (*GetScenarioResponse, error) {
	resp, err := c.get(fmt.Sprintf("/metadata/%s", name), map[string]string{})
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

	s := &types.Scenario{}
	if err := json.Unmarshal(body, s); err != nil {
		return nil, err
	}

	return &GetScenarioResponse{
		Scenario: s,
	}, nil
}
