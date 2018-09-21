package internal

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"golang.org/x/net/context/ctxhttp"
)

const (
	MAX_CONECTION_COUNT = 30
	DELAY_SECONDS       = 2 // TODO: remove, for demonstration purpose only
)

var processesLimit = make(chan int, MAX_CONECTION_COUNT)

func ProcessFile(ctx context.Context, fileName *string) chan Status {
	progress := NewProgress(2 * MAX_CONECTION_COUNT)
	go processFile(ctx, fileName, progress)
	return progress.status
}

func processFile(ctx context.Context, fileName *string, status *Progress) {
	defer status.close()

	file, closeFile, err := OpenFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("Error while open file: %s\n", err.Error()))
	}
	defer closeFile()

	var wg sync.WaitGroup

READING:
	for file.Scan() {
		select {
		case <-ctx.Done():
			status.interrupt()
			break READING
		default:
			status.incTotal()
			url := file.Text()
			queueURL(ctx, status, &wg, url)
		}
	}

	wg.Wait()
}

func queueURL(ctx context.Context, status *Progress, wg *sync.WaitGroup, url string) {
	// TODO: validate url
	select {
	case <-ctx.Done():
		return
	case processesLimit <- 1:
		subctx, _ := context.WithCancel(ctx)
		go processURL(subctx, status, wg, url)
	}
}

func processURL(ctx context.Context, status *Progress, wg *sync.WaitGroup, url string) {
	wg.Add(1)
	defer func() {
		wg.Done()
		<-processesLimit
		status.incDone()
	}()

	time.Sleep(DELAY_SECONDS * time.Second) // TODO: remove, for demonstration purpose only

	resp, err := ctxhttp.Get(ctx, nil, url)
	if err != nil {
		// TODO: process error
		//fmt.Printf(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO: process error
		//fmt.Printf(err.Error())
		return
	}
	Save(url, body)
}
