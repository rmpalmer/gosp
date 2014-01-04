package filter

import (
	"fmt"
	"operation"
	"sync"
	"records"
)

type Filter struct {
	opera.Operation
}

func NewFilter(waiter *sync.WaitGroup) *Filter {
	f := &Filter{}
	f.Operation.Waiter = waiter
	f.Operation.Waiter.Add(1)
	return f
}


func (f *Filter) Execute() {
	fmt.Printf("filter execute\n")
	if (f.Source != nil) {
		for rec := range *f.Source {
			switch recType := rec.(type) {
				case *gprec.Global:
					fmt.Printf("filter received global\n") 
				case *gprec.Trace:
					t := rec.(*gprec.Trace)
					fmt.Printf("filter received trace %d\n",t.Header[0])
					
				default:
					fmt.Printf("filter received unrecognized type %v\n",recType) 
			}
			if (f.Sink != nil) {
				f.Sink <- rec
			}
		}
		if (f.Sink != nil) {
			close(f.Sink)
		}
	} 
	f.Operation.Waiter.Done()
}