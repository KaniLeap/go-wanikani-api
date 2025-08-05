package wk

import (
	"time"
)

type Shared struct {
	Object    string    `json:"object"`
	URL       string    `json:"url"`
	UpdatedAt time.Time `json:"data_updated_at"`
}

type CollectionBase struct {
	Shared
	Pages struct {
		NextURL     string `json:"next_url"`
		PreviousURL string `json:"previous_url"`
		PerPage     int    `json:"per_page"`
	} `json:"pages"`
	Count int `json:"total_count"`
}

type ResourceBase struct {
	Shared
	Id int `json:"id"`
}

type Resource[T any] struct {
	ResourceBase
	Data T `json:"data"`
}

type Item[T any] struct {
	ResourceBase
	Data T `json:"data"`
}

type Collection[T any] struct {
	CollectionBase
	Data []Item[T] `json:"data"`
}
