package internal

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func openDB(path string) (err error) {
	if db == nil {
		db, err = leveldb.OpenFile(path, nil)
	}
	return err
}

func closeDB() {
	db.Close()
}

func save(key string, val []byte) {
	db.Put([]byte(key), val, nil)
}
