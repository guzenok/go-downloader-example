// Обработка url,
package main

import (
	"io/ioutil"
	"net/http"
	"sync"
)

var procSemafor = make(chan int, MAX_CONECTION_COUNT)
var procWaitCounter sync.WaitGroup

// Ожидание свободного местечка в семафоре, и запуск скачивания "в фоне"
func QuiueURL(url string) {
	procSemafor <- 1
	go ProcessURL(url)
}

// Скачивание по http и сохранение в БД
func ProcessURL(url string) {
	// по окончании процесса обязательно освободить семафор и уменьшить счетчик процессов
	defer func() {
		<-procSemafor
		procWaitCounter.Done()
		IncBarValue()
	}()
	// увеличить счетчик процессов
	procWaitCounter.Add(1)
	// http-запрос
	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()
		// тело ответа
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			// запись в БД
			Save(url, body)
		}
	}
}

// Ожидание завершения всех скачиваний
func WaitAll() {
	procWaitCounter.Wait()
}
