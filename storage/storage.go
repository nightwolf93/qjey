package storage

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Init .. initialize the storage db
func Init() {
	if _, err := os.Stat("./data"); os.IsNotExist(err) { // create the storage directory if not exist
		logrus.Info("creating the storate path because the path doesnt exit")
		err = os.Mkdir("./data", 0777)
		if err != nil {
			logrus.Fatal("can't create the storage path, exit ..")
		}
		logrus.Info("storage path created with success")
	}

}