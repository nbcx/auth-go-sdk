package sdk

func SendEmail(title string, content string, sender string, receivers ...string) error {
	return globalClient.SendEmail(title, content, sender, receivers...)
}

func SendEmailByProvider(title string, content string, sender string, provider string, receivers ...string) error {
	return globalClient.SendEmailByProvider(title, content, sender, provider, receivers...)
}
