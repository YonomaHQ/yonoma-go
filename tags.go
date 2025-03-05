package yonoma

type Tag struct {
	Name string `json:"name"`
}

func (c *Client) ListTags() ([]byte, error) {
	return c.request("GET", "/tags/list", nil)
}

func (c *Client) CreateTag(tag Tag) ([]byte, error) {
	return c.request("POST", "/tags", tag)
}

func (c *Client) UpdateTag(id string, tag Tag) ([]byte, error) {
	return c.request("PUT", "/tags/"+id, tag)
}

func (c *Client) RetrieveTag(id string) ([]byte, error) {
	return c.request("GET", "/tags/"+id, nil)
}

func (c *Client) DeleteTag(id string) ([]byte, error) {
	return c.request("DELETE", "/tags/"+id, nil)
}
