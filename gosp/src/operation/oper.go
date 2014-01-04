package operation

import (
	"records"
	"sync"
)

type Operation struct {
	// records come into the operation via its source
	Source *chan records.Record
	
	// and leave via its sink
	Sink chan records.Record
	
	Waiter *sync.WaitGroup
}

// operation a is the predecessor
// operation b is the new successor
func (a *Operation) Append(b *Operation) {
	if (a.Sink == nil) {
		// make a route out of the predecessor
		a.Sink = make(chan records.Record)
		
		// have the successor connect to it.
		b.Source = &a.Sink
	}
}

