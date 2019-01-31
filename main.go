package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"qjey/qjeyserver"
	"qjey/storage"
	"qjey/utils"
	"runtime"
)

var (
	qjayServer *qjeyserver.QjeyServer
)

func main() {
	if runtime.GOOS != "windows" {
		os.Clearenv()
	}

	// initial basic things
	utils.InitLog()
	utils.InitConfig()

	// starting the engine
	logrus.Info("starting qjey engine")
	storage.Init()

	qjeyServerConfig := qjeyserver.NewQjeyServerConfigFromEnv()
	qjeyServer := qjeyserver.NewQjeyServer(qjeyServerConfig)

	_ = qjeyServer
}