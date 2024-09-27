package sdk

import (
	"testing"
)

func TestProvider(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Provider")

	// Add a new object
	provider := &Provider{
		Owner:       "admin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		DisplayName: name,
		Category:    "Captcha",
		Type:        "Default",
	}
	_, err := AddProvider(provider)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	providers, err := GetProviders()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range providers {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	provider, err = GetProvider(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if provider.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", provider.Name, name)
	}

	// Update the object
	updatedDisplayName := "Updated Casdoor Website"
	provider.DisplayName = updatedDisplayName
	_, err = UpdateProvider(provider)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedProvider, err := GetProvider(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedProvider.DisplayName != updatedDisplayName {
		t.Fatalf("Failed to update object, DisplayName mismatch: %s != %s", updatedProvider.DisplayName, updatedDisplayName)
	}

	// Delete the object
	_, err = DeleteProvider(provider)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedProvider, err := GetProvider(name)
	if err != nil || deletedProvider != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
