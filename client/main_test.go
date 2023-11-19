package client

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type apiTestCase struct {
	Url     string
	Body    string
	Code    int
	Headers map[string]string
	Extra   func(*testing.T)
	Errors  bool
}

func (a *apiTestCase) Prepare(t *testing.T) (*httptest.Server, *Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// assert.Equal(t, r.Header.Get("Content-Type"), "application/json")
		// assert.Equal(t, r.Header.Get("Accept"), "application/json")

		for key, val := range a.Headers {
			assert.Equal(t, val, r.Header.Get(key))
		}

		assert.Equal(t, a.Url, r.URL.Path)

		if a.Extra != nil {
			a.Extra(t)
		}

		w.WriteHeader(a.Code)
		w.Write([]byte(a.Body))
	}))

	c := NewClient(&ClientOptions{
		Url:    strings.TrimPrefix(server.URL, "http://"),
		Scheme: "http",
	})
	return server, c
}
