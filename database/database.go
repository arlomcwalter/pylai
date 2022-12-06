package database

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"os"
	"path/filepath"
)

var DB *leveldb.DB

func initDb(path string) error {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func closeDb() error {
	return DB.Close()
}

func Init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error reading home directory.")
	}

	db, err := leveldb.OpenFile(filepath.Join(homeDir, ".pylai"), nil)
	if err != nil {
		log.Fatal("Error opening database.")
	}

	DB = db
}

func Shutdown() {
	err := DB.Close()
	if err != nil {
		log.Fatal("Error closing database.")
	}
}
