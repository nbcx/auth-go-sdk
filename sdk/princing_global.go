package sdk

func GetPricings() ([]*Pricing, error) {
	return globalClient.GetPricings()
}

func GetPaginationPricings(p int, pageSize int, queryMap map[string]string) ([]*Pricing, int, error) {
	return globalClient.GetPaginationPricings(p, pageSize, queryMap)
}

func GetPricing(name string) (*Pricing, error) {
	return globalClient.GetPricing(name)
}

func UpdatePricing(pricing *Pricing) (bool, error) {
	return globalClient.UpdatePricing(pricing)
}

func AddPricing(pricing *Pricing) (bool, error) {
	return globalClient.AddPricing(pricing)
}

func DeletePricing(pricing *Pricing) (bool, error) {
	return globalClient.DeletePricing(pricing)
}
