package wk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/KaniLeap/go-wanikani-api/data"
	"github.com/KaniLeap/go-wanikani-api/options"
)

func makeURL(baseURL, endpoint string, cfg *options.RequestConfig) (string, error) {
	fullURL, err := url.Parse(baseURL + endpoint)
	if err != nil {
		return "", err
	}

	fullURL.RawQuery = cfg.QueryParams.Encode()

	return fullURL.String(), nil
}

func processPayload(payload interface{}) (io.Reader, error) {
	if payload == nil {
		return nil, nil
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(body), nil
}

func newRequest(method, urlStr, token string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	if body != nil {
		req.Header.Add("Content-Type", "application/json; charset=utf-8")
	}

	return req, nil
}

func setupRequest(client *Client, method, endpoint string, opts ...options.Option) (*http.Request, error) {
	cfg, err := options.NewRequestConfig(opts...)
	if err != nil {
		return nil, err
	}

	urlStr, err := makeURL(client.BaseURL, endpoint, cfg)
	if err != nil {
		return nil, err
	}

	body, err := processPayload(cfg.Payload)
	if err != nil {
		return nil, err
	}

	return newRequest(method, urlStr, client.Token, body)
}

func doRequest[T any](req *http.Request) (*data.Collection[T], error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result data.Collection[T]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func request[T any](client *Client, method, endpoint string, opts ...options.Option) (*data.Collection[T], error) {
	req, err := setupRequest(client, method, endpoint, opts...)
	if err != nil {
		return nil, err
	}

	return doRequest[T](req)
}

func paginate[T any](client *Client, endpoint string, opts ...options.Option) (*pagination[T], error) {
	result, err := request[T](client, "GET", endpoint, opts...)
	if err != nil {
		return nil, err
	}

	return &pagination[T]{
		Data:   *result,
		Client: client,
	}, nil
}

func paginateFromURL[T any](client *Client, url string) (*pagination[T], error) {
	req, err := newRequest("GET", url, client.Token, nil)
	if err != nil {
		return nil, err
	}

	result, err := doRequest[T](req)
	if err != nil {
		return nil, err
	}

	return &pagination[T]{
		Data:   *result,
		Client: client,
	}, nil
}
