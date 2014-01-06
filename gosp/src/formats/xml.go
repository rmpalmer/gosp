package formats

import (
	"io"
	"encoding/xml"
	"records"
	"errors"
	"fmt"
)

type XmlMarshaler struct {
	encoder *xml.Encoder
	decoder *xml.Decoder
}

func (x *XmlMarshaler) InitFile(writer io.Writer) error {
	x.encoder = xml.NewEncoder(writer)
	if err := x.encoder.Encode(magicNumber); err != nil {
        return err
    }
    if err := x.encoder.Encode(fileVersion); err != nil {
        return err
    }
	return nil
}

func (x *XmlMarshaler) ValidateFile(reader io.Reader) (error) {
    x.decoder = xml.NewDecoder(reader)
    var magic int
    if err := x.decoder.Decode(&magic); err != nil {
        return err
    }
    if magic != magicNumber {
        return errors.New("cannot read non-trace gob file")
    } else {
    	fmt.Printf("read magic number %d\n",magic)
    }
    var version int
    if err := x.decoder.Decode(&version); err != nil {
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

func (x *XmlMarshaler) MarshalRecord(rec records.Record) error {
	fmt.Printf("starting gob MarshalRecord\n")
	err := x.encoder.Encode(rec.Rectyp())
	err  = x.encoder.Encode(rec)
	return err
}

func (x *XmlMarshaler) UnmarshalRecord() (records.Record, error ) {
	fmt.Printf("starting gob UnmarshalRecord\n")
	var recid int
	var tr *records.Trace
	var gl *records.Global
	err := x.decoder.Decode(&recid)
	switch recid {
		case records.TraceType:
			tr = new(records.Trace)
			err = x.decoder.Decode(tr)
			return tr, err
		case records.GlobalType:
			gl = new(records.Global)
			err = x.decoder.Decode(gl)
			return gl, err
		}
	return nil, err
}
