package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
	"os"
)

// Init .. init the config from .env file
func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Reset the utils level from the config file
	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		log.SetLevel(log.DebugLevel)
		break

	case "info":
		log.SetLevel(log.InfoLevel)
		break

	case "warn":
		log.SetLevel(log.WarnLevel)
		break

	case "error":
		log.SetLevel(log.ErrorLevel)
		break

	default:
		log.Fatal("unknown log level, exit ..")
	}
	log.Info(fmt.Sprintf("log level set to %s", os.Getenv("LOG_LEVEL")))
}