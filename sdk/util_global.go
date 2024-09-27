package sdk

import "io"

func GetUrl(action string, queryMap map[string]string) string {
	return globalClient.GetUrl(action, queryMap)
}

// DoGetResponse is a general function to get response from param url through HTTP Get method.
func DoGetResponse(url string) (*Response, error) {
	return globalClient.DoGetResponse(url)
}

// DoGetBytes is a general function to get response data in bytes from param url through HTTP Get method.
func DoGetBytes(url string) ([]byte, error) {
	return globalClient.DoGetBytes(url)
}

// DoGetBytesRaw is a general function to get response from param url through HTTP Get method.
func DoGetBytesRaw(url string) ([]byte, error) {
	return globalClient.DoGetBytesRaw(url)
}

func DoPost(action string, queryMap map[string]string, postBytes []byte, isForm, isFile bool) (*Response, error) {
	return globalClient.DoPost(action, queryMap, postBytes, isForm, isFile)
}

// DoPostBytesRaw is a general function to post a request from url, body through HTTP Post method.
func DoPostBytesRaw(url string, contentType string, body io.Reader) ([]byte, error) {
	return globalClient.DoPostBytesRaw(url, contentType, body)
}
