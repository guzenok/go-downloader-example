package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/guzenok/go_downloader/internal"
)

var fileName = flag.String("urls", "urls.txt", "name of file with URLs")
var dbDir = flag.String("datadir", "./db", "DB directory")

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
	internal.InitBar()

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
	internal.ProcessFile(fileName)
	// успешное завершение
	exitCode = 0
	// прощальное сообщение
	internal.FinishBar()
	fmt.Fprintf(os.Stdout, "Done. Result in %s\n", *dbDir)
}
