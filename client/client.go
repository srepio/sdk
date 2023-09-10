package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	hc    *http.Client
	Url   string
	Token string
}

type ClientOptions struct {
	Timeout time.Duration
	Token   string
	Url     string
}

func NewClient(opts *ClientOptions) *Client {
	if opts.Url == "" {
		opts.Url = "api.srep.io"
	}

	return &Client{
		Url:   opts.Url,
		Token: opts.Token,
		hc: &http.Client{
			Timeout: opts.Timeout,
		},
	}
}

func (c *Client) get(path string, params map[string]string) (*http.Response, error) {
	req := &http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme: "https",
			Host:   c.Url,
			Path:   path,
		},
	}
	query := req.URL.Query()
	for key, val := range params {
		query.Add(key, val)
	}
	req.URL.RawQuery = query.Encode()

	return c.request(req)
}

func (c *Client) post(path string, body []byte) (*http.Response, error) {
	req := &http.Request{
		Method: http.MethodPost,
		URL: &url.URL{
			Scheme: "https",
			Host:   c.Url,
			Path:   path,
		},
		Body: io.NopCloser(bytes.NewReader(body)),
	}
	req.Header.Add("Content-Type", "application/json")

	return c.request(req)
}

func (c *Client) request(req *http.Request) (*http.Response, error) {
	if c.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}
	resp, err := c.hc.Do(req)
	if resp.StatusCode > 299 {
		err = fmt.Errorf("status code: %d", resp.StatusCode)
	}
	return resp, err
}
