package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserRequestValidation(t *testing.T) {
	type testCase struct {
		request CreateUserRequest
		passes  bool
	}

	cases := []testCase{
		{
			request: CreateUserRequest{
				Name:     gofakeit.Name(),
				Email:    gofakeit.Email(),
				Password: gofakeit.Password(true, true, true, true, false, 10),
			},
			passes: true,
		},
		{
			request: CreateUserRequest{
				Name:     gofakeit.Name(),
				Email:    gofakeit.Email(),
				Password: gofakeit.Password(true, true, true, true, false, 9),
			},
			passes: false,
		},
		{
			request: CreateUserRequest{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
			},
			passes: false,
		},
		{
			request: CreateUserRequest{
				Name:     gofakeit.Name(),
				Password: gofakeit.Password(true, true, true, true, false, 10),
			},
			passes: false,
		},
		{
			request: CreateUserRequest{
				Email:    gofakeit.Email(),
				Password: gofakeit.Password(true, true, true, true, false, 10),
			},
			passes: false,
		},
		{
			request: CreateUserRequest{
				Name:     gofakeit.Name(),
				Email:    gofakeit.Email(),
				Password: gofakeit.Password(true, true, true, true, false, 9),
			},
			passes: false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("create_user_validation_%v_%t", c.request, c.passes), func(t *testing.T) {
			err := c.request.Validate()
			if c.passes {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestLoginRequestValidation(t *testing.T) {
	type testCase struct {
		request LoginRequest
		passes  bool
	}

	cases := []testCase{
		{
			request: LoginRequest{
				Email:    gofakeit.Name(),
				Password: gofakeit.Email(),
			},
			passes: true,
		},
		{
			request: LoginRequest{
				Email: gofakeit.Name(),
			},
			passes: false,
		},
		{
			request: LoginRequest{
				Password: gofakeit.Email(),
			},
			passes: false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("login_validation_%v_%t", c.request, c.passes), func(t *testing.T) {
			err := c.request.Validate()
			if c.passes {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	cases := []apiTestCase{
		{
			Url:  "/auth/login",
			Code: http.StatusOK,
			Body: `{ "user": {
                    "id": "f41ffd64-e726-4f5b-8b59-190be848014d",
                    "name": "Henry Whitaker",
                    "email": "henrywhitaker3@outlook.com"
                },
                "token": "bongo"
            }`,
		},
		{
			Url:    "/auth/login",
			Code:   http.StatusUnauthorized,
			Body:   `{}`,
			Errors: true,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("login_%v", tc), func(t *testing.T) {
			s, c := tc.Prepare(t)
			defer s.Close()

			resp, err := c.Login(context.Background(), &LoginRequest{})
			fmt.Println(resp)
			if tc.Errors {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	cases := []apiTestCase{
		{
			Url:  "/auth",
			Code: http.StatusOK,
			Body: `{
                "user": {
                    "id": "6aac65e9-17d2-4a34-8503-490138aa3ed5",
                    "name": "Henry Whitaker",
                    "email": "henrywhitaker34@outlook.com"
                }
            }`,
		},
		{
			Url:  "/auth",
			Code: http.StatusUnprocessableEntity,
			Body: `{
                "message": "Request failed validation",
                "errors": {
                    "password": "cannot be blank"
                }
            }`,
			Errors: true,
		},
		{
			Url:  "/auth",
			Code: http.StatusUnprocessableEntity,
			Body: `{
                "message": "Request failed validation",
                "errors": {
                    "email": "user henrywhitaker3@outlook.com already exists"
                }
            }`,
			Errors: true,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("register_%v", tc), func(t *testing.T) {
			s, c := tc.Prepare(t)
			defer s.Close()

			resp, err := c.CreateUser(context.Background(), &CreateUserRequest{})
			fmt.Println(resp)
			if tc.Errors {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
