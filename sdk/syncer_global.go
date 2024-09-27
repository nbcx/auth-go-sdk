package sdk

func GetSyncers() ([]*Syncer, error) {
	return globalClient.GetSyncers()
}

func GetPaginationSyncers(p int, pageSize int, queryMap map[string]string) ([]*Syncer, int, error) {
	return globalClient.GetPaginationSyncers(p, pageSize, queryMap)
}

func GetSyncer(name string) (*Syncer, error) {
	return globalClient.GetSyncer(name)
}

func UpdateSyncer(syncer *Syncer) (bool, error) {
	return globalClient.UpdateSyncer(syncer)
}

func AddSyncer(syncer *Syncer) (bool, error) {
	return globalClient.AddSyncer(syncer)
}

func DeleteSyncer(syncer *Syncer) (bool, error) {
	return globalClient.DeleteSyncer(syncer)
}
