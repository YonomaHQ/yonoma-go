package yonoma

import "fmt"

type Contact struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

func (c *Client) CreateContact(listID string, contactData map[string]interface{}) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/%s/create", listID)
	return c.request("POST", endpoint, contactData)
}

func (c *Client) UnsubscribeContact(listID string, contactId string, contactData map[string]interface{}) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/%s/status/%s", listID, contactId)
	return c.request("POST", endpoint, contactData)
}

func (c *Client) AddTag(contactId string, contactData map[string]interface{}) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/tags/%s/add", contactId)
	return c.request("POST", endpoint, contactData)
}

func (c *Client) RemoveContactTag(contactId string, contactData map[string]interface{}) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/tags/%s/remove", contactId)
	return c.request("POST", endpoint, contactData)
}
