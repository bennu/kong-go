package client

import (
	"context"
	"fmt"

	"github.com/bennu/kong-go/types"
)

func (cli *Client) APIDelete(ctx context.Context, name string) error {
	resp, err := cli.delete(ctx, fmt.Sprintf("/apis/%s", name), nil)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var apiErrorResp types.APIErrorResponse
	if err := checkResponseErr(resp, apiErrorResp); err != nil {
		return err
	}

	return nil
}
