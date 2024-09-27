package sdk

import (
	"testing"
)

func TestGroup(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Group")

	// Add a new object
	group := &Group{
		Owner:       "admin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		DisplayName: name,
	}
	_, err := AddGroup(group)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	groups, err := GetGroups()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range groups {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	group, err = GetGroup(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if group.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", group.Name, name)
	}

	// Update the object
	updatedDisplayName := "Updated Casdoor Website"
	group.DisplayName = updatedDisplayName
	_, err = UpdateGroup(group)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedGroup, err := GetGroup(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedGroup.DisplayName != updatedDisplayName {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedGroup.DisplayName, updatedDisplayName)
	}

	// Delete the object
	_, err = DeleteGroup(group)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedGroup, err := GetGroup(name)
	if err != nil || deletedGroup != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
