// downloader
package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	DB_PATH             = "./db"
	MAX_CONECTION_COUNT = 10
)

var fileName = flag.String("urls", "", "name of file with URLs")

func main() {

	// код завершения
	var exitCode = 0
	defer os.Exit(exitCode)

	// разбор параметров командной строки
	flag.Parse()
	if *fileName == "" {
		exitCode = 2
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}
	// основная работа
	doAll(fileName, &exitCode)

}

func doAll(fileName *string, exitCode *int) {

	*exitCode = 1
	// Открываем входной файл
	scanner, err := OpenFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while open file: %s\n", err.Error())
		return
	}
	defer CloseFile()

	// Открываем выходную БД
	err = OpenDB(DB_PATH)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while open DB: %s\n", err.Error())
		return
	}
	defer CloseDB()

	// Читаем строки из файла и запускаем обработку каждого
	for scanner.Scan() {
		url := scanner.Text()
		// TODO: валидация url
		fmt.Fprintf(os.Stdout, " - begin download %s\n", url)
		QuiueURL(url)
	}

	// Ожидание завершения всех обработок
	WaitAll()

	// Успешное завершение
	*exitCode = 0
	fmt.Fprintf(os.Stdout, "Done. Result in %s\n", DB_PATH)
}
