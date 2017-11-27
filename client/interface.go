package client

import (
	"context"
	"github.com/bennu/kong-go/types"
)

type Kong interface {
	API
}

type API interface {
	APIList(ctx context.Context, options types.APIListOptions) (types.APIList, error)
}
