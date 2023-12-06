package wk

import (
	"strconv"

	"github.com/KaniLeap/go-wanikani-api/data"

	"github.com/KaniLeap/go-wanikani-api/options"
)

func (c *Client) GetAssignments(opts ...options.Option) (*pagination[data.Assignment], error) {
	return paginate[data.Assignment](c, "assignments", opts...)
}

func (c *Client) GetAssignment(id int, opts ...options.Option) (*data.Collection[data.Assignment], error) {
	return request[data.Assignment](c, "assignments/"+strconv.Itoa(id), "POST", opts...)
}

func (c *Client) StartAssignment(id int, opts ...options.Option) (*data.Collection[data.Assignment], error) {
	return request[data.Assignment](c, "assignments/"+strconv.Itoa(id)+"/start", "PUT", opts...)
}
