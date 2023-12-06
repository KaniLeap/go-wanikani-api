package wk

type Client struct {
	Token   string
	BaseURL string
}

func NewClient(token string) *Client {
	return &Client{
		Token:   token,
		BaseURL: "https://api.wanikani.com/v2/",
	}
}
