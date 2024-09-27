package sdk

import (
	"testing"
)

func TestSms(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	sms := &smsForm{
		Content:   "casdoor",
		Receivers: []string{"+8613854673829", "+441932567890"},
	}
	err := SendSms(sms.Content, sms.Receivers...)
	if err != nil {
		t.Fatalf("Failed to send sms: %v", err)
	}

}
