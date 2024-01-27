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
				Scenario: "bongo",
			},
			passes: true,
		},
		{
			request: StartPlayRequest{},
			passes:  false,
		},
		{
			request: StartPlayRequest{},
			passes:  false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("start_play_validation_%s_%t", c.request.Scenario, c.passes), func(t *testing.T) {
			err := c.request.Validate()
			if c.passes {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestCheckPlayRequestValidation(t *testing.T) {
	type testCase struct {
		request CheckPlayRequest
		passes  bool
	}

	cases := []testCase{
		{
			request: CheckPlayRequest{
				ID: uuid.NewString(),
			},
			passes: true,
		},
		{
			request: CheckPlayRequest{
				ID: "bongo",
			},
			passes: false,
		},
		{
			request: CheckPlayRequest{},
			passes:  false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("check_play_validation_%s_%t", c.request.ID, c.passes), func(t *testing.T) {
			err := c.request.Validate()
			if c.passes {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestCancelPlayRequestValidation(t *testing.T) {
	type testCase struct {
		request CancelPlayRequest
		passes  bool
	}

	cases := []testCase{
		{
			request: CancelPlayRequest{
				ID: uuid.NewString(),
			},
			passes: true,
		},
		{
			request: CancelPlayRequest{
				ID: "bongo",
			},
			passes: false,
		},
		{
			request: CancelPlayRequest{},
			passes:  false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("cancel_play_validation_%s_%t", c.request.ID, c.passes), func(t *testing.T) {
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

func TestGetShellRequestValidation(t *testing.T) {
	type testCase struct {
		request GetShellRequest
		passes  bool
	}

	cases := []testCase{
		{
			request: GetShellRequest{},
			passes:  false,
		},
		{
			request: GetShellRequest{
				ID: uuid.NewString(),
			},
			passes: true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("get_shell_validation_%t", c.passes), func(t *testing.T) {
			err := c.request.Validate()
			if c.passes {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestFindPlayRequestValidation(t *testing.T) {
	type testCase struct {
		request GetPlayRequest
		passes  bool
	}

	cases := []testCase{
		{
			request: GetPlayRequest{},
			passes:  false,
		},
		{
			request: GetPlayRequest{
				ID: uuid.NewString(),
			},
			passes: true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("find_play_validation_%t", c.passes), func(t *testing.T) {
			err := c.request.Validate()
			if c.passes {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

// func TestGetShell(t *testing.T) {
// 	cases := []apiTestCase{
// 		{
// 			Url:  "/plays/shell",
// 			Code: http.StatusOK,
// 			Body: `{
//                 "play": {
//                     "id": "6aac65e9-17d2-4a34-8503-490138aa3ed5"
//                 }
//             }`,
// 		},
// 	}

// 	for _, tc := range cases {
// 		t.Run(fmt.Sprintf("get_shell_%v", tc), func(t *testing.T) {
// 			s, c := tc.Prepare(t)
// 			defer s.Close()

// 			_, err := c.GetPlays(context.Background(), &GetPlaysRequest{})
// 			if tc.Errors {
// 				assert.Error(t, err)
// 			} else {
// 				assert.Nil(t, err)
// 			}
// 		})
// 	}
// }
