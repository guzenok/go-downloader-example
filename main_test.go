package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/go-chi/chi"
	"github.com/guzenok/go_downloader/internal"
)

// TODO: nice tests

var (
	LOCAL_PORT = "3000"
	URLS_COUNT = 100
	FILE_NAME  = path.Join(os.TempDir(), "urls.txt")
	DB_DIR     = path.Join(os.TempDir(), "go_downloader/")
)

func TestGeneral(t *testing.T) {
	defer func() {
		if msg := recover(); msg != nil {
			t.Fatalf("General error: %s\n", msg)
		}
	}()

	os.RemoveAll(DB_DIR)
	err := internal.OpenDB(DB_DIR)
	if err != nil {
		t.Fatalf("Error while open DB: %s\n", err.Error())
	}
	defer internal.CloseDB()

	err = genURLs(FILE_NAME, URLS_COUNT)
	if err != nil {
		t.Fatalf("Error while open DB: %s\n", err.Error())
	}

	server := startHttpServer()
	defer server.Close()

	ctx := context.TODO()
	for _ = range internal.ProcessFile(ctx, &FILE_NAME) {
	}

	db := internal.GetDB()
	iter := db.NewIterator(nil, nil)
	defer iter.Release()
	n := 0
	for iter.Next() {
		n += 1
		key := iter.Key()
		value := iter.Value()
		if string(key[0:17]) != "http://localhost:" || string(value[0:15]) != "<!DOCTYPE html>" {
			t.Errorf("Error in saved data: '%s' => '%s'\n", key, value[0:15])
		}
	}
	if n != URLS_COUNT {
		t.Errorf("Only %d rows in DB, wait > 150", n)
	}

	err = iter.Error()
	if err != nil {
		t.Fatalf("Error while iter DB: %s\n", err.Error())
	}

}

func genURLs(fileName string, count int) (err error) {
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer func() {
		err = file.Close()
	}()

	for i := 0; i < count; i++ {
		_, err = file.WriteString(fmt.Sprintf("http://localhost:%s/?%d\n", LOCAL_PORT, i))
		if err != nil {
			return
		}
	}

	return
}

func startHttpServer() *http.Server {
	handler := chi.NewRouter()
	handler.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<!DOCTYPE html>\n<html></html>\n"))
	})
	server := &http.Server{Addr: ":" + LOCAL_PORT, Handler: handler}
	go func() {
		_ = server.ListenAndServe()
	}()
	return server
}
