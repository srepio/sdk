package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	hc      *http.Client
	Options *ClientOptions
}

type ClientOptions struct {
	Timeout time.Duration
	Token   string
	Url     string
	Scheme  string
}

func NewClient(opts *ClientOptions) *Client {
	if opts.Url == "" {
		opts.Url = "api.srep.io"
	}
	if opts.Scheme == "" {
		opts.Scheme = "https"
	}

	return &Client{
		Options: opts,
		hc: &http.Client{
			Timeout: opts.Timeout,
		},
	}
}

func (c *Client) get(path string, params map[string]string) *http.Request {
	req := &http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme: c.Options.Scheme,
			Host:   c.Options.Url,
			Path:   path,
		},
	}
	if params != nil {
		query := req.URL.Query()
		for key, val := range params {
			query.Add(key, val)
		}
		req.URL.RawQuery = query.Encode()
	}

	return req
}

func (c *Client) post(path string, body []byte) *http.Request {
	req := &http.Request{
		Method: http.MethodPost,
		URL: &url.URL{
			Scheme: c.Options.Scheme,
			Host:   c.Options.Url,
			Path:   path,
		},
		Body: io.NopCloser(bytes.NewReader(body)),
	}
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req
}

func (c *Client) delete(path string, body []byte) *http.Request {
	req := &http.Request{
		Method: http.MethodDelete,
		URL: &url.URL{
			Scheme: c.Options.Scheme,
			Host:   c.Options.Url,
			Path:   path,
		},
		Body: io.NopCloser(bytes.NewReader(body)),
	}
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req
}

func (c *Client) headers(req *http.Request) *http.Request {
	if req.Header == nil {
		req.Header = http.Header{}
	}
	if c.Options.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Options.Token))
	}
	req.Header.Set("User-Agent", "SrepGoSDK/0.1.51")

	return req
}

func do[T any](ctx context.Context, c *http.Client, req *http.Request) (*T, error) {
	req = req.WithContext(ctx)
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	bout, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var out T
	if len(bout) > 0 {
		if err := json.Unmarshal(bout, &out); err != nil {
			return nil, err
		}
	}

	return &out, err
}

type request interface {
	Validate() error
}

func (c *Client) buildRequest(method, path string, req request, params map[string]string) (*http.Request, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var hreq *http.Request
	switch method {
	case http.MethodGet:
		hreq = c.get(path, params)
	case http.MethodPost:
		hreq = c.post(path, body)
	case http.MethodDelete:
		hreq = c.delete(path, body)
	default:
		return nil, errors.New("unsupported http method")
	}
	hreq = c.headers(hreq)

	return hreq, nil
}
