package main 

import (
	"datgen"
	"filter"
	"sync"
	"fmt"
)	

func main() {
	fmt.Printf("Hello gosp version 0.2\n")
	
	waiter := &sync.WaitGroup{}
	d := datgen.NewDatgen(waiter,5)
	f := filter.NewFilter(waiter)
	d.Append(&f.Operation)
	go d.Execute()
	go f.Execute()
	waiter.Wait()
	
	fmt.Printf("Goodbye gosp version 0.2\n")
}

