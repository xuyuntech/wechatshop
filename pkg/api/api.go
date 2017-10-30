package api

type api struct {
	listen string
}

type ApiConfig struct {
	Listen string
}

func NewApi(config *ApiConfig) (*api, error) {
	return &api{
		listen: config.Listen,
	}, nil
}

func (*api) Run() error {
	return nil
}