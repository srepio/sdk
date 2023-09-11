package client

import (
	"context"
	"encoding/json"
)

type StartPlayRequest struct {
	Scenario string `json:"scenario" validate:"required"`
	Driver   string `json:"driver" validate:"required"`
}

type StartPlayResponse struct {
	ID int64 `json:"id"`
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
	ID int64 `json:"id" validate:"required"`
}

type CompletePlayResponse struct{}

func (c *Client) CompletePlay(ctx context.Context, req *CompletePlayRequest) (*CompletePlayResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	out := &CompletePlayResponse{}
	if _, err := c.post("/plays/complete", body, out); err != nil {
		return nil, err
	}
	return out, nil
}
