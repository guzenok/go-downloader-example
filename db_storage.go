// Работа с целевой БД
package main

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

// Открыть
func OpenDB(path string) (err error) {
	db, err = leveldb.OpenFile(path, nil)
	return err
}

// Закрыть
func CloseDB() {
	db.Close()
}

// Записать
func Save(key string, val []byte) {
	db.Put([]byte(key), val, nil)
}
