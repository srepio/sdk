package client

import (
	"context"
	"testing"
)

func TestGetScenarios(t *testing.T) {
	client := NewClient(&ClientOptions{
		Url: "api.srep.io",
	})

	_, err := client.GetScenarios(context.Background(), &GetScenariosRequest{})

	if err != nil {
		t.Error(err)
	}
}

func TestScenarioFind(t *testing.T) {
	client := NewClient(&ClientOptions{
		Url: "api.srep.io",
	})

	_, err := client.FindScenario(context.Background(), &FindScenarioRequest{
		Scenario: "mango",
	})

	if err != nil {
		t.Error(err)
	}
}
