package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/websocket"
	"github.com/srepio/sdk/types"
)

type StartPlayRequest struct {
	Scenario string `json:"scenario" validate:"required"`
}

func (r StartPlayRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Scenario, validation.Required),
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

type CheckPlayRequest struct {
	ID string `json:"id" validate:"required"`
}

func (r CheckPlayRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.ID, validation.Required, validation.Match(regexp.MustCompile(uuidRegex))),
	)
}

type CheckPlayResponse struct {
	Passed bool `json:"passed"`
}

func (c *Client) CheckPlay(ctx context.Context, req *CheckPlayRequest) (*CheckPlayResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	out := &CheckPlayResponse{}
	if _, err := c.post("/plays/check", body, out); err != nil {
		return nil, err
	}
	return out, nil
}

type CancelPlayRequest struct {
	ID string `json:"id" validate:"required"`
}

func (r CancelPlayRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.ID, validation.Required, validation.Match(regexp.MustCompile(uuidRegex))),
	)
}

type CancelPlayResponse struct{}

func (c *Client) CancelPlay(ctx context.Context, req *CancelPlayRequest) (*CancelPlayResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	out := &CancelPlayResponse{}
	if _, err := c.post("/plays/cancel", body, out); err != nil {
		return nil, err
	}
	return out, nil
}

type GetPlaysRequest struct{}

func (r GetPlaysRequest) Validate() error {
	return validation.ValidateStruct(&r)
}

type GetPlaysResponse struct {
	Plays []*types.Play `json:"plays"`
}

func (c *Client) GetPlays(ctx context.Context, req *GetPlaysRequest) (*GetPlaysResponse, error) {
	out := &GetPlaysResponse{}
	if _, err := c.get("/plays", nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

type GetShellRequest struct {
	ID   string `json:"id"`
	Rows uint16 `json:"rows"`
	Cols uint16 `json:"cols"`
}

func (r GetShellRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.ID, validation.Required, validation.Match(regexp.MustCompile(uuidRegex))),
	)
}

func (c *Client) GetShell(ctx context.Context, req *GetShellRequest, stdin io.Reader, stdout io.Writer) error {
	var scheme string
	if c.Options.Scheme == "https" {
		scheme = "wss"
	} else {
		scheme = "ws"
	}

	url := url.URL{
		Scheme: scheme,
		Host:   c.Options.Url,
		Path:   fmt.Sprintf("/plays/%s/shell", req.ID),
	}
	headers := make(http.Header)
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", c.Options.Token))

	ws, resp, err := websocket.DefaultDialer.Dial(url.String(), headers)
	if err != nil {
		if resp.StatusCode == http.StatusTooEarly {
			return ErrTooEarly
		}
		return fmt.Errorf("%v: %d", err, resp.StatusCode)
	}
	defer ws.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_, message, err := ws.ReadMessage()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err) {
						return
					}
					fmt.Println(err)
					return
				}

				msg := &types.TerminalMessage{}
				if err := json.Unmarshal(message, msg); err != nil {
					fmt.Println(err)
					return
				}

				if msg.Type == types.Ping {
					if err := ws.WriteJSON(types.TerminalMessage{Type: types.Pong}); err != nil {
						fmt.Println(err)
						return
					}
				} else {
					stdout.Write([]byte(msg.Content))
				}
			}
		}
	}()

	go func() {
		buffer := make([]byte, 1024)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n, err := stdin.Read(buffer)
				if err != nil {
					fmt.Println(err)
					return
				}
				if n == 0 {
					continue
				}
				data := make([]byte, n)
				copy(data, buffer)
				msg := &types.TerminalMessage{
					Type:    types.Input,
					Content: string(data),
				}
				if err := ws.WriteJSON(msg); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}()

	<-done
	return nil
}

type GetActivePlayRequest struct{}

func (r GetActivePlayRequest) Validate() error {
	return validation.ValidateStruct(&r)
}

type GetActivePlayResponse struct {
	Play *types.Play `json:"play"`
}

func (c *Client) GetActivePlay(ctx context.Context, req *GetActivePlayRequest) (*GetActivePlayResponse, error) {
	out := &GetActivePlayResponse{}
	if _, err := c.get("/plays/active", nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

type FindPlayRequest struct {
	ID string `json:"id"`
}

func (r FindPlayRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.ID, validation.Required, validation.Match(regexp.MustCompile(uuidRegex))),
	)
}

type FindPlayResponse struct {
	Play *types.Play `json:"play"`
}

func (c *Client) FindPlay(ctx context.Context, req *FindPlayRequest) (*FindPlayResponse, error) {
	out := &FindPlayResponse{}
	if _, err := c.get(fmt.Sprintf("/plays/%s", req.ID), nil, out); err != nil {
		return nil, err
	}
	return out, nil
}
