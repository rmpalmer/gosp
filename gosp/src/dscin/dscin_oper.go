package dscin

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
	"log"
)

type Dscin struct {
	operation.Operation
	closer func()
	marshaler formats.RecordMarshaler
}

func NewDscin (waiter *sync.WaitGroup, filename string) *Dscin {
	d := new(Dscin)
	d.Operation.Waiter = waiter
	d.Operation.Waiter.Add(1)
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	d.closer = func() {
		file.Close()
	}
	var reader io.ReadCloser = file
	var uncompressor *gzip.Reader
	if strings.HasSuffix(filename, ".gz") {
		uncompressor, err = gzip.NewReader(file)
		d.closer = func() { uncompressor.Close(); file.Close() }
		reader = uncompressor
	}
	uncompressed_name := strings.TrimRight(filename, ".gz")
	switch {
		case strings.HasSuffix(uncompressed_name, ".gob"):
			d.marshaler = new(formats.GobMarshaler)
		case strings.HasSuffix(uncompressed_name, ".xml"):
			d.marshaler = new(formats.XmlMarshaler)
	}
	if (d.marshaler != nil) {
		d.marshaler.ValidateFile(reader)
	}
	return d
}

func (d *Dscin) Execute() {
	fmt.Printf("dscin execute\n")
	var r records.Record
	for {
		r = d.HandleRecord()
		if (r == nil) {
			fmt.Printf("dscin execute break\n")
			break
		} else if (d.Sink != nil) {
			d.Sink <- r
		}
	}
	if (d.Sink != nil) {
		close(d.Sink)
	}
	d.Operation.Waiter.Done()
}

func (d *Dscin) HandleRecord() records.Record {
	r, err := d.marshaler.UnmarshalRecord()
	if (err != nil) {
		fmt.Printf("error returned from unmarshalRecord, found in HandleRecord\n")
		log.Print(err)
		return nil
	}
	return r
} 