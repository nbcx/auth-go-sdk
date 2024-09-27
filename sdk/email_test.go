package sdk

import (
	"strings"
	"testing"
)

func TestEmail(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	email := &emailForm{
		Title:     "casbin",
		Content:   "casdoor-go-sdk website test",
		Sender:    "admin",
		Receivers: []string{"TestSmtpServer"},
	}
	err := SendEmail(email.Title, email.Content, email.Sender, email.Receivers...)
	if err != nil {
		if !strings.Contains(err.Error(), "535 Error") {
			t.Fatalf("Failed to send Email: %v", err)
		}
	}
}
