package wk

import (
	"strconv"
	"time"
)

type Assignment struct {
	Type          string    `json:"subject_type"`
	Id            int       `json:"subject_id"`
	Level         int       `json:"level"`
	Stage         int       `json:"srs_stage"`
	UnlockedAt    time.Time `json:"unlocked_at"`
	StartedAt     time.Time `json:"started_at"`
	PassedAt      time.Time `json:"passed_at"`
	BurnedAt      time.Time `json:"burned_at"`
	AvailableAt   time.Time `json:"available_at"`
	ResurrectedAt time.Time `json:"resurrected_at"`
}

func WithAvailableAfter(t time.Time) Option {
	return func(cfg *options) error {
		cfg.params["available_after"] = append(cfg.params["available_after"], t.Format(time.RFC3339))
		return nil
	}
}

func WithAvailableBefore(t time.Time) Option {
	return func(cfg *options) error {
		cfg.params["available_before"] = append(cfg.params["available_before"], t.Format(time.RFC3339))
		return nil
	}
}

func WithBurned() Option {
	return func(cfg *options) error {
		cfg.params["burned"] = append(cfg.params["burned"], "true")
		return nil
	}
}

func WithHidden() Option {
	return func(cfg *options) error {
		cfg.params["hidden"] = append(cfg.params["hidden"], "true")
		return nil
	}
}

func WithAvailableLessons() Option {
	return func(cfg *options) error {
		cfg.params["available_lessons"] = append(cfg.params["available_lessons"], "true")
		return nil
	}
}

func WithAvailableReviews() Option {
	return func(cfg *options) error {
		cfg.params["available_reviews"] = append(cfg.params["available_reviews"], "true")
		return nil
	}
}

func WithLevels(levels []int) Option {
	return func(cfg *options) error {
		for _, level := range levels {
			cfg.params["levels"] = append(cfg.params["levels"], strconv.Itoa(level))
		}
		return nil
	}
}

func WithStage(stages []int) Option {
	return func(cfg *options) error {
		for _, s := range stages {
			cfg.params["srs_stages"] = append(cfg.params["srs_stages"], strconv.Itoa(s))
		}
		return nil
	}
}

func WithSubjectIDs(ids []int) Option {
	return func(cfg *options) error {
		for _, id := range ids {
			cfg.params["subject_ids"] = append(cfg.params["subject_ids"], strconv.Itoa(id))
		}
		return nil
	}
}

func WithSubjectTypes(types []string) Option {
	return func(cfg *options) error {
		for _, t := range types {
			cfg.params["subject_types"] = append(cfg.params["subject_types"], t)
		}
		return nil
	}
}

func WithUnlocked() Option {
	return func(cfg *options) error {
		cfg.params["unlocked"] = append(cfg.params["unlocked"], "true")
		return nil
	}
}

func (c *Client) GetAssignments(opts ...Option) (*Paginate[Assignment], error) {
	return paginate[Assignment](c, "assignments", opts...)
}

func (c *Client) GetAssignment(id int, opts ...Option) (*Resource[Assignment], error) {
	return resource[Assignment](c, "assignments/"+strconv.Itoa(id), "GET", nil, opts...)
}

func (c *Client) StartAssignment(payload Assignments, id int, opts ...Option) (*Resource[Assignment], error) {
	return resource[Assignment](c, "assignments/"+strconv.Itoa(id)+"/start", "PUT", payload, opts...)
}
