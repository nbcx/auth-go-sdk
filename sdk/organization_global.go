package sdk

func GetOrganization(name string) (*Organization, error) {
	return globalClient.GetOrganization(name)
}

func GetOrganizations() ([]*Organization, error) {
	return globalClient.GetOrganizations()
}

func GetOrganizationNames() ([]*Organization, error) {
	return globalClient.GetOrganizationNames()
}

func AddOrganization(organization *Organization) (bool, error) {
	return globalClient.AddOrganization(organization)
}

func DeleteOrganization(organization *Organization) (bool, error) {
	return globalClient.DeleteOrganization(organization)
}

func UpdateOrganization(organization *Organization) (bool, error) {
	return globalClient.UpdateOrganization(organization)
}
