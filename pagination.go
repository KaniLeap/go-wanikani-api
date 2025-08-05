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

func (p *Paginate[T]) fromURL(url string) error {
	if url == "" {
		return errors.New("url is empty")
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

func (p *Paginate[T]) Next() error {
	return p.fromURL(p.Data.Pages.NextURL)
}

func (p *Paginate[T]) Previous() error {
	return p.fromURL(p.Data.Pages.PreviousURL)
}
