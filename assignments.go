package wk

import (
	"context"
	"strconv"
	"time"
)

type Assignment struct {
	SubjectType   string     `json:"subject_type"`
	SubjectId     int        `json:"subject_id"`
	Level         int        `json:"level"`
	Stage         int        `json:"srs_stage"`
	UnlockedAt    *time.Time `json:"unlocked_at"`
	StartedAt     *time.Time `json:"started_at"`
	PassedAt      *time.Time `json:"passed_at"`
	BurnedAt      *time.Time `json:"burned_at"`
	AvailableAt   *time.Time `json:"available_at"`
	ResurrectedAt *time.Time `json:"resurrected_at"`
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

func WithBurned(burned bool) Option {
	return func(cfg *options) error {
		cfg.params["burned"] = []string{strconv.FormatBool(burned)}
		return nil
	}
}

func WithAvailableLessons(available bool) Option {
	return func(cfg *options) error {
		cfg.params["available_lessons"] = []string{strconv.FormatBool(available)}
		return nil
	}
}

func WithAvailableReviews(available bool) Option {
	return func(cfg *options) error {
		cfg.params["available_reviews"] = []string{strconv.FormatBool(available)}
		return nil
	}
}

func WithStages(stages ...int) Option {
	return func(cfg *options) error {
		strStages := make([]string, len(stages))
		for i, s := range stages {
			strStages[i] = strconv.Itoa(s)
		}
		cfg.params["srs_stages"] = strStages
		return nil
	}
}

func WithSubjectIDs(ids []int) Option {
	return func(cfg *options) error {
		strId := make([]string, len(ids))
		for i, id := range ids {
			strId[i] = strconv.Itoa(id)
		}
		cfg.params["subject_ids"] = strId
		return nil
	}
}

func WithSubjectTypes(types []string) Option {
	return func(cfg *options) error {
		strTypes := make([]string, len(types))
		copy(strTypes, types)
		cfg.params["subject_types"] = strTypes
		return nil
	}
}

func WithUnlocked(unlocked bool) Option {
	return func(cfg *options) error {
		cfg.params["unlocked"] = []string{strconv.FormatBool(unlocked)}
		return nil
	}
}

func (c *Client) GetAssignments(ctx context.Context, opts ...Option) (*Paginate[Assignment], error) {
	return paginate[Assignment](c, ctx, "assignments", opts...)
}

func (c *Client) GetAssignment(ctx context.Context, id int, opts ...Option) (*Resource[Assignment], error) {
	return resource[Assignment](c, ctx, "assignments/"+strconv.Itoa(id), "GET", nil, opts...)
}

func (c *Client) StartAssignment(ctx context.Context, payload Assignments, id int, opts ...Option) (*Resource[Assignment], error) {
	return resource[Assignment](c, ctx, "assignments/"+strconv.Itoa(id)+"/start", "PUT", payload, opts...)
}
