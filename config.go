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
		strId := make([]string, len(ids))
		for i, id := range ids {
			strId[i] = strconv.Itoa(id)
		}
		cfg.params["ids"] = strId
		return nil
	}
}

func WithUpdatedAfter(t time.Time) Option {
	return func(cfg *options) error {
		cfg.params["updated_after"] = []string{t.Format(time.RFC3339)}
		return nil
	}
}
