package client

import (
	"context"
	"errors"

	"github.com/srepio/sdk/types"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateUserResponse struct {
	User *types.User
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User  *types.User
	Token string
}

func (c *Client) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, errors.New("not implemented yet")
}

func (c *Client) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, errors.New("not implemented yet")
}
