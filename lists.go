package yonoma

type List struct {
	Name string `json:"name"`
}

func (c *Client) ListLists() ([]byte, error) {
	return c.request("GET", "lists/list", nil)
}

func (c *Client) CreateList(listData map[string]interface{}) ([]byte, error) {
	return c.request("POST", "lists/create", listData)
}

func (c *Client) UpdateList(listID string, listData map[string]interface{}) ([]byte, error) {
	return c.request("POST", "lists/"+listID+"/update", listData)
}

func (c *Client) RetrieveList(listID string) ([]byte, error) {
	return c.request("GET", "lists/"+listID, nil)
}

func (c *Client) DeleteList(listID string) ([]byte, error) {
	return c.request("POST", "lists/"+listID+"/delete", nil)
}
