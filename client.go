package wk

import (
	"context"
	"errors"
	"net/http"
	"strconv"
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

	attempts := 0
	for {
		var headers http.Header
		err := req.
			CheckStatus(200, 299).
			CopyHeaders(headers).
			ErrorJSON(&apiErr).
			Fetch(ctx)
		if err == nil {
			break
		}
		var se *requests.ResponseError
		if errors.As(err, &se) {
			reset := headers.Get("RateLimit-Reset")
			if reset != "" {
				if sec, perr := strconv.ParseInt(reset, 10, 64); perr == nil {
					wait := time.Until(time.Unix(sec, 0))
					if wait > 0 {
						timer := time.NewTimer(wait)
						select {
						case <-ctx.Done():
							if !timer.Stop() {
								<-timer.C
							}
							return nil, nil, ctx.Err()
						case <-timer.C:
						}
					}
				}
			}
			if se.StatusCode == http.StatusTooManyRequests {
				continue
			}
			if se.StatusCode >= 500 && se.StatusCode <= 599 {
				if attempts < 2 {
					attempts++
					time.Sleep(time.Duration(attempts) * 500 * time.Millisecond)
					continue
				}
			}
		}
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
