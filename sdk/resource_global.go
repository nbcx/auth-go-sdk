package sdk

func GetResource(id string) (*Resource, error) {
	return globalClient.GetResource(id)
}

func GetResourceEx(owner, name string) (*Resource, error) {
	return globalClient.GetResourceEx(owner, name)
}

func GetResources(owner, user, field, value, sortField, sortOrder string) ([]*Resource, error) {
	return globalClient.GetResources(owner, user, field, value, sortField, sortOrder)
}

func GetPaginationResources(owner, user, field, value string, pageSize, page int, sortField, sortOrder string) ([]*Resource, error) {
	return globalClient.GetPaginationResources(owner, user, field, value, pageSize, page, sortField, sortOrder)
}

func UploadResource(user string, tag string, parent string, fullFilePath string, fileBytes []byte) (string, string, error) {
	return globalClient.UploadResource(user, tag, parent, fullFilePath, fileBytes)
}

func UploadResourceEx(user string, tag string, parent string, fullFilePath string, fileBytes []byte, createdTime string, description string) (string, string, error) {
	return globalClient.UploadResourceEx(user, tag, parent, fullFilePath, fileBytes, createdTime, description)
}

func DeleteResource(resource *Resource) (bool, error) {
	return globalClient.DeleteResource(resource)
}

func DeleteResourceByName(name string) (bool, error) {
	return globalClient.DeleteResourceByName(name)
}
