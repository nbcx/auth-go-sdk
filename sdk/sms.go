package sdk

import "encoding/json"

type smsForm struct {
	Content   string   `json:"content"`
	Receivers []string `json:"receivers"`
}

func (c *Client) SendSms(content string, receivers ...string) error {
	form := smsForm{
		Content:   content,
		Receivers: receivers,
	}
	postBytes, err := json.Marshal(form)
	if err != nil {
		return err
	}

	_, err = c.DoPost("send-sms", nil, postBytes, false, false)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SendSmsByProvider(content string, provider string, receivers ...string) error {
	form := smsForm{
		Content:   content,
		Receivers: receivers,
	}
	postBytes, err := json.Marshal(form)
	if err != nil {
		return err
	}

	providerMap := map[string]string{
		"provider": provider,
	}
	_, err = c.DoPost("send-sms", providerMap, postBytes, false, false)
	if err != nil {
		return err
	}
	return nil
}
