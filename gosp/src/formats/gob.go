package formats

import (
	"io"
	"encoding/gob"
	"records"
	"errors"
	"fmt"
)

type GobMarshaler struct{
	encoder		*gob.Encoder
	decoder		*gob.Decoder
}

func (g *GobMarshaler) InitFile(writer io.Writer) error { 
	fmt.Printf("creating gob encoder for init file\n")
    g.encoder = gob.NewEncoder(writer)
    if err := g.encoder.Encode(magicNumber); err != nil {
        return err
    }
    if err := g.encoder.Encode(fileVersion); err != nil {
        return err
    }
    fmt.Printf("init gob file done\n")
    return nil
}

func (g *GobMarshaler) ValidateFile(reader io.Reader) (error) {
    g.decoder = gob.NewDecoder(reader)
    var magic int
    if err := g.decoder.Decode(&magic); err != nil {
        return err
    }
    if magic != magicNumber {
        return errors.New("cannot read non-trace gob file")
    } else {
    	fmt.Printf("read magic number %d\n",magic)
    }
    var version int
    if err := g.decoder.Decode(&version); err != nil {
        return err
    }
    if version > fileVersion {
        return fmt.Errorf("version %d is too new to read", version)
    } else {
    	fmt.Printf("read file version %d\n",version)
    }
    fmt.Printf("ValidateFile no errors\n")
	return nil
}

func (g *GobMarshaler) MarshalGlobal(global *records.Global) error {
	fmt.Printf("starting gob MarshalGlobal\n")
	err := g.encoder.Encode(global.Rectyp())
	err  = g.encoder.Encode(global)
	return err
}

func (g *GobMarshaler) MarshalTrace(trace *records.Trace) error {
    fmt.Printf("starting gob MarshalTrace\n")
    err := g.encoder.Encode(trace.Rectyp())
    err  = g.encoder.Encode(trace)
    //fmt.Printf("done calling encoder.Encode %s\n",err)
    return err
}

func (g *GobMarshaler) UnmarshalGlobal() (*records.Global, error) {
	fmt.Printf("starting gob UnmarshalGlobal\n")
	var global records.Global
	fmt.Printf("unmarshaller about to decode global\n")
	err := g.decoder.Decode(&global)
	return &global, err
}

func (g *GobMarshaler) UnmarshalTrace() (*records.Trace, error) {
    fmt.Printf("starting gob UnmarshalTrace\n")
    var trace records.Trace
    fmt.Printf("unmarshaller about to decode trace\n")
    err := g.decoder.Decode(&trace)
    //fmt.Printf("done calling decoder.Decode %s\n",err)
    return &trace, err
}

func (g *GobMarshaler) MarshalRecord(rec records.Record) error {
	fmt.Printf("starting gob MarshalRecord\n")
	err := g.encoder.Encode(rec.Rectyp())
	err  = g.encoder.Encode(rec)
	return err
}

func (g *GobMarshaler) UnmarshalRecord() (records.Record, error ) {
	fmt.Printf("starting gob UnmarshalRecord\n")
	var recid int
	var tr *records.Trace
	var gl *records.Global
	err := g.decoder.Decode(&recid)
	switch recid {
		case 4095:
			tr = new(records.Trace)
			err = g.decoder.Decode(tr)
			return tr, err
		case 255:
			gl = new(records.Global)
			err = g.decoder.Decode(gl)
			return gl, err
		}
	return nil, err
}
