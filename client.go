package wk

import (
	"context"
	"net/http"
	"time"

	"github.com/carlmjohnson/requests"
)

type Client struct {
	client  *http.Client
	token   string
	baseURL string
	ua      string
}

func NewClient(token string) *Client {
	return NewClientWithHTTP(nil, token)
}

func NewClientWithHTTP(hc *http.Client, token string) *Client {
	if hc == nil {
		hc = &http.Client{Timeout: 10 * time.Second}
	}
	return &Client{
		client:  hc,
		token:   token,
		baseURL: "https://api.wanikani.com/v2/",
		ua:      "go-wanikani-api (+https://github.com/KaniLeap/go-wanikani-api)",
	}
}

func (c *Client) getRequest() *requests.Builder {
	return requests.
		URL(c.baseURL).
		Client(c.client).
		Bearer(c.token).
		Header("User-Agent", c.ua).
		Header("Accept", "application/json")
}

func createRequest[T any](c *Client, ctx context.Context, path, method string, payload any, opts ...Option) (*T, *requests.Builder, error) {
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
		Params(cfg.params).
		Method(method).
		ToJSON(&result)
	if method != http.MethodGet && payload != nil {
		req = req.BodyJSON(payload)
	}

	var apiErr struct {
		Error string `json:"error"`
		Code  int    `json:"code"`
	}
	err := req.
		CheckStatus(200, 299).
		ErrorJSON(&apiErr).
		Fetch(ctx)
	if err != nil {
		return nil, nil, err
	}

	return &result, req, nil
}

func resource[T any](c *Client, ctx context.Context, path, method string, payload any, opts ...Option) (*Resource[T], error) {
	result, _, err := createRequest[Resource[T]](c, ctx, path, method, payload, opts...)
	return result, err
}

func paginate[T any](c *Client, ctx context.Context, path string, opts ...Option) (*Paginate[T], error) {
	response, req, err := createRequest[Collection[T]](c, ctx, path, "GET", nil, opts...)
	if err != nil {
		return nil, err
	}

	return &Paginate[T]{
		Data:    *response,
		request: req,
		ctx:     ctx,
	}, nil
}
