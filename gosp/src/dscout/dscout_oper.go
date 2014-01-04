package dscout

import (
	"fmt"
	"operation"
	"records"
	"sync"
	"io"
	"os"
	"strings"
	"compress/gzip"
	"path/filepath"
)

type RecordMarshaller interface {
	MarshalRecord(writer io.Writer, rec *records.Rec) error 
}

type RecordUnmarshaller interface {
	UnmarshalRecord(reader io.Reader) (*records.Rec, error)
}

type Dscout struct {
	fname stringMarshalTrace
	waiter *sync.WaitGroup
	predecessor chan records.Rec
	successor chan records.Rec
}

func New(fname string, waiter *sync.WaitGroup, pred operation.Oper) *Dscout {
	dscout := &Dscout{fname, waiter, nil, nil}
	dscout.waiter.Add(1)
	dscout.predecessor = make(chan records.Rec)
	pred.Follow(dscout.predecessor)
	return dscout
}

func (f *Dscout) Follow(c chan records.Rec) {
	f.successor = c
}

func (d *Dscout) Exec() {
	done := false
	fmt.Print(done)
	fmt.Printf("dscout Exec Start\n")
	
	file, closer, err := createOutFile(d.fname)
	
	var marshaler RecordMarshaller
	
    switch suffixOf(d.fname) {
    case ".gob":
        marshaler = GobMarshaler{}
//    case ".txt":
//        marshaler = TxtMarshaler{}
    }
	
	if (closer != nil) {
		defer closer()
	}
	if (err != nil) {
		done = true
	}	
	
	for !done {
		r := <- d.predecessor
		switch  r.(type) {
		case *records.Eod:
			fmt.Printf("dscout read eod\n")
			done = true
		case *records.Trace:
			fmt.Printf("dscout read record\n")
		}
		marshaler.MarshalRecord(file, *r)
		if (d.successor != nil) {
			fmt.Printf("dscout sending to successor\n")
			d.successor <- r
		}
	}
	d.waiter.Done()
}

func createOutFile(fname string) (io.WriteCloser, func(), error) {
	file, err := os.Create(fname)
	if err != nil {
        return nil, nil, err
    }
	closer := func() {
		fmt.Printf("closing the file\n")
		file.Close()
	}
    var writer io.WriteCloser = file
    var compressor *gzip.Writer
    if strings.HasSuffix(fname, ".gz") {
        compressor = gzip.NewWriter(file)
        closer = func() { compressor.Close(); file.Close() }
        writer = compressor
    }
    return writer, closer, nil
}

func suffixOf(filename string) string {
    suffix := filepath.Ext(filename)
    if suffix == ".gz" {
        suffix = filepath.Ext(filename[:len(filename)-3])
    }
    return suffix
}
