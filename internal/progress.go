package internal

type Status struct {
	Total       int64
	Done        int64
	Interrupted bool
}

type Progress struct {
	Status
	total       chan int64
	done        chan int64
	interrupted chan bool
	status      chan Status
}

func NewProgress(capacity int) *Progress {
	p := &Progress{
		total:       make(chan int64, capacity),
		done:        make(chan int64, capacity),
		interrupted: make(chan bool, capacity),
		status:      make(chan Status, capacity),
	}
	go p.notify()
	return p
}

func (p *Progress) incTotal() {
	p.total <- 1
}

func (p *Progress) incDone() {
	p.done <- 1
}

func (p *Progress) interrupt() {
	p.interrupted <- true
}

func (p *Progress) close() {
	p.interrupted <- false
}

func (p *Progress) notify() {
	closed := false
	for {
		select {
		case val := <-p.total:
			p.Total += val
		case val := <-p.done:
			p.Done += val
		case val := <-p.interrupted:
			p.Interrupted = p.Interrupted || val
			closed = true
		}
		p.status <- p.Status

		if closed {
			close(p.status)
			break
		}
	}
}
