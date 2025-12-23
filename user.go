package wk

import (
	"context"
	"time"
)

type User struct {
	Id                       string     `json:"id"`
	Name                     string     `json:"username"`
	Level                    int        `json:"level"`
	ProfileURL               string     `json:"profile_url"`
	StartedAt                time.Time  `json:"started_at"`
	CurrentVacationStartedAt *time.Time `json:"current_vacation_started_at"`
	Subscription             struct {
		Active          bool      `json:"active"`
		Type            string    `json:"type"`
		MaxLevelGranted int       `json:"max_level_granted"`
		PeriodEndsAt    time.Time `json:"period_ends_at"`
	} `json:"subscription"`
	Preferences Preferences `json:"preferences"`
}

func (c *Client) GetUser(ctx context.Context, opts ...Option) (*Resource[User], error) {
	return resource[User](c, ctx, "user", "GET", nil, opts...)
}

func (c *Client) UpdateUser(ctx context.Context, payload Users, opts ...Option) (*Resource[User], error) {
	return resource[User](c, ctx, "user", "PUT", payload, opts...)
}
