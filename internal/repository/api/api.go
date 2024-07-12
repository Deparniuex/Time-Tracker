package api

type ApiClientConfig struct {
	APIURL string
}

type ApiClient struct {
	ApiClientConfig
}

func New(cfg ApiClientConfig) *ApiClient {
	return &ApiClient{
		ApiClientConfig: cfg,
	}
}
