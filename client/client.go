package client

import (
	"net/http"
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
		opts.Url = "https://api.srep.io/"
	}

	return &Client{
		Url:   opts.Url,
		Token: opts.Token,
		hc: &http.Client{
			Timeout: opts.Timeout,
		},
	}
}
