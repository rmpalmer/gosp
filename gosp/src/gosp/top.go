package main 

import (
	//"datgen"
	//"filter"
	"sync"
	"fmt"
	//"dscout"
	"dscin"
)	

func main() {
	fmt.Printf("Hello gosp version 0.1\n")
	
	// will probably want to have some sort of factory to create the operations
	// then Execute() is a method on the operation.   
	 
	waiter := &sync.WaitGroup{}
	/* writing
	d := datgen.NewDatgen(waiter,5)
	f := filter.NewFilter(waiter)
	o := dscout.NewDscout(waiter,"foobar.gob")
	d.Append(&f.Operation)
	f.Append(&o.Operation)
	go d.Execute()
	go f.Execute()
	go o.Execute()
	*/
	/* reading */
	d := dscin.NewDscin(waiter,"foobar.gob")
	go d.Execute()
	waiter.Wait()
	
	fmt.Printf("Goodbye gosp version 0.1\n")
}

