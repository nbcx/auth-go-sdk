package sdk

func GetRecords() ([]*Record, error) {
	return globalClient.GetRecords()
}

func GetPaginationRecords(p int, pageSize int, queryMap map[string]string) ([]*Record, int, error) {
	return globalClient.GetPaginationRecords(p, pageSize, queryMap)
}

func GetRecord(name string) (*Record, error) {
	return globalClient.GetRecord(name)
}

func AddRecord(record *Record) (bool, error) {
	return globalClient.AddRecord(record)
}
