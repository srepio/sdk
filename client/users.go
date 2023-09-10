package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/srepio/sdk/types"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateUserResponse struct {
	User *types.User `json:"user"`
}

func (c *Client) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, errors.New("not implemented yet")
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
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

	resp, err := c.post("/auth/login", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	out := &LoginResponse{}
	if err := json.Unmarshal(bs, out); err != nil {
		return nil, err
	}
	return out, nil
}

type MeResponse struct {
	User *types.User `json:"user"`
}

func (c *Client) Me(ctx context.Context) (*MeResponse, error) {
	resp, err := c.get("/auth/me", map[string]string{})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	me := &MeResponse{}
	if err := json.Unmarshal(bs, me); err != nil {
		return nil, err
	}
	return me, nil
}
