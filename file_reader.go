// Работа с входным файлом
package main

import (
	"bufio"
	"os"
)

var inputFile *os.File

func OpenFile(fileName *string) (scanner *bufio.Scanner, err error) {
	inputFile, err = os.Open(*fileName)
	if err == nil {
		scanner = bufio.NewScanner(inputFile)
	}
	return scanner, err
}

func CloseFile() {
	inputFile.Close()
}
