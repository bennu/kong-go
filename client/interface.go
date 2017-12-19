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
	APICreate(ctx context.Context, name string, options types.APICreate) (types.API, error)
	APIUpdate(ctx context.Context, name string, options types.APIUpdate) (types.API, error)
	APILookup(ctx context.Context, name string) (types.API, error)
	APIDelete(ctx context.Context, name string) error
}
