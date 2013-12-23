package datgen

import (
	"records"
	"fmt"
	"sync"
)

type Datgen struct {
	count int
	hSize int
	dSize int
	waiter *sync.WaitGroup
	successor chan records.Rec
}

func New(count, hSize, dSize int, waiter *sync.WaitGroup) *Datgen {
	datgen := Datgen{count,hSize,dSize,waiter,nil}
	datgen.waiter.Add(1)
	return &datgen
}

func (d *Datgen) Follow(c chan records.Rec) {
	if (c != nil) {
		fmt.Print("adding a follower\n")
		d.successor = c
	} else {
		fmt.Print("follower channel was nil!\n")
	}
}

func (d *Datgen) Exec() {
	
	fmt.Printf("DATGEN exec start\n")
	t:= records.NewTrace(d.hSize,d.dSize)
	for i:=0; i<d.count; i++ {
		fmt.Printf("datgen generating trace %d\n",i)
		if (d.successor != nil) {
			fmt.Print("datgen sending to my follower\n")
			d.successor <- t
		} else {
			fmt.Printf("datgen no follower\n")
		}
	}
	if (d.successor != nil) {
		fmt.Printf("datgen sending eod\n")
		var eod *records.Eod
		d.successor <- eod
	} else {
		fmt.Printf("datgen hiding eod\n")
	}
	d.waiter.Done()
}