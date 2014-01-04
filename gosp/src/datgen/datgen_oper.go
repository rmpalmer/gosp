package datgen

import (
	"fmt"
	"records"
	"operation"
	"sync"
)

type Datgen struct {
	operation.Operation
	nrecs int
}

func NewDatgen(waiter *sync.WaitGroup, n int) *Datgen {
	d := &Datgen{}
	d.nrecs = n
	d.Operation.Waiter = waiter
	d.Operation.Waiter.Add(1)
	return d
}

func (d *Datgen) Execute() {
	fmt.Printf("d will generate %d recs\n",d.nrecs)
	var t *records.Trace
	var g *records.Global
	g = records.NewGlobal(0, 4, 1000)
	if (d.Sink != nil) {
		d.Sink <- g
	}
	for i := 0; i<d.nrecs; i++ {
		t = records.NewTrace(64,1024)
		t.Header[0] = i
		fmt.Printf("generate %d\n",t.Header[0])
		if (d.Sink != nil) {
			d.Sink <- t
		}
	}
	if (d.Sink != nil) {
		close(d.Sink)
	}
	d.Operation.Waiter.Done()
}