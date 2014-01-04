package dscout

import (
	"os"
	"compress/gzip"
	"io"
	"sync"
	"strings"
	"records"
	"formats"
	"operation"
	"fmt"
)

type Dscout struct {
	operation.Operation
	closer		func()
	marshaler	formats.RecordMarshaler
}

func NewDscout (waiter *sync.WaitGroup, filename string) *Dscout {
	d := new(Dscout)
	d.Operation.Waiter = waiter
	d.Operation.Waiter.Add(1)
	file, err := os.Create(filename)
	if err != nil {
		return nil
	}
	d.closer = func() {
		file.Close()
	}
	var writer io.WriteCloser = file
	var compressor *gzip.Writer
	if strings.HasSuffix(filename, ".gz") {
		compressor = gzip.NewWriter(file)
		d.closer = func() { compressor.Close(); file.Close() }
		writer = compressor
	}
	uncompressed_name := strings.TrimRight(filename, ".gz")
	switch {
		case strings.HasSuffix(uncompressed_name, ".gob"):
			d.marshaler = new(formats.GobMarshaler)
		case strings.HasSuffix(uncompressed_name, ".xml"):
			d.marshaler = new(formats.XmlMarshaler)
	}
	if (d.marshaler != nil) {
		d.marshaler.InitFile(writer)
	}
	return d
}

func (d *Dscout) Execute() {
	fmt.Printf("dscout execute\n")
	if (d.Source != nil) {
		for rec := range *d.Source {
			switch recType := rec.(type) {
				case *records.Global:
					fmt.Printf("dscout received global\n") 
				case *records.Trace:
					t := rec.(*records.Trace)
					fmt.Printf("dscout received trace %d\n",t.Header[0])
					d.HandleTrace(t)
				default:
					fmt.Printf("dscout received unrecognized type %v\n",recType) 
			}
			if (d.Sink != nil) {
				d.Sink <- rec
			}
		}
		if (d.Sink != nil) {
			close(d.Sink)
		}
		d.closer()
	} 
	d.Operation.Waiter.Done()
}

func (d *Dscout) HandleTrace(trace *records.Trace) {
	d.marshaler.MarshalTrace(trace)
}

