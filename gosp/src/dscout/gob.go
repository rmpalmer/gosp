package dscout

import (
	"records"
	"io"
	"errors"
	"fmt"
	"encoding/gob"
)

type GobMarshaler struct{}

func (GobMarshaler) MarshalRecord(writer io.Writer,
    rec *records.Rec) error {
    encoder := gob.NewEncoder(writer)
    fmt.Printf("here is magic number\n")
    fmt.Print(magicNumber)
    fmt.Printf("done with magic number\n")
    if err := encoder.Encode(magicNumber); err != nil {
        return err
    }
    if err := encoder.Encode(fileVersion); err != nil {
        return err
    }
    encoder.Encode(*rec.Header())
    encoder.Encode(*rec.Data())
    return nil
}

func (GobMarshaler) UnmarshalRecord(reader io.Reader) (*records.Rec,
    error) {
    decoder := gob.NewDecoder(reader)
    var magic int
    if err := decoder.Decode(&magic); err != nil {
        return nil, err
    }
    if magic != magicNumber {
        return nil, errors.New("cannot read non-seismic gob file")
    }
    var version int
    if err := decoder.Decode(&version); err != nil {
        return nil, err
    }
    if version > fileVersion {
        return nil, fmt.Errorf("version %d is too new to read", version)
    }
    var rec records.Rec
    err := decoder.Decode(rec)
    return rec, err
}