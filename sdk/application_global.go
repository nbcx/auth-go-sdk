package sdk

func GetApplications() ([]*Application, error) {
	return globalClient.GetApplications()
}

func GetOrganizationApplications() ([]*Application, error) {
	return globalClient.GetOrganizationApplications()
}

func GetApplication(name string) (*Application, error) {
	return globalClient.GetApplication(name)
}

func AddApplication(application *Application) (bool, error) {
	return globalClient.AddApplication(application)
}

func DeleteApplication(application *Application) (bool, error) {
	return globalClient.DeleteApplication(application)
}

func UpdateApplication(application *Application) (bool, error) {
	return globalClient.UpdateApplication(application)
}
