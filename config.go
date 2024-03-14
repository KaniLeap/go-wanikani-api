package wk

import (
	"strconv"
	"time"
)

type options struct {
	params map[string][]string
}

type Option func(*options) error

func WithIds(ids []int) Option {
	return func(cfg *options) error {
		for _, id := range ids {
			cfg.params["ids"] = append(cfg.params["ids"], strconv.Itoa(id))
		}
		return nil
	}
}

func WithUpdatedAfter(t time.Time) Option {
	return func(cfg *options) error {
		cfg.params["updated_after"] = append(cfg.params["updated_after"], t.Format(time.RFC3339))
		return nil
	}
}
