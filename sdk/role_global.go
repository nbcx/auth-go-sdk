package sdk

func GetRoles() ([]*Role, error) {
	return globalClient.GetRoles()
}

func GetPaginationRoles(p int, pageSize int, queryMap map[string]string) ([]*Role, int, error) {
	return globalClient.GetPaginationRoles(p, pageSize, queryMap)
}

func GetRole(name string) (*Role, error) {
	return globalClient.GetRole(name)
}

func UpdateRole(role *Role) (bool, error) {
	return globalClient.UpdateRole(role)
}

func UpdateRoleForColumns(role *Role, columns []string) (bool, error) {
	return globalClient.UpdateRoleForColumns(role, columns)
}

func AddRole(role *Role) (bool, error) {
	return globalClient.AddRole(role)
}

func DeleteRole(role *Role) (bool, error) {
	return globalClient.DeleteRole(role)
}
