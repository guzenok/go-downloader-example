package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/guzenok/go_downloader/internal"
)

var fileName = flag.String("urls", "", "name of file with URLs")
var dbDir = flag.String("datadir", "./db", "DB directory")

func main() {
	flag.Parse()

	if *fileName == "" {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			fmt.Fprint(os.Stderr, err)
			os.Exit(2)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())

	interrupts := make(chan os.Signal, 1)
	signal.Notify(interrupts, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT, syscall.SIGKILL)
	go func() {
		for _ = range interrupts {
			cancel()
		}
	}()

	internal.InitBar()

	for progress := range internal.ProcessFile(ctx, fileName) {
		// TODO: display progress
		_ = progress
	}

	internal.FinishBar()

	fmt.Printf("Done. Result in %s\n", *dbDir)
}
