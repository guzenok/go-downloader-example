package main

import (
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

func download(urls *[]string) {
	var db, err = leveldb.OpenFile(DB_FILE_NAME, nil)
	if err != nil {

	}
	defer db.Close()

	var wg sync.WaitGroup
	for _, url := range *urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					db.Put([]byte(url), body, nil)
				}
			}
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
