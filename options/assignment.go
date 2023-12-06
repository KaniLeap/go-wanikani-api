package options

import (
	"strconv"
	"time"
)

func WithAvailableAfter(t time.Time) Option {
	return func(cfg *RequestConfig) error {
		cfg.QueryParams.Set("available_after", t.Format(time.RFC3339))
		return nil
	}
}

func WithAvailableBefore(t time.Time) Option {
	return func(cfg *RequestConfig) error {
		cfg.QueryParams.Set("available_before", t.Format(time.RFC3339))
		return nil
	}
}

func WithBurned() Option {
	return func(cfg *RequestConfig) error {
		cfg.QueryParams.Set("burned", "true")
		return nil
	}
}

func WithHidden() Option {
	return func(cfg *RequestConfig) error {
		cfg.QueryParams.Set("hidden", "true")
		return nil
	}
}

func WithAvailableLessons() Option {
	return func(cfg *RequestConfig) error {
		cfg.QueryParams.Set("available_lessions", "true")
		return nil
	}
}

func WithAvailableReviews() Option {
	return func(cfg *RequestConfig) error {
		cfg.QueryParams.Set("available_reviews", "true")
		return nil
	}
}

func WithLevels(levels []int) Option {
	return func(cfg *RequestConfig) error {
		for _, l := range levels {
			cfg.QueryParams.Add("levels", strconv.Itoa(l))
		}
		return nil
	}
}

func WithStage(stages []int) Option {
	return func(cfg *RequestConfig) error {
		for _, s := range stages {
			cfg.QueryParams.Add("srs_stages", strconv.Itoa(s))
		}
		return nil
	}
}

func WithSubjectIDs(ids []int) Option {
	return func(cfg *RequestConfig) error {
		for _, id := range ids {
			cfg.QueryParams.Add("subject_ids", strconv.Itoa(id))
		}
		return nil
	}
}

func WithSubjectTypes(types []string) Option {
	return func(cfg *RequestConfig) error {
		for _, t := range types {
			cfg.QueryParams.Add("subject_types", t)
		}
		return nil
	}
}

func WithUnlocked() Option {
	return func(cfg *RequestConfig) error {
		cfg.QueryParams.Set("unlocked", "true")
		return nil
	}
}
