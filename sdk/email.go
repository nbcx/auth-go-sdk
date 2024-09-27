package sdk

import "encoding/json"

type emailForm struct {
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Sender    string   `json:"sender"`
	Receivers []string `json:"receivers"`
}

func (c *Client) SendEmail(title string, content string, sender string, receivers ...string) error {
	form := emailForm{
		Title:     title,
		Content:   content,
		Sender:    sender,
		Receivers: receivers,
	}
	postBytes, err := json.Marshal(form)
	if err != nil {
		return err
	}

	_, err = c.DoPost("send-email", nil, postBytes, false, false)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SendEmailByProvider(title string, content string, sender string, provider string, receivers ...string) error {
	form := emailForm{
		Title:     title,
		Content:   content,
		Sender:    sender,
		Receivers: receivers,
	}
	postBytes, err := json.Marshal(form)
	if err != nil {
		return err
	}

	providerMap := map[string]string{
		"provider": provider,
	}
	_, err = c.DoPost("send-email", providerMap, postBytes, false, false)
	if err != nil {
		return err
	}
	return nil
}
