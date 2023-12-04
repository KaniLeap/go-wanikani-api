package wk

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/KaniLeap/go-wanikani-api/data"
	"github.com/KaniLeap/go-wanikani-api/options"
)

type Client struct {
	Token   string
	BaseUrl string
}

func NewClient(token string) *Client {
	return &Client{
		Token:   token,
		BaseUrl: "https://api.wanikani.com/v2/",
	}
}

func initialRequest[T any](client *Client, url string) (*data.Collection[T], error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+client.Token)

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

func requestWithOpts[T any](client *Client, endpoint string, opts ...options.Option) (*data.Collection[T], error) {
	cfg, err := options.NewRequestConfig(opts...)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(client.BaseUrl)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += endpoint
	baseUrl.RawQuery = cfg.QueryParams.Encode()

	return initialRequest[T](client, baseUrl.String())
}
