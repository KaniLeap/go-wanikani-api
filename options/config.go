package options

import (
	"net/url"
	"strconv"
	"time"
)

type requestConfig struct {
	QueryParams *url.Values
}

type Option func(*requestConfig) error

func NewRequestConfig(opts ...Option) (*requestConfig, error) {
	cfg := &requestConfig{
		QueryParams: &url.Values{},
	}

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

func WithIds(ids []int) Option {
	return func(cfg *requestConfig) error {
		for _, id := range ids {
			cfg.QueryParams.Add("ids", strconv.Itoa(id))
		}
		return nil
	}
}

func WithUpdatedAfter(t time.Time) Option {
	return func(cfg *requestConfig) error {
		cfg.QueryParams.Set("updated_after", t.Format(time.RFC3339))
		return nil
	}
}
