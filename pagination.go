package wk

import (
	"github.com/KaniLeap/go-wanikani-api/data"
)

type pagination[T any] struct {
	Data   data.Collection[T]
	Client *Client
}

func (p *pagination[T]) Next() (*pagination[T], error) {
	if p.Data.Pages.NextURL == "" {
		return nil, nil
	}

	return paginateFromURL[T](p.Client, p.Data.Pages.NextURL)
}

func (p *pagination[T]) Previous() (*pagination[T], error) {
	if p.Data.Pages.PreviousURL == "" {
		return nil, nil
	}

	return paginateFromURL[T](p.Client, p.Data.Pages.PreviousURL)
}
