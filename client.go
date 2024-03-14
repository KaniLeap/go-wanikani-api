package wk

import (
	"context"
	"github.com/carlmjohnson/requests"
	"net/http"
)

type Client struct {
	client *http.Client
	token  string
}

var ctx = context.Background()

func NewClient(token string) *Client {
	return &Client{
		token: token,
	}
}

func (c *Client) getRequest() *requests.Builder {
	return requests.
		URL("https://api.wanikani.com/v2/").
		Bearer(c.token)
}

func createRequest[T any](c *Client, path, method string, payload any, opts ...Option) (*T, *requests.Builder, error) {
	cfg := &options{
		params: make(map[string][]string),
	}

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, nil, err
		}
	}

	var result T
	req := c.getRequest().
		Path(path).
		BodyJSON(payload).
		Params(cfg.params).
		Method(method).
		ToJSON(&result)

	err := req.Fetch(ctx)
	if err != nil {
		return nil, nil, err
	}

	return &result, req, nil
}

func resource[T any](c *Client, path, method string, payload any, opts ...Option) (*Resource[T], error) {
	result, _, err := createRequest[Resource[T]](c, path, method, payload, opts...)
	return result, err
}

func paginate[T any](c *Client, path string, opts ...Option) (*Paginate[T], error) {
	response, req, err := createRequest[Collection[T]](c, path, "GET", nil, opts...)
	if err != nil {
		return nil, err
	}

	return &Paginate[T]{
		Data:    *response,
		request: req,
	}, nil
}
