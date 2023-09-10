package client

import (
	"context"
	"testing"
)

func TestMetadataGet(t *testing.T) {
	client := NewClient(&ClientOptions{
		Url: "api.srep.io",
	})

	_, err := client.GetMetadata(context.Background())

	if err != nil {
		t.Error(err)
	}
}
