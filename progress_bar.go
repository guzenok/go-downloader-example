package main

import (
	"github.com/cheggaaa/pb"
)

var bar *pb.ProgressBar

func InitBar() {
	bar = pb.New(0).Start()
}

func IncBarTotal() {
	if bar != nil {
		bar.SetTotal(bar.Total() + 1)
	}
}

func IncBarValue() {
	if bar != nil {
		bar.Increment()
	}
}

func FinishBar() {
	if bar != nil {
		bar.Finish()
	}
}
