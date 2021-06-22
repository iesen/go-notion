package notion

const (
	DefaultBaseUrl = "https://api.notion.com"
)

type ClientConfig struct {
	BaseUrl string
}

type Client struct {
	config ClientConfig
}

func NewClient() Client {
	return Client{
		config: ClientConfig{BaseUrl: DefaultBaseUrl},
	}
}

func NewClientWithConfig(config ClientConfig) Client {
	return Client{
		config: config,
	}
}
