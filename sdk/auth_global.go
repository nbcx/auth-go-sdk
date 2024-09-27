package sdk

import "golang.org/x/oauth2"

func GetOAuthToken(code string, state string) (*oauth2.Token, error) {
	return globalClient.GetOAuthToken(code, state)
}

func RefreshOAuthToken(refreshToken string) (*oauth2.Token, error) {
	return globalClient.RefreshOAuthToken(refreshToken)
}
