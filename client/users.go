package client

import (
	"context"
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/srepio/sdk/types"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(10, 255)),
	)
}

type CreateUserResponse struct {
	User *types.User `json:"user"`
}

func (c *Client) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	out := &CreateUserResponse{}
	if _, err := c.post("/auth", body, out); err != nil {
		return nil, err
	}
	return out, nil
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r LoginRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}

type LoginResponse struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

func (c *Client) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	out := &LoginResponse{}
	if _, err := c.post("/auth/login", body, out); err != nil {
		return nil, err
	}
	return out, nil
}

type MeResponse struct {
	User *types.User `json:"user"`
}

func (c *Client) Me(ctx context.Context) (*MeResponse, error) {
	me := &MeResponse{}
	if _, err := c.get("/auth/me", map[string]string{}, me); err != nil {
		return nil, err
	}
	return me, nil
}

type GetApiTokensRequest struct{}

type GetApiTokensResponse struct {
	Tokens []types.ApiToken
}

func (c *Client) GetApiTokens(ctx context.Context, req *GetApiTokensRequest) (*GetApiTokensResponse, error) {
	resp := &GetApiTokensResponse{}
	if _, err := c.get("/auth/tokens", map[string]string{}, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type CreateApiTokenRequest struct {
	Name string `json:"name"`
}

func (r CreateApiTokenRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required),
	)
}

type CreateApiTokenResponse struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

func (c *Client) CreateApiToken(ctx context.Context, req *CreateApiTokenRequest) (*CreateApiTokenResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp := &CreateApiTokenResponse{}
	if _, err := c.post("/auth/tokens", body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type DeleteApiTokenRequest struct {
	Name string `json:"name"`
}

type DeleteApiTokenResponse struct{}

func (c *Client) DeleteApiToken(ctx context.Context, req *DeleteApiTokenRequest) (*DeleteApiTokenResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp := &DeleteApiTokenResponse{}
	if _, err := c.delete("/auth/tokens", body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
