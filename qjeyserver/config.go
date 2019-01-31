package qjeyserver

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type QjeyServerConfig struct {
	Host string // host binding
	Port int // port binding
}

// NewQjeyServerConfigFromEnv .. create new server config with the env variables
func NewQjeyServerConfigFromEnv() QjeyServerConfig {
	if os.Getenv("SERVER_HOST") == "" { // check SERVER_HOST
		logrus.Fatal("SERVER_HOST env var invalid, exit ..")
	}
	if os.Getenv("SERVER_PORT") == "" {  // check SERVER_PORT
		logrus.Fatal("SERVER_PORT env var invalid, exit ..")
	}

	serverHost := os.Getenv("SERVER_HOST")
	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		logrus.Fatal("SERVER_PORT env var need to be a integer, exit ..")
	}

	c := QjeyServerConfig{
		Host: serverHost,
		Port: serverPort,
	}
	return c
}