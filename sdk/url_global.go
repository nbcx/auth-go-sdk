package sdk

func GetSignupUrl(enablePassword bool, redirectUri string) string {
	return globalClient.GetSignupUrl(enablePassword, redirectUri)
}

func GetSigninUrl(redirectUri string) string {
	return globalClient.GetSigninUrl(redirectUri)
}

func GetUserProfileUrl(userName string, accessToken string) string {
	return globalClient.GetUserProfileUrl(userName, accessToken)
}

func GetMyProfileUrl(accessToken string) string {
	return globalClient.GetMyProfileUrl(accessToken)
}
