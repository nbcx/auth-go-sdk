package sdk

func GetGlobalCerts() ([]*Cert, error) {
	return globalClient.GetGlobalCerts()
}

func GetCerts() ([]*Cert, error) {
	return globalClient.GetCerts()
}

func GetCert(name string) (*Cert, error) {
	return globalClient.GetCert(name)
}

func UpdateCert(cert *Cert) (bool, error) {
	return globalClient.UpdateCert(cert)
}

func AddCert(cert *Cert) (bool, error) {
	return globalClient.AddCert(cert)
}

func DeleteCert(cert *Cert) (bool, error) {
	return globalClient.DeleteCert(cert)
}
