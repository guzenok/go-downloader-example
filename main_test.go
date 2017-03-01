// Тест
package main

import (
	"os"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

var TEST_FILE_NAME = "testdata/urls.txt"

func TestWriteDB(t *testing.T) {
	// Очистка БД
	os.RemoveAll(DB_PATH)

	// Проверка кода завершения
	exitCode := 0
	doAll(&TEST_FILE_NAME, &exitCode)

	if exitCode != 0 {
		t.Errorf("Exit code is %d, wait 0\n", exitCode)
		t.FailNow()
		return
	}

	// Проверка наличия данных в БД
	db, err := leveldb.OpenFile(DB_PATH, nil)
	if err != nil {
		t.Fatalf("Error while open DB: %s\n", err.Error())
	}
	defer db.Close()

	// И что начинаются на известные префиксы
	iter := db.NewIterator(nil, nil)
	n := 0
	defer iter.Release()
	for iter.Next() {
		n += 1
		key := iter.Key()
		value := iter.Value()
		if string(key[0:22]) != "http://localhost:6060/" || string(value[0:15]) != "<!DOCTYPE html>" {
			t.Errorf("Error in saved data: '%s' => '%s'\n", key, value[0:15])
		}
	}
	if n < 150 {
		t.Errorf("Only %d rows in DB, wait > 150", n)
	}

	err = iter.Error()
	if err != nil {
		t.Fatalf("Error while iter DB: %s\n", err.Error())
	}

}
