package main 

import (
	"datgen"
	"filter"
	"sync"
	"fmt"
	"dscout"
	"dscin"
)	

func main() {
	fmt.Printf("Hello gosp version 0.1\n")
	
	// will probably want to have some sort of factory to create the operations
	// then Execute() is a method on the operation.   
	 
	waiter := &sync.WaitGroup{}
	/* writing */
	if (false) {
		d := datgen.NewDatgen(waiter,5)
		f := filter.NewFilter(waiter)
		o := dscout.NewDscout(waiter,"foobar.xml")
		d.Append(&f.Operation)
		f.Append(&o.Operation)
		go d.Execute()
		go f.Execute()
		go o.Execute()
	} else {
	/* reading */ 
		d := dscin.NewDscin(waiter,"foobar.xml")
		go d.Execute()
	}
	waiter.Wait()
	
	fmt.Printf("Goodbye gosp version 0.1\n")
}

