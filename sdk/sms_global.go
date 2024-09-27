package sdk

func SendSms(content string, receivers ...string) error {
	return globalClient.SendSms(content, receivers...)
}

func SendSmsByProvider(content string, provider string, receivers ...string) error {
	return globalClient.SendSmsByProvider(content, provider, receivers...)
}
