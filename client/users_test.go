package client

import (
	"context"
	"testing"
)

func TestMe(t *testing.T) {
	client := NewClient(&ClientOptions{
		Url: "api.srep.io",
	})

	_, err := client.Me(context.Background())

	if err != nil {
		t.Error(err)
	}
}
