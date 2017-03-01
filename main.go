// downloader
package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	DB_FILE_NAME = "./db"
)

var fileName = flag.String("urls", "", "имя файла с URLs")

func printHelp() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Parse()
	if *fileName == "" {
		printHelp()
		os.Exit(2)
	} else {
		var urls, err = readFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while read file: %s\n", err.Error())
			os.Exit(1)
		}
		download(&urls)
	}
	os.Exit(0)
}
