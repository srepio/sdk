package client

import (
	"context"
	"encoding/json"
)

type StartPlayRequest struct {
	Scenario string `json:"scenario"`
	Driver   string `json:"driver"`
}

type StartPlayResponse struct {
	ID int `json:"string"`
}

func (c *Client) StartPlay(ctx context.Context, req *StartPlayRequest) (*StartPlayResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	out := &StartPlayResponse{}
	if _, err := c.post("/plays", body, out); err != nil {
		return nil, err
	}
	return out, nil
}

type CompletePlayRequest struct {
	ID int `json:"id"`
}

type CompletePlayResponse struct{}

func (c *Client) CompletePlay(ctx context.Context, req *StartPlayRequest) (*StartPlayResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	out := &StartPlayResponse{}
	if _, err := c.post("/plays/complete", body, out); err != nil {
		return nil, err
	}
	return out, nil
}
