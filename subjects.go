package wk

import (
	"strconv"
	"time"
)

type Subject struct {
	AuxiliaryMeanings []struct {
		Meaning        string `json:"meaning"`
		Primary        bool   `json:"primary"`
		AcceptedAnswer bool   `json:"accepted_answer"`
	}
	Characters      string    `json:"characters"`
	CreatedAt       time.Time `json:"created_at"`
	DocumentURL     string    `json:"document_url"`
	HiddenAt        time.Time `json:"hidden_at"`
	Level           int       `json:"level"`
	MeaningMnemonic string    `json:"meaning_mnemonic"`
	Meanings        []struct {
		Meaning string `json:"meaning"`
		Type    string `json:"type"`
	}
	Slug  string `json:"slug"`
	SRSId int    `json:"spaced_repetition_system_id"`
}

func WithTypes(types []string) Option {
	return func(cfg *options) error {
		cfg.params["types"] = types
		return nil
	}
}

func WithSlugs(slugs []string) Option {
	return func(cfg *options) error {
		cfg.params["slugs"] = slugs
		return nil
	}
}

func WithLevels(levels []int) Option {
	return func(cfg *options) error {
		strLevels := make([]string, len(levels))
		for i, l := range levels {
			strLevels[i] = strconv.Itoa(l)
		}
		cfg.params["levels"] = strLevels
		return nil
	}
}

func WithHidden(hidden bool) Option {
	return func(cfg *options) error {
		cfg.params["hidden"] = []string{strconv.FormatBool(hidden)}
		return nil
	}
}

func (c *Client) GetSubjects(opts ...Option) (*Paginate[Subject], error) {
	return paginate[Subject](c, "subjects", opts...)
}

func (c *Client) GetSubject(id int, opts ...Option) (*Resource[Subject], error) {
	return resource[Subject](c, "subjects/"+strconv.Itoa(id), "GET", nil, opts...)
}
