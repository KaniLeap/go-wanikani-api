package data

import (
	"time"
)

type Shared struct {
	Object    string    `json:"object"`
	URL       string    `json:"url"`
	UpdatedAt time.Time `json:"data_updated_at"`
}

type Page struct {
	NextURL     string `json:"next_url"`
	PreviousURL string `json:"previous_url"`
	PerPage     int    `json:"per_page"`
}

type CollectionBase struct {
	Shared
	Pages Page `json:"pages"`
	Count int  `json:"total_count"`
}

type ResourceBase struct {
	Shared
	Id int `json:"id"`
}

type Resource[T any] struct {
	ResourceBase
	Data T `json:"data"`
}

type Collection[T any] struct {
	CollectionBase
	Data []struct {
		ResourceBase
		Data T `json:"data"`
	} `json:"data"`
}
