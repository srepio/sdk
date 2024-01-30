package client

import (
	"context"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/srepio/sdk/types"
)

type GetScenariosRequest struct{}

func (r GetScenariosRequest) Validate() error {
	return nil
}

type GetScenariosResponse struct {
	Scenarios *types.Metadata `json:"scenarios"`
}

// Get all scenarios
func (c *Client) GetScenarios(ctx context.Context, req *GetScenariosRequest) (*GetScenariosResponse, error) {
	hreq, err := c.buildRequest(http.MethodGet, "/scenarios", req, nil)
	if err != nil {
		return nil, err
	}

	return do[GetScenariosResponse](ctx, c.hc, hreq)
}

type FindScenarioRequest struct {
	Scenario string `json:"scenario" param:"name"`
	Page     int    `json:"page" query:"page"`
}

func (r FindScenarioRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Scenario, validation.Required),
	)
}

type FindScenarioResponse struct {
	Scenario *types.Scenario               `json:"scenario"`
	History  *types.Paginated[*types.Play] `json:"history,omitempty"`
}

// Get all scenarioa metdata
func (c *Client) FindScenario(ctx context.Context, req *FindScenarioRequest) (*FindScenarioResponse, error) {
	if req.Page == 0 {
		req.Page = 1
	}
	hreq, err := c.buildRequest(http.MethodGet, fmt.Sprintf("/scenarios/%s", req.Scenario), req, map[string]string{"page": fmt.Sprintf("%d", req.Page)})
	if err != nil {
		return nil, err
	}

	return do[FindScenarioResponse](ctx, c.hc, hreq)
}
