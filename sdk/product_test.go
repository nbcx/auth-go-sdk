package sdk

import (
	"testing"
)

func TestProduct(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Product")

	// Add a new object
	product := &Product{
		Owner:       "admin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		DisplayName: name,

		Image:       "https://cdn.casbin.org/img/casdoor-logo_1185x256.png",
		Description: "Casdoor Website",
		Tag:         "auto_created_product_for_plan",

		Quantity: 999,
		Sold:     0,
		State:    "Published",
	}
	_, err := AddProduct(product)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	products, err := GetProducts()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range products {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	product, err = GetProduct(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if product.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", product.Name, name)
	}

	// Update the object
	updatedDescription := "Updated Casdoor Website"
	product.Description = updatedDescription
	_, err = UpdateProduct(product)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedProduct, err := GetProduct(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedProduct.Description != updatedDescription {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedProduct.Description, updatedDescription)
	}

	// Delete the object
	_, err = DeleteProduct(product)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedProduct, err := GetProduct(name)
	if err != nil || deletedProduct != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
