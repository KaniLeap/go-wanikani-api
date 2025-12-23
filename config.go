package wk

import (
	"strconv"
	"time"
)

type options struct {
	params map[string][]string
}

type Option func(*options) error

func WithIDs(ids ...int) Option {
	return func(cfg *options) error {
		strIDs := make([]string, len(ids))
		for i, id := range ids {
			strIDs[i] = strconv.Itoa(id)
		}
		cfg.params["ids"] = strIDs
		return nil
	}
}

func WithUpdatedAfter(t time.Time) Option {
	return func(cfg *options) error {
		cfg.params["updated_after"] = []string{t.Format(time.RFC3339)}
		return nil
	}
}
