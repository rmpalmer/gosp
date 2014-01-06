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
		case records.TraceType:
			fmt.Printf("attempting to decode trace\n")
			tr = new(records.Trace)
			err = g.decoder.Decode(tr)
			return tr, err
		case records.GlobalType:
			fmt.Printf("attempting to decode global\n")
			gl = new(records.Global)
			err = g.decoder.Decode(gl)
			return gl, err
		}
	return nil, err
}
