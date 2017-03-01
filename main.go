// downloader
package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	DB_FILE_NAME        = "./db"
	MAX_CONECTION_COUNT = 10
)

var fileName = flag.String("urls", "", "имя файла с URLs")

func printHelp() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {

	// Разбор параметров командной строки
	flag.Parse()
	if *fileName == "" {
		printHelp()
		os.Exit(2)
	} else {

		// Открываем входной файл
		scanner, err := OpenFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while open file: %s\n", err.Error())
			// os.Exit(1)
			return
		} else {
			defer CloseFile()
		}

		// Открываем выходную БД
		err = OpenDB(DB_FILE_NAME)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while open DB: %s\n", err.Error())
			// os.Exit(1)
			return
		} else {
			defer CloseDB()
		}

		// Читаем строки из файла и запускаем обработку каждого
		for scanner.Scan() {
			url := scanner.Text()
			// TODO: валидация url
			fmt.Fprintf(os.Stdout, " - begin download %s\n", url)
			QuiueURL(url)
		}

		// Ожидание завершения всех обработок
		WaitAll()
	}
}
