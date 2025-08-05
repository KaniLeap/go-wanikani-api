package wk

import (
	"context"
	"strconv"
	"time"
)

type ReviewBase struct {
	AssignmentId      int       `json:"assignment_id"`
	IncorrectMeanings int       `json:"incorrect_meaning_answers"`
	IncorrectReadings int       `json:"incorrect_reading_answers"`
	CreatedAt         time.Time `json:"created_at"`
}

type Review struct {
	ReviewBase
	SubjectId   int `json:"subject_id"`
	StartingSRS int `json:"starting_srs_stage"`
	EndingSRS   int `json:"ending_srs_stage"`
	SRSId       int `json:"spaced_repetition_system_id"`
}

func (c *Client) GetReview(ctx context.Context, id int, opts ...Option) (*Resource[Review], error) {
	return resource[Review](c, ctx, "reviews/"+strconv.Itoa(id), "GET", nil, opts...)
}

func (c *Client) CreateReview(ctx context.Context, payload Reviews, opts ...Option) (*Resource[Review], error) {
	return resource[Review](c, ctx, "reviews/", "POST", payload, opts...)
}
