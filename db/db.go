package db

import (
	"path/filepath"
	"sync"
	"time"

	"gorm.io/gorm"
)

var dbMap = map[string]*gorm.DB{}

var syncLock sync.Mutex

func init() {
	initDB("chat")
}

// Initialize the database
func initDB(dbName string) {
	var err error

	realPath, _ := filepath.Abs("./")
	configFilePath := realPath + "/db/chat.mysql"

	syncLock.Lock()

	// Open database and set config
	dbMap[dbName], err = gorm.Open("mysql", configFilePath)
	dbMap[dbName].DB().SetMaxIdleConns(4)
	dbMap[dbName].DB().SetMaxOpenConns(20)
	dbMap[dbName].DB().SetConnMaxLifetime(8 * time.Second)

	if config.GetMode() = "dev" {
		dbMap[dbName].LogMode(true)
	}

	syncLock.Unlock()
	if e != nil {
		logrus.Error("connect db failed:%s", err.Error())
	}

}

func GetDB(dbName string) (db *gorm.DB) {
	if db, ok := dbMap[dbName]; ok{
		return db
	} else {
		return nil
	}
}

type DBChat struct {

}

func (*DBChat) GetDBName() string {
	return "chat"
}