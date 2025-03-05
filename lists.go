package yonoma

type List struct {
	Name string `json:"name"`
}

func (c *Client) ListLists() ([]byte, error) {
	return c.request("GET", "/lists/list", nil)
}

func (c *Client) CreateList(list List) ([]byte, error) {
	return c.request("POST", "/lists", list)
}

func (c *Client) UpdateList(id string, list List) ([]byte, error) {
	return c.request("PUT", "/lists/"+id, list)
}

func (c *Client) RetrieveList(id string) ([]byte, error) {
	return c.request("GET", "/lists/"+id, nil)
}

func (c *Client) DeleteList(id string) ([]byte, error) {
	return c.request("DELETE", "/lists/"+id, nil)
}
