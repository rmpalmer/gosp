package main 

import (
	"datgen"
	"filter"
	"sync"
	"fmt"
	"dscout"
)	

func main() {
	fmt.Printf("Hello gosp version 0.1\n")
	
	waiter := &sync.WaitGroup{}
	d := datgen.NewDatgen(waiter,5)
	f := filter.NewFilter(waiter)
	o := dscout.NewDscout(waiter,"foobar.gob")
	d.Append(&f.Operation)
	f.Append(&o.Operation)
	go d.Execute()
	go f.Execute()
	go o.Execute()
	waiter.Wait()
	
	fmt.Printf("Goodbye gosp version 0.1\n")
}

