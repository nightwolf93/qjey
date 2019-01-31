package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// InitLog .. initialize the logger
func InitLog() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
