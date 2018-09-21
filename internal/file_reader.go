package internal

import (
	"bufio"
	"os"
)

func OpenFile(fileName *string) (scanner *bufio.Scanner, close func(), err error) {
	var inputFile *os.File

	inputFile, err = os.Open(*fileName)
	if err == nil {
		scanner = bufio.NewScanner(inputFile)
	}

	close = func() {
		inputFile.Close()
	}

	return
}
