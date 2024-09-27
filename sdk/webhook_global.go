package sdk

func GetWebhooks() ([]*Webhook, error) {
	return globalClient.GetWebhooks()
}

func GetPaginationWebhooks(p int, pageSize int, queryMap map[string]string) ([]*Webhook, int, error) {
	return globalClient.GetPaginationWebhooks(p, pageSize, queryMap)
}

func GetWebhook(name string) (*Webhook, error) {
	return globalClient.GetWebhook(name)
}

func UpdateWebhook(webhook *Webhook) (bool, error) {
	return globalClient.UpdateWebhook(webhook)
}

func AddWebhook(webhook *Webhook) (bool, error) {
	return globalClient.AddWebhook(webhook)
}

func DeleteWebhook(webhook *Webhook) (bool, error) {
	return globalClient.DeleteWebhook(webhook)
}
