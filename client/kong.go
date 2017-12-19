package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

func (cli *Client) post(ctx context.Context, path string, obj interface{}) (*http.Response, error) {
	body, err := encodeBody(obj)
	if err != nil {
		return nil, err
	}

	return cli.request(ctx, http.MethodPost, path, nil, body)
}

func (cli *Client) get(ctx context.Context, path string, query url.Values) (*http.Response, error) {
	return cli.request(ctx, http.MethodGet, path, query, nil)
}

func (cli *Client) delete(ctx context.Context, path string, query url.Values) (*http.Response, error) {
	return cli.request(ctx, http.MethodDelete, path, query, nil)
}

func (cli *Client) request(ctx context.Context, method, path string, query url.Values, body io.Reader) (*http.Response, error) {
	p, err := cli.apiPath(path, query)
	if err != nil {
		return nil, err
	}
	fmt.Println(body, method, p)

	req, err := http.NewRequest(method, p, body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := ctxhttp.Do(ctx, cli.client, req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resp, err
}

func (cli *Client) apiPath(path string, query url.Values) (string, error) {
	url, err := url.Parse(path)
	if err != nil {
		return "", err
	}

	ref := cli.BaseURL.ResolveReference(url)

	if query != nil {
		ref.RawQuery = query.Encode()
	}
	return ref.String(), nil
}

func encodeBody(obj interface{}) (*bytes.Buffer, error) {
	buff := bytes.NewBuffer(nil)
	if obj != nil {
		if err := json.NewEncoder(buff).Encode(obj); err != nil {
			return nil, err
		}
	}
	return buff, nil
}

func checkResponseErr(resp *http.Response, obj interface{}) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &obj); err != nil {
		return fmt.Errorf("Error reading JSON: %v", err)
	}

	return fmt.Errorf("Error response from kong: %s", obj)

}
