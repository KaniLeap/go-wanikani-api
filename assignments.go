package wk

import (
	"github.com/KaniLeap/go-wanikani-api/data"

	"github.com/KaniLeap/go-wanikani-api/options"
)

func (c *Client) GetAssignments(opts ...options.Option) (*data.Collection[data.Assignment], error) {
	return requestWithOpts[data.Assignment](c, "assignments", opts...)
}
