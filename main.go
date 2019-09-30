package main

import (
	"psmockserver/pkg/config"
	"psmockserver/pkg/logger"
	"psmockserver/pkg/server"

	"github.com/kataras/golog"
)

func init() {
	logger.Setup()
}
func main() {
	golog.Infof("Server started on http://localhost:%s", config.Cfg.Server.Port)
	server.Start(config.Cfg.Server.Port)
}
