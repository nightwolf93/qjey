package storage

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
)

const (
	storageBasePath string = "./data"
)

var (
	db leveldb.DB
)

// Init .. initialize the storage db
func Init() {
	if _, err := os.Stat(storageBasePath); os.IsNotExist(err) { // create the storage directory if not exist
		logrus.Info("creating the storate path because the path doesnt exit")
		err = os.Mkdir(storageBasePath, 0777)
		if err != nil {
			logrus.Fatal("can't create the storage path, exit ..")
		}
		logrus.Info("storage path created with success")
	}

	_, err := leveldb.OpenFile(fmt.Sprintf("%s/data.db", storageBasePath), nil)
	if err != nil {
		logrus.Fatal("can't open db file")
	}
	logrus.Info("db file opened with success")
}