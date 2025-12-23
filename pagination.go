package wk

import (
	"context"
	"errors"
	"github.com/carlmjohnson/requests"
)

type Paginate[T any] struct {
	Data    Collection[T]
	request *requests.Builder
	ctx     context.Context
}

var (
	ErrNoNextPage     = errors.New("no next page")
	ErrNoPreviousPage = errors.New("no previous page")
	ErrEmptyURL       = errors.New("url is empty")
)

func (p *Paginate[T]) fromURL(url string) error {
	if url == "" {
		return ErrEmptyURL
	}

	var data Collection[T]
	err := p.request.
		BaseURL(url).
		ToJSON(&data).
		Fetch(p.ctx)
	if err != nil {
		return err
	}

	p.Data = data

	return nil
}

func (p *Paginate[T]) HasNext() bool {
	return p.Data.Pages.NextURL != ""
}

func (p *Paginate[T]) HasPrevious() bool {
	return p.Data.Pages.PreviousURL != ""
}

func (p *Paginate[T]) Next() error {
	if !p.HasNext() {
		return ErrNoNextPage
	}
	return p.fromURL(p.Data.Pages.NextURL)
}

func (p *Paginate[T]) Previous() error {
	if !p.HasPrevious() {
		return ErrNoPreviousPage
	}
	return p.fromURL(p.Data.Pages.PreviousURL)
}
