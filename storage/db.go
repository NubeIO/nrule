package storage

import (
	"github.com/tidwall/buntdb"
	"log"
	"os"
	"path/filepath"
)

type db struct {
	DB *buntdb.DB
}

func New(dbFile string) IStorage {
	if dbFile == "" {
		dbFile = "data/data.db"
	}
	parentDir := filepath.Dir(dbFile)
	if parentDir != "" {
		err := os.MkdirAll(parentDir, 0755)
		if err != nil {
			panic("data directory creation issue")
		}
	}
	newDb, err := buntdb.Open(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	return &db{DB: newDb}
}

func (inst *db) Close() error {
	return inst.DB.Close()
}

func matchRuleUUID(uuid string) bool {
	if len(uuid) == 16 {
		if uuid[0:4] == "rql_" {
			return true
		}
	}
	return false
}

func matchVarUID(uuid string) bool {
	if len(uuid) == 16 {
		if uuid[0:4] == "var_" {
			return true
		}
	}
	return false
}
