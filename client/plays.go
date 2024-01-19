package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/websocket"
	"github.com/srepio/sdk/types"
	"golang.org/x/crypto/ssh/terminal"
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
	hreq, err := c.buildRequest(http.MethodPost, "/plays", req, nil)
	if err != nil {
		return nil, err
	}

	return do[StartPlayResponse](ctx, c.hc, hreq)
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
	hreq, err := c.buildRequest(http.MethodPost, "/plays/check", req, nil)
	if err != nil {
		return nil, err
	}

	return do[CheckPlayResponse](ctx, c.hc, hreq)
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
	hreq, err := c.buildRequest(http.MethodPost, "/plays/cancel", req, nil)
	if err != nil {
		return nil, err
	}

	return do[CancelPlayResponse](ctx, c.hc, hreq)
}

type GetPlaysRequest struct{}

func (r GetPlaysRequest) Validate() error {
	return validation.ValidateStruct(&r)
}

type GetPlaysResponse struct {
	Plays []*types.Play `json:"plays"`
}

func (c *Client) GetPlays(ctx context.Context, req *GetPlaysRequest) (*GetPlaysResponse, error) {
	hreq, err := c.buildRequest(http.MethodGet, "/plays", req, nil)
	if err != nil {
		return nil, err
	}

	return do[GetPlaysResponse](ctx, c.hc, hreq)
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

func (c *Client) GetShell(ctx context.Context, req *GetShellRequest, stdin *os.File, stdout *os.File) error {
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

	wso, resp, err := websocket.DefaultDialer.Dial(url.String(), headers)
	if err != nil {
		if resp.StatusCode == http.StatusTooEarly {
			return ErrTooEarly
		}
		return fmt.Errorf("%v: %d", err, resp.StatusCode)
	}
	defer wso.Close()
	sock := newWs(wso)

	go func() {
		var oldRows int
		var oldCols int

		for {
			select {
			case <-ctx.Done():
				return
			default:
				cols, rows, err := terminal.GetSize(int(stdout.Fd()))
				if err != nil {
					fmt.Println(err)
					return
				}
				if oldRows != rows || oldCols != cols {
					msg := &types.TerminalMessage{
						Type:    types.Resize,
						Content: fmt.Sprintf("%d,%d", rows, cols),
					}
					sock.Write(msg)
				}
				time.Sleep(time.Second * 1)
			}
		}
	}()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := sock.Read()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err) {
						return
					}
					fmt.Println(err)
					return
				}

				if msg.Type == types.Ping {
					if err := sock.Write(&types.TerminalMessage{Type: types.Pong}); err != nil {
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
				if err := sock.Write(msg); err != nil {
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
	hreq, err := c.buildRequest(http.MethodGet, "/plays/active", req, nil)
	if err != nil {
		return nil, err
	}

	return do[GetActivePlayResponse](ctx, c.hc, hreq)
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
	hreq, err := c.buildRequest(http.MethodPost, fmt.Sprintf("/plays/%s", req.ID), req, nil)
	if err != nil {
		return nil, err
	}

	return do[FindPlayResponse](ctx, c.hc, hreq)
}
