package qjeyserver

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type QjeyServer struct {
	Config QjeyServerConfig
}

// NewQjeyServer .. create a new qjay server instance
func NewQjeyServer(config QjeyServerConfig) *QjeyServer {
	qjeyServer := &QjeyServer{
		Config: config,
	}

	if os.Getenv("API_KEY") == "changemenow" { // warn about the default api key
		logrus.Warn("please change the default API_KEY env var because is not secure")
	}
	logrus.Info(fmt.Sprintf("starting the qjey server on %s:%d", qjeyServer.Config.Host, qjeyServer.Config.Port))
	qjeyServer.init()

	return qjeyServer
}

func (s *QjeyServer) init() {

}