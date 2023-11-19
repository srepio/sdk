package client

import (
	"context"
	"fmt"
	"net/http"
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

func TestGetPlaysRequestValidation(t *testing.T) {
	type testCase struct {
		request GetPlaysRequest
		passes  bool
	}

	cases := []testCase{
		{
			request: GetPlaysRequest{},
			passes:  true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("get_plays_validation_%t", c.passes), func(t *testing.T) {
			err := c.request.Validate()
			if c.passes {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestGetPlays(t *testing.T) {
	cases := []apiTestCase{
		{
			Url:  "/plays",
			Code: http.StatusOK,
			Body: `{
                "play": {
                    "id": "6aac65e9-17d2-4a34-8503-490138aa3ed5"
                }
            }`,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("get_plays_%v", tc), func(t *testing.T) {
			s, c := tc.Prepare(t)
			defer s.Close()

			_, err := c.GetPlays(context.Background(), &GetPlaysRequest{})
			if tc.Errors {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
