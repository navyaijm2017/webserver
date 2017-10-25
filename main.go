package main

import (
	"github.com/luopengift/golibs/logger"
	"github.com/luopengift/types"
	"github.com/luopengift/webserver/apps"
)

func main() {
	config := &apps.Config{}
	err := types.ParseConfigFile("config.ini", config)
	if err != nil {
		logger.Error("%v", err)
	}
	logger.Info("%#v", config)
	apps.Init(config)
	select {}
}
