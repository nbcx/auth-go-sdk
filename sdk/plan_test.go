package sdk

import (
	"testing"
)

func TestPlan(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Plan")

	// Add a new object
	plan := &Plan{
		Owner:       "admin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		DisplayName: name,
		Description: "Casdoor Website",
	}
	_, err := AddPlan(plan)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	Plans, err := GetPlans()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range Plans {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	plan, err = GetPlan(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if plan.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", plan.Name, name)
	}

	// Update the object
	updatedDescription := "Updated Casdoor Website"
	plan.Description = updatedDescription
	_, err = UpdatePlan(plan)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedPlan, err := GetPlan(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedPlan.Description != updatedDescription {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedPlan.Description, updatedDescription)
	}

	// Delete the object
	_, err = DeletePlan(plan)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedPlan, err := GetPlan(name)
	if err != nil || deletedPlan != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
