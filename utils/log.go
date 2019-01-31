package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// InitLog .. initialize the logger
func InitLog() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
