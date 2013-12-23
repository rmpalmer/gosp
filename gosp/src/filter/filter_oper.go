package filter

import (
	"fmt"
	"operation"
	"records"
	"sync"
)

type Filter struct {
	f1 float64
	waiter *sync.WaitGroup
	predecessor chan records.Rec
	successor chan records.Rec
}

func New(f1 float64, waiter *sync.WaitGroup, pred operation.Oper) *Filter {
	filter := &Filter{f1, waiter, nil, nil}
	filter.waiter.Add(1)
	filter.predecessor = make(chan records.Rec)
	pred.Follow(filter.predecessor)
	return filter
}

func (f *Filter) Follow(c chan records.Rec) {
	f.successor = c
}

func (f *Filter) Exec() {
	done := false
	fmt.Print(done)
	fmt.Printf("FILTER Exec Start\n")
	for !done {
		r := <- f.predecessor
		switch  r.(type) {
		case *records.Eod:
			fmt.Printf("filter read eod\n")
			done = true
		case *records.Trace:
			fmt.Printf("filter read record\n")
		}
		if (f.successor != nil) {
			fmt.Printf("filter sending to successor\n")
			f.successor <- r
		}
	}
	f.waiter.Done()
}