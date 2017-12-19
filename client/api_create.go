package client

import (
	"context"
	"encoding/json"

	"github.com/bennu/kong-go/types"
)

func (cli *Client) APICreate(ctx context.Context, name string, options types.APICreate) (types.API, error) {
	apiCreateRequest := types.APICreateRequest{
		APICreate: options,
		Name:      name,
	}

	var api types.API

	resp, err := cli.post(ctx, "/apis", apiCreateRequest)
	if err != nil {
		return api, err
	}

	defer resp.Body.Close()

	var apiErrorResp types.APIErrorResponse
	if err := checkResponseErr(resp, apiErrorResp); err != nil {
		return api, err
	}

	err = json.NewDecoder(resp.Body).Decode(&api)
	return api, err
}
