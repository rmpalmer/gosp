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

func (x *XmlMarshaler) MarshalTrace(trace *records.Trace) error {
    fmt.Printf("starting xml MarshalTrace\n")
    err := x.encoder.Encode(trace)
    fmt.Printf("done calling encoder.Encode %s\n",err)
    return err
}

func (x *XmlMarshaler) UnmarshalTrace() (*records.Trace, error) {
    fmt.Printf("starting xml UnmarshalTrace\n")
    var trace records.Trace
    fmt.Printf("unmarshaller about to decode trace\n")
    err := x.decoder.Decode(&trace)
    fmt.Printf("done calling decoder.Decode %s\n",err)
    return &trace, err
}

