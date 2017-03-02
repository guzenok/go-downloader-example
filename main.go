// downloader
package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	DB_PATH             = "./db"
	MAX_CONECTION_COUNT = 100
)

var fileName = flag.String("urls", "", "name of file with URLs")

func main() {

	// код завершения
	var exitCode = 0
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Fprint(os.Stderr, msg)
		}
		os.Exit(exitCode)
	}()

	// Запуск прогрессбара
	InitBar()

	// разбор параметров командной строки
	flag.Parse()
	if *fileName == "" {
		exitCode = 2
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}
	// код возврата на случай, если что-то пойдет не так
	exitCode = 1
	// основная работа
	doAll(fileName)
	// успешное завершение
	exitCode = 0
	// прощальное сообщение
	FinishBar()
	fmt.Fprintf(os.Stdout, "Done. Result in %s\n", DB_PATH)
}

func doAll(fileName *string) {

	// Открываем входной файл
	scanner, err := OpenFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("Error while open file: %s\n", err.Error()))
	}
	defer CloseFile()

	// Открываем выходную БД
	err = OpenDB(DB_PATH)
	if err != nil {
		panic(fmt.Sprintf("Error while open DB: %s\n", err.Error()))
	}
	defer CloseDB()

	// Читаем строки из файла и запускаем обработку каждого
	for scanner.Scan() {
		url := scanner.Text()
		// TODO: валидация url
		QuiueURL(url)
		IncBarTotal()
	}

	// Ожидание завершения всех обработок
	WaitAll()

}
