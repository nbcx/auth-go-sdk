package sdk

import (
	"testing"
)

func TestSubscription(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Subscription")

	// Add a new object
	subscription := &Subscription{
		Owner:       "admin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		DisplayName: name,
		Description: "Casdoor Website",
	}
	_, err := AddSubscription(subscription)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	subscriptions, err := GetSubscriptions()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range subscriptions {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	subscription, err = GetSubscription(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if subscription.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", subscription.Name, name)
	}

	// Update the object
	updatedDescription := "Updated Casdoor Website"
	subscription.Description = updatedDescription
	_, err = UpdateSubscription(subscription)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedSubscription, err := GetSubscription(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedSubscription.Description != updatedDescription {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedSubscription.Description, updatedDescription)
	}

	// Delete the object
	_, err = DeleteSubscription(subscription)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedSubscription, err := GetSubscription(name)
	if err != nil || deletedSubscription != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
