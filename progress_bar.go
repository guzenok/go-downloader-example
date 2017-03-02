// progress_bar
package main

import (
	"sync/atomic"

	"gopkg.in/cheggaaa/pb.v1"
)

var bar *pb.ProgressBar

func InitBar() {
	bar = pb.New(0).SetUnits(pb.U_NO).Start()
}

func IncBarTotal() {
	if bar != nil {
		atomic.AddInt64(&bar.Total, 1)
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
