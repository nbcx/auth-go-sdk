package sdk

func GetTokens() ([]*Token, error) {
	return globalClient.GetTokens()
}

func GetPaginationTokens(p int, pageSize int, queryMap map[string]string) ([]*Token, int, error) {
	return globalClient.GetPaginationTokens(p, pageSize, queryMap)
}

func GetToken(name string) (*Token, error) {
	return globalClient.GetToken(name)
}

func UpdateToken(token *Token) (bool, error) {
	return globalClient.UpdateToken(token)
}

func UpdateTokenForColumns(token *Token, columns []string) (bool, error) {
	return globalClient.UpdateTokenForColumns(token, columns)
}

func AddToken(token *Token) (bool, error) {
	return globalClient.AddToken(token)
}

func DeleteToken(token *Token) (bool, error) {
	return globalClient.DeleteToken(token)
}
