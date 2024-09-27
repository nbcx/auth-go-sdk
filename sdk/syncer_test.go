package sdk

import (
	"testing"
)

func TestSyncer(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Syncer")

	// Add a new object
	syncer := &Syncer{
		Owner:        "admin",
		Name:         name,
		CreatedTime:  GetCurrentTime(),
		Organization: "casbin",
		Host:         "localhost",
		Port:         3306,
		User:         "root",
		Password:     "123",
		DatabaseType: "mysql",
		Database:     "syncer_db",
		Table:        "user-table",
		SyncInterval: 1,
	}
	_, err := AddSyncer(syncer)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	syncers, err := GetSyncers()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range syncers {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	syncer, err = GetSyncer(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if syncer.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", syncer.Name, name)
	}

	// Update the object
	updatedHost := "Updated Host"
	syncer.Host = updatedHost
	_, err = UpdateSyncer(syncer)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedSyncer, err := GetSyncer(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedSyncer.Host != updatedHost {
		t.Fatalf("Failed to update object, host mismatch: %s != %s", updatedSyncer.Host, updatedHost)
	}

	// Delete the object
	_, err = DeleteSyncer(syncer)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedSyncer, err := GetSyncer(name)
	if err != nil || deletedSyncer != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
