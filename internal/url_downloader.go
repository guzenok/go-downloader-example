package internal

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"golang.org/x/net/context/ctxhttp"
)

type Progress struct {
	read  int
	saved int
}

var procSemafor = make(chan int, MAX_CONECTION_COUNT)
var procWaitCounter sync.WaitGroup

const (
	MAX_CONECTION_COUNT = 10
	DB_PATH             = "./db"
)

func ProcessFile(ctx context.Context, fileName *string) chan Progress {
	progress := make(chan Progress)
	go processFile(ctx, fileName, progress)
	return progress
}

func processFile(ctx context.Context, fileName *string, progress chan<- Progress) {
	defer close(progress)

	err := openDB(DB_PATH)
	if err != nil {
		panic(fmt.Sprintf("Error while open DB: %s\n", err.Error()))
	}
	defer closeDB()

	file, err := OpenFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("Error while open file: %s\n", err.Error()))
	}
	defer CloseFile()
READING:
	for file.Scan() {
		select {
		case <-ctx.Done():
			break READING
		default:
			IncBarTotal()
			url := file.Text()
			quiueURL(ctx, url)
		}
	}

	procWaitCounter.Wait()
}

func quiueURL(ctx context.Context, url string) {
	// TODO: валидация url
	select {
	case <-ctx.Done():
		return
	case procSemafor <- 1:
		subctx, _ := context.WithCancel(ctx)
		go processURL(subctx, url)
	}
}

func processURL(ctx context.Context, url string) {
	procWaitCounter.Add(1)
	defer func() {
		procWaitCounter.Done()
		<-procSemafor
		IncBarValue()
	}()
	time.Sleep(2 * time.Second)
	resp, err := ctxhttp.Get(ctx, nil, url)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// TODO: process error
			return
		}
		save(url, body)
	}
}
