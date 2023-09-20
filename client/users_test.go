package client

import (
	"fmt"
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
