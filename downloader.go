// downloader
package main

import (
	"flag"
	"fmt"
	"os"
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
		fmt.Println("Hello World! " + *fileName)
	}
	os.Exit(0)
}
