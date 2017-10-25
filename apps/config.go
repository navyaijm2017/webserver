package apps

type HttpConfig struct {
	Addr string `json:"addr"`
}

type Config struct {
	HttpConfig
	WeChatConfig
}
