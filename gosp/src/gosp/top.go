package main 

import (
	"fmt"
	"datgen"
	"filter"
	"dscout"
	"records"
	"sync"
)	

func main() {
	fmt.Printf("Hello gosp version 0.1\n")
	
	waiter := &sync.WaitGroup{}
	
	fmt.Printf("making a trace\n")
	var r interface{}
	r = records.NewTrace(5,7)
	
	switch r.(type) {
		case *records.Trace:
			fmt.Printf("it really is a trace!\n")
	}
	
	d := datgen.New(5,100,1000,waiter)
	f := filter.New(60.0, waiter, d)
	g := filter.New(50.0, waiter, f)
	o := dscout.New("foobar.gob",waiter, g)
	
	go d.Exec()
	go f.Exec()
	go g.Exec()
	go o.Exec()
	
	waiter.Wait()
	
	fmt.Printf("Goodbye gosp version 0.1\n")
}

