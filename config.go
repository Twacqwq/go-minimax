package minimax

import "net/http"

const (
	APIV1                    = "https://api.minimax.chat/v1"
	defaultEmptyMessageLimit = 300
)

type Config struct {
	apiToken string
	groupId  string

	BaseURL           string
	HTTPClient        *http.Client
	EmptyMessageLimit uint
}

func DefaultConfig(apiToken, groupId string) *Config {
	return &Config{
		apiToken:          apiToken,
		groupId:           groupId,
		BaseURL:           APIV1,
		HTTPClient:        &http.Client{},
		EmptyMessageLimit: defaultEmptyMessageLimit,
	}
}
