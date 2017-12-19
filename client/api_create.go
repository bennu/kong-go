package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

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

	if resp.StatusCode != 201 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return api, err
		}
		var apiErrorResp types.APIErrorResponse
		if err := json.Unmarshal(body, &apiErrorResp); err != nil {
			return api, fmt.Errorf("Error reading JSON: %v", err)
		}

		return api, fmt.Errorf("Error response from kong: %s", apiErrorResp)
	}

	err = json.NewDecoder(resp.Body).Decode(&api)
	return api, err
}
