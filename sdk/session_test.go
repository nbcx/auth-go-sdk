package sdk

import (
	"testing"
)

func TestSession(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Session")

	// Add a new object
	Session := &Session{
		Owner:       "casbin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		Application: "app-built-in",
		SessionId:   []string{},
	}
	_, err := AddSession(Session)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	Sessions, err := GetSessions()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range Sessions {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	Session, err = GetSession(name, Session.Application)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if Session.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", Session.Name, name)
	}

	// Update the object
	UpdateTime := GetCurrentTime()
	Session.CreatedTime = UpdateTime
	_, err = UpdateSession(Session)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedSession, err := GetSession(name, Session.Application)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedSession.CreatedTime != UpdateTime {
		t.Fatalf("Failed to update object, Application mismatch: %s != %s", updatedSession.CreatedTime, UpdateTime)
	}

	// Delete the object
	_, err = DeleteSession(Session)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedSession, err := GetSession(name, Session.Application)
	if err != nil || deletedSession != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
