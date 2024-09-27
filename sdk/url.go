package sdk

import (
	"fmt"
	"net/url"
	"strings"
)

func (c *Client) GetSignupUrl(enablePassword bool, redirectUri string) string {
	// redirectUri can be empty string if enablePassword == true (only password enabled signup page is required)
	if enablePassword {
		return fmt.Sprintf("%s/signup/%s", c.Endpoint, c.ApplicationName)
	} else {
		return strings.ReplaceAll(c.GetSigninUrl(redirectUri), "/login/oauth/authorize", "/signup/oauth/authorize")
	}
}

func (c *Client) GetSigninUrl(redirectUri string) string {
	// origin := "https://door.casbin.com"
	// redirectUri := fmt.Sprintf("%s/callback", origin)
	scope := "read"
	state := c.ApplicationName
	return fmt.Sprintf("%s/login/oauth/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=%s&state=%s",
		c.Endpoint, c.ClientId, url.QueryEscape(redirectUri), scope, state)
}

func (c *Client) GetUserProfileUrl(userName string, accessToken string) string {
	param := ""
	if accessToken != "" {
		param = fmt.Sprintf("?access_token=%s", accessToken)
	}
	return fmt.Sprintf("%s/users/%s/%s%s", c.Endpoint, c.OrganizationName, userName, param)
}

func (c *Client) GetMyProfileUrl(accessToken string) string {
	param := ""
	if accessToken != "" {
		param = fmt.Sprintf("?access_token=%s", accessToken)
	}
	return fmt.Sprintf("%s/account%s", c.Endpoint, param)
}
