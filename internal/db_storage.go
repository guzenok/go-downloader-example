package internal

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func OpenDB(path string) (err error) {
	if db == nil {
		db, err = leveldb.OpenFile(path, nil)
	}
	return err
}

func CloseDB() {
	db.Close()
	db = nil
}

func Save(key string, val []byte) {
	db.Put([]byte(key), val, nil)
}
