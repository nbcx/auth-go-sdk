package sdk

import (
	"testing"
)

func TestRole(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Role")

	// Add a new object
	role := &Role{
		Owner:       "admin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		DisplayName: name,
		Description: "Casdoor Website",
	}
	_, err := AddRole(role)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	roles, err := GetRoles()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range roles {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	role, err = GetRole(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if role.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", role.Name, name)
	}

	// Update the object
	updatedDescription := "Updated Casdoor Website"
	role.Description = updatedDescription
	_, err = UpdateRole(role)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedRole, err := GetRole(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedRole.Description != updatedDescription {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedRole.Description, updatedDescription)
	}

	// Delete the object
	_, err = DeleteRole(role)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedRole, err := GetRole(name)
	if err != nil || deletedRole != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}