package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var procSemafor = make(chan int, MAX_CONECTION_COUNT)
var procWaitCounter sync.WaitGroup

const (
	MAX_CONECTION_COUNT = 100
	DB_PATH             = "./db"
)

func ProcessFile(fileName *string) {

	// Открываем входной файл
	scanner, err := OpenFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("Error while open file: %s\n", err.Error()))
	}
	defer CloseFile()

	// Открываем выходную БД
	err = openDB(DB_PATH)
	if err != nil {
		panic(fmt.Sprintf("Error while open DB: %s\n", err.Error()))
	}
	defer closeDB()

	// Читаем строки из файла и запускаем обработку каждого
	for scanner.Scan() {
		url := scanner.Text()
		// TODO: валидация url
		quiueURL(url)
		IncBarTotal()
	}

	// Ожидание завершения всех обработок
	waitAll()

}

// Ожидание свободного местечка в семафоре, и запуск скачивания "в фоне"
func quiueURL(url string) {
	procSemafor <- 1
	go processURL(url)
}

// Скачивание по http и сохранение в БД
func processURL(url string) {
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
			save(url, body)
		}
	}
}

// Ожидание завершения всех скачиваний
func waitAll() {
	procWaitCounter.Wait()
}
