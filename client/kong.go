package client

import (
	"context"
	"net/http"
	"net/url"

	"golang.org/x/net/context/ctxhttp"
)

const (
	defaultBaseURL = "http://127.0.0.1:8001"
)

type Client struct {
	client  *http.Client
	BaseURL *url.URL
}

func NewClient(client *http.Client) (*Client, error) {
	if client == nil {
		client = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	return &Client{client: client, BaseURL: baseURL}, nil
}

func (cli *Client) get(ctx context.Context, path string, query url.Values) (*http.Response, error) {
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	ref := cli.BaseURL.ResolveReference(url)

	if query != nil {
		ref.RawQuery = query.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, ref.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := ctxhttp.Do(ctx, cli.client, req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
