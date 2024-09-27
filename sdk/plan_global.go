package sdk

func GetPlans() ([]*Plan, error) {
	return globalClient.GetPlans()
}

func GetPaginationPlans(p int, pageSize int, queryMap map[string]string) ([]*Plan, int, error) {
	return globalClient.GetPaginationPlans(p, pageSize, queryMap)
}

func GetPlan(name string) (*Plan, error) {
	return globalClient.GetPlan(name)
}

func UpdatePlan(plan *Plan) (bool, error) {
	return globalClient.UpdatePlan(plan)
}

func AddPlan(plan *Plan) (bool, error) {
	return globalClient.AddPlan(plan)
}

func DeletePlan(plan *Plan) (bool, error) {
	return globalClient.DeletePlan(plan)
}
