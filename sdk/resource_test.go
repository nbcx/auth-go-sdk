package sdk

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func (resource *Resource) GetId() string {
	return fmt.Sprintf("%s/%s", resource.Owner, resource.Name)
}

func TestResource(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	filename := "resource.go"
	file, err := os.Open(filename)

	if err != nil {
		t.Fatalf("Failed to open the file: %v\n", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read data from the file: %v\n", err)
	}

	name := fmt.Sprintf("/casdoor/%s", filename)
	// Add a new object
	resource := &Resource{
		Owner:       "casbin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		Description: "Casdoor Website",
		User:        "casbin",
		FileName:    filename,
		FileSize:    len(data),
		Tag:         name,
	}
	_, _, err = UploadResource(resource.User, resource.Tag, "", resource.FileName, data)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	Resources, err := GetResources(resource.Owner, resource.User, "", "", "", "")
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range Resources {
		if item.Tag == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	resource, err = GetResource(resource.GetId())
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if resource.Tag != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", resource.Name, name)
	}

	// Delete the object
	_, err = DeleteResource(resource)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedResource, err := GetResource(name)
	if err != nil || deletedResource != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
