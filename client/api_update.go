package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bennu/kong-go/types"
)

func (cli *Client) APIUpdate(ctx context.Context, key string, options types.APIUpdate) (types.API, error) {
	var api types.API

	resp, err := cli.patch(ctx, fmt.Sprintf("/apis/%s", key), options)
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
