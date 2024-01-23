package client

import (
	"context"
	"net/http"

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
	hreq, err := c.buildRequest(http.MethodPost, "/auth", req, nil)
	if err != nil {
		return nil, err
	}

	return do[CreateUserResponse](ctx, c.hc, hreq)
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
	User             *types.User        `json:"user,omitempty"`
	Token            string             `json:"token,omitempty"`
	Details          *types.UserDetails `json:"details,omitempty"`
	MFARequired      bool               `json:"mfa_required,omitempty"`
	AuthenticationID string             `json:"authentication_id,omitempty"`
}

func (c *Client) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	hreq, err := c.buildRequest(http.MethodPost, "/auth/login", req, nil)
	if err != nil {
		return nil, err
	}

	return do[LoginResponse](ctx, c.hc, hreq)
}

type VerifyMFARequest struct {
	AuthenticationID string `json:"authentication_id"`
	Code             string `json:"code"`
}

func (r VerifyMFARequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.AuthenticationID, validation.Required),
		validation.Field(&r.Code, validation.Required),
	)
}

func (c *Client) VerifyMFA(ctx context.Context, req *VerifyMFARequest) (*LoginResponse, error) {
	hreq, err := c.buildRequest(http.MethodPost, "/auth/mfa/verify", req, nil)
	if err != nil {
		return nil, err
	}

	return do[LoginResponse](ctx, c.hc, hreq)
}

type MeRequest struct{}

func (r MeRequest) Validate() error {
	return nil
}

type MeResponse struct {
	User    *types.User        `json:"user"`
	Details *types.UserDetails `json:"details"`
}

func (c *Client) Me(ctx context.Context, req *MeRequest) (*MeResponse, error) {
	hreq, err := c.buildRequest(http.MethodGet, "/auth/me", req, nil)
	if err != nil {
		return nil, err
	}

	return do[MeResponse](ctx, c.hc, hreq)
}

type GetApiTokensRequest struct{}

func (r GetApiTokensRequest) Validate() error {
	return nil
}

type GetApiTokensResponse struct {
	Tokens []types.ApiToken `json:"tokens"`
}

func (c *Client) GetApiTokens(ctx context.Context, req *GetApiTokensRequest) (*GetApiTokensResponse, error) {
	hreq, err := c.buildRequest(http.MethodGet, "/auth/tokens", req, nil)
	if err != nil {
		return nil, err
	}

	return do[GetApiTokensResponse](ctx, c.hc, hreq)
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
	hreq, err := c.buildRequest(http.MethodPost, "/auth/tokens", req, nil)
	if err != nil {
		return nil, err
	}

	return do[CreateApiTokenResponse](ctx, c.hc, hreq)
}

type DeleteApiTokenRequest struct {
	Name string `json:"name"`
}

func (r DeleteApiTokenRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required),
	)
}

type DeleteApiTokenResponse struct{}

func (c *Client) DeleteApiToken(ctx context.Context, req *DeleteApiTokenRequest) (*DeleteApiTokenResponse, error) {
	hreq, err := c.buildRequest(http.MethodDelete, "/auth/tokens", req, nil)
	if err != nil {
		return nil, err
	}

	return do[DeleteApiTokenResponse](ctx, c.hc, hreq)
}

type ConfirmPasswordRequest struct {
	Password string `json:"password"`
}

func (r ConfirmPasswordRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Password, validation.Required),
	)
}

type ConfirmPasswordResponse struct {
	Confirmed bool `json:"confirmed"`
}

func (c *Client) ConfirmPassword(ctx context.Context, req *ConfirmPasswordRequest) (*ConfirmPasswordResponse, error) {
	hreq, err := c.buildRequest(http.MethodPost, "/auth/password/confirm", req, nil)
	if err != nil {
		return nil, err
	}

	return do[ConfirmPasswordResponse](ctx, c.hc, hreq)
}

type UpdatePasswordRequest struct {
	CurrentPassword         string `json:"current_password"`
	NewPassword             string `json:"new_password"`
	NewPasswordConfirmation string `json:"new_password_confirmation"`
}

func (r UpdatePasswordRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.CurrentPassword, validation.Required),
		validation.Field(&r.NewPassword, validation.Required),
		validation.Field(&r.NewPasswordConfirmation, validation.Required),
	)
}

type UpdatePasswordResponse struct {
	Updated bool `json:"updated"`
}

func (c *Client) UpdatePassword(ctx context.Context, req *UpdatePasswordRequest) (*UpdatePasswordResponse, error) {
	hreq, err := c.buildRequest(http.MethodPost, "/auth/password/update", req, nil)
	if err != nil {
		return nil, err
	}

	return do[UpdatePasswordResponse](ctx, c.hc, hreq)
}

type ConfigureMFARequest struct {
	Provider types.MFAType `json:"provider"`
}

func (r ConfigureMFARequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Provider, validation.Required, validation.In(types.TOTP)),
	)
}

type ConfigureMFAResponse struct {
	Data *types.MFAData `json:"mfa_data"`
}

func (c *Client) ConfigureMFA(ctx context.Context, req *ConfigureMFARequest) (*ConfigureMFAResponse, error) {
	hreq, err := c.buildRequest(http.MethodPost, "/auth/mfa/configure", req, nil)
	if err != nil {
		return nil, err
	}

	return do[ConfigureMFAResponse](ctx, c.hc, hreq)
}

type RemoveMFARequest struct{}

func (r RemoveMFARequest) Validate() error {
	return nil
}

type RemoveMFAResponse struct{}

func (c *Client) RemoveMFA(ctx context.Context, req *RemoveMFARequest) (*RemoveMFAResponse, error) {
	hreq, err := c.buildRequest(http.MethodDelete, "/auth/mfa", req, nil)
	if err != nil {
		return nil, err
	}
	return do[RemoveMFAResponse](ctx, c.hc, hreq)
}

type LogoutRequest struct{}

func (r LogoutRequest) Validate() error {
	return nil
}

type LogoutResponse struct{}

func (c *Client) Logout(ctx context.Context, req *LogoutRequest) (*LogoutResponse, error) {
	hreq, err := c.buildRequest(http.MethodPost, "/auth/logout", req, nil)
	if err != nil {
		return nil, err
	}
	return do[LogoutResponse](ctx, c.hc, hreq)
}
