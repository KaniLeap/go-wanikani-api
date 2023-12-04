package options

import (
	"strconv"
	"time"
)

func WithAvailableAfter(t time.Time) Option {
	return func(cfg *requestConfig) error {
		cfg.QueryParams.Set("available_after", t.Format(time.RFC3339))
		return nil
	}
}

func WithAvailableBefore(t time.Time) Option {
	return func(cfg *requestConfig) error {
		cfg.QueryParams.Set("available_before", t.Format(time.RFC3339))
		return nil
	}
}

func WithBurned() Option {
	return func(cfg *requestConfig) error {
		cfg.QueryParams.Set("burned", "true")
		return nil
	}
}

func WithHidden() Option {
	return func(cfg *requestConfig) error {
		cfg.QueryParams.Set("hidden", "true")
		return nil
	}
}

func WithAvailableLessons() Option {
	return func(cfg *requestConfig) error {
		cfg.QueryParams.Set("available_lessions", "true")
		return nil
	}
}

func WithAvailableReviews() Option {
	return func(cfg *requestConfig) error {
		cfg.QueryParams.Set("available_reviews", "true")
		return nil
	}
}

func WithLevels(level []int) Option {
	return func(cfg *requestConfig) error {
		for _, l := range level {
			cfg.QueryParams.Add("levels", strconv.Itoa(l))
		}
		return nil
	}
}
