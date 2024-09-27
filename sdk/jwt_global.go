package sdk

func ParseJwtToken(token string) (*Claims, error) {
	return globalClient.ParseJwtToken(token)
}
