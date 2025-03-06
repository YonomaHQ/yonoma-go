package yonoma

import "fmt"

type Tag struct {
	Name string `json:"name"`
}

func (c *Client) ListTags() ([]byte, error) {
	return c.request("GET", "/tags/list", nil)
}

func (c *Client) CreateTag(tagData map[string]interface{}) ([]byte, error) {
	return c.request("POST", "/tags/create", tagData)
}

func (c *Client) UpdateTag(tagID string, tagData map[string]interface{}) ([]byte, error) {
	endpoint := fmt.Sprintf("tags/%s/update", tagID)
	return c.request("POST", endpoint, tagData)
}

func (c *Client) RetrieveTag(tagID string) ([]byte, error) {
	endpoint := fmt.Sprintf("tags/%s", tagID)
	return c.request("GET", endpoint, nil)
}

func (c *Client) DeleteTag(tagID string) ([]byte, error) {
	endpoint := fmt.Sprintf("tags/%s/delete", tagID)
	return c.request("POST", endpoint, nil)
}
