package sdk

import (
	"testing"
)

func TestModel(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Model")

	// Add a new object
	model := &Model{
		Owner:       "casbin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		DisplayName: name,
		ModelText: `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act`,
	}
	_, err := AddModel(model)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	models, err := GetModels()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range models {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	model, err = GetModel(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if model.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", model.Name, name)
	}

	// Update the object
	updatedDisplayName := "UpdatedName"
	model.DisplayName = updatedDisplayName
	_, err = UpdateModel(model)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedModel, err := GetModel(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedModel.DisplayName != updatedDisplayName {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedModel.DisplayName, updatedDisplayName)
	}

	// Delete the object
	_, err = DeleteModel(model)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedModel, err := GetModel(name)
	if err != nil || deletedModel != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
