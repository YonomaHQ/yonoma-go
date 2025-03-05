package yonoma

type Contact struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

func (c *Client) CreateContact(contact Contact) ([]byte, error) {
	return c.request("POST", "/contacts", contact)
}

func (c *Client) UnsubscribeContact(email string) ([]byte, error) {
	payload := map[string]string{"email": email}
	return c.request("POST", "/contacts/unsubscribe", payload)
}

func (c *Client) LabelContactWithTag(email, tag string) ([]byte, error) {
	payload := map[string]string{"email": email, "tag": tag}
	return c.request("POST", "/contacts/tag", payload)
}

func (c *Client) RemoveContactTag(email, tag string) ([]byte, error) {
	payload := map[string]string{"email": email, "tag": tag}
	return c.request("POST", "/contacts/untag", payload)
}
