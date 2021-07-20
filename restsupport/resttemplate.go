package restsupport

import "net/http"

type RestTemplate struct{}

func (t RestTemplate) Get(url string) (*http.Response, error) {
	return http.Get(url)
}
