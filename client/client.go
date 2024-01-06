package client

import (
	"bytes"
	"encoding/json"
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

func (c *Client) get(path string, params map[string]string, data any) (*http.Response, error) {
	req := &http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme: c.Options.Scheme,
			Host:   c.Options.Url,
			Path:   path,
		},
	}
	query := req.URL.Query()
	for key, val := range params {
		query.Add(key, val)
	}
	req.URL.RawQuery = query.Encode()

	return c.request(req, data)
}

func (c *Client) post(path string, body []byte, data any) (*http.Response, error) {
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

	return c.request(req, data)
}

func (c *Client) delete(path string, body []byte, data any) (*http.Response, error) {
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

	return c.request(req, data)
}

func (c *Client) request(req *http.Request, data any) (*http.Response, error) {
	if req.Header == nil {
		req.Header = http.Header{}
	}
	if c.Options.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Options.Token))
	}
	resp, err := c.hc.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return resp, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	bout, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bout, data); err != nil {
		return resp, err
	}

	return resp, err
}
