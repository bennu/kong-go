package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/bennu/kong-go/types"
)

func (cli *Client) APIList(ctx context.Context, options types.APIListOptions) (types.APIList, error) {
	query := url.Values{}

	if options.Id != "" {
		query.Set("id", options.Id)
	}

	if options.Name != "" {
		query.Set("name", options.Name)
	}

	var apis types.APIList

	resp, err := cli.get(ctx, "/apis", query)
	if err != nil {
		return apis, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&apis)
	return apis, err
}
