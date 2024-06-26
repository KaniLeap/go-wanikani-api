package wk

import "time"

type User struct {
	UserId                   string    `json:"id"`
	Name                     string    `json:"username"`
	Level                    int       `json:"level"`
	ProfileURL               string    `json:"profile_url"`
	StartedAt                time.Time `json:"started_at"`
	CurrentVacationStartedAt time.Time `json:"current_vacation_started_at"`
	Subscription             struct {
		Active          bool      `json:"active"`
		Type            string    `json:"type"`
		MaxLevelGranted int       `json:"max_level_granted"`
		PeriodEndsAt    time.Time `json:"period_ends_at"`
	} `json:"subscription"`
	Preferences Preferences `json:"preferences"`
}

func (c *Client) GetUser(opts ...Option) (*Resource[User], error) {
	return resource[User](c, "user", "GET", nil, opts...)
}

func (c *Client) UpdateUser(payload Users, opts ...Option) (*Resource[User], error) {
	return resource[User](c, "user", "PUT", payload, opts...)
}
