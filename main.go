package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"qjey/config"
	"qjey/storage"
	"qjey/utils"
	"runtime"
)

func main() {
	if runtime.GOOS != "windows" {
		os.Clearenv()
	}

	// initial basic things
	utils.InitLog()
	config.Init()

	// starting the engine
	logrus.Info("starting qjey engine")
	storage.Init()
}