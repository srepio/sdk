package client

import (
	"fmt"
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

func (c *Client) request(req *http.Request) (*http.Response, error) {
	if c.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}
	return c.hc.Do(req)
}
