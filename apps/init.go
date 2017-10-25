package apps

import (
	"github.com/luopengift/gohttp"
)

type BaseHttpHandler struct {
	gohttp.HttpHandler
}

func Init(cfg *Config) {
	WeChatInit(cfg)
	app := gohttp.Init()
	app.Route("^/wechat$", &WeChatHandler{})
	app.Route("^/wechat/sendwx$", &TextHandler{})
	go app.Run(cfg.HttpConfig.Addr)
}
