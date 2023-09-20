package client

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStartPlayRequestValidation(t *testing.T) {
	type testCase struct {
		request StartPlayRequest
		passes  bool
	}

	cases := []testCase{
		{
			request: StartPlayRequest{
				Driver:   "bongo",
				Scenario: "bongo",
			},
			passes: true,
		},
		{
			request: StartPlayRequest{
				Driver: "bongo",
			},
			passes: false,
		},
		{
			request: StartPlayRequest{
				Scenario: "bongo",
			},
			passes: false,
		},
		{
			request: StartPlayRequest{},
			passes:  false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("start_play_validation_%s_%s_%t", c.request.Scenario, c.request.Driver, c.passes), func(t *testing.T) {
			err := c.request.Validate()
			if c.passes {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestCompletePlayRequestValidation(t *testing.T) {
	type testCase struct {
		request CompletePlayRequest
		passes  bool
	}

	cases := []testCase{
		{
			request: CompletePlayRequest{
				ID: uuid.NewString(),
			},
			passes: true,
		},
		{
			request: CompletePlayRequest{
				ID: "bongo",
			},
			passes: false,
		},
		{
			request: CompletePlayRequest{},
			passes:  false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("complete_play_validation_%s_%t", c.request.ID, c.passes), func(t *testing.T) {
			err := c.request.Validate()
			if c.passes {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestFailedPlayRequestValidation(t *testing.T) {
	type testCase struct {
		request FailedPlayRequest
		passes  bool
	}

	cases := []testCase{
		{
			request: FailedPlayRequest{
				ID: uuid.NewString(),
			},
			passes: true,
		},
		{
			request: FailedPlayRequest{
				ID: "bongo",
			},
			passes: false,
		},
		{
			request: FailedPlayRequest{},
			passes:  false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("failed_play_validation_%s_%t", c.request.ID, c.passes), func(t *testing.T) {
			err := c.request.Validate()
			if c.passes {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
