package internal

import (
	"sync/atomic"
)

type Progress struct {
	Total       int64
	Done        int64
	Interrupted bool
	data        chan Progress
}

func NewProgress(capacity int) *Progress {
	return &Progress{
		Total:       0,
		Done:        0,
		Interrupted: false,
		data:        make(chan Progress, capacity),
	}
}

func (p *Progress) incTotal() {
	atomic.AddInt64(&p.Total, 1)
	p.data <- *p
}

func (p *Progress) incDone() {
	atomic.AddInt64(&p.Done, 1)
	p.data <- *p
}

func (p *Progress) interrupt() {
	p.Interrupted = true
	p.data <- *p
}

func (p *Progress) close() {
	close(p.data)
}
