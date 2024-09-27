package sdk

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("adapter")

	// Add a new object
	adapter := &Adapter{
		Owner:       "admin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		User:        name,
		Host:        "https://casdoor.org",
	}
	_, err := AddAdapter(adapter)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	adapters, err := GetAdapters()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range adapters {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	adapter, err = GetAdapter(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if adapter.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", adapter.Name, name)
	}

	// Update the object
	updatedUser := "Updated Casdoor Website"
	adapter.User = updatedUser
	_, err = UpdateAdapter(adapter)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedadapter, err := GetAdapter(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedadapter.User != updatedUser {
		t.Fatalf("Failed to update object, User mismatch: %s != %s", updatedadapter.User, updatedUser)
	}

	// Delete the object
	_, err = DeleteAdapter(adapter)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedadapter, err := GetAdapter(name)
	if err != nil || deletedadapter != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
