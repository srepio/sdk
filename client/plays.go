package client

import (
	"context"
	"encoding/json"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type StartPlayRequest struct {
	Scenario string `json:"scenario" validate:"required"`
	Driver   string `json:"driver" validate:"required"`
}

func (r StartPlayRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Scenario, validation.Required),
		validation.Field(&r.Driver, validation.Required),
	)
}

type StartPlayResponse struct {
	ID string `json:"id"`
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
	ID string `json:"id" validate:"required"`
}

func (r CompletePlayRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.ID, validation.Required, validation.Match(regexp.MustCompile(uuidRegex))),
	)
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

type FailedPlayRequest struct {
	ID string `json:"id" validate:"required"`
}

func (r FailedPlayRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.ID, validation.Required, validation.Match(regexp.MustCompile(uuidRegex))),
	)
}

type FailedPlayResponse struct{}

func (c *Client) FailPlay(ctx context.Context, req *FailedPlayRequest) (*FailedPlayResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	out := &FailedPlayResponse{}
	if _, err := c.post("/plays/fail", body, out); err != nil {
		return nil, err
	}
	return out, nil
}
