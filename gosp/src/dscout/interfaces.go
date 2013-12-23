package dscout

import (
	"records"
	"io"
)

type RecordMarshaler interface {
    MarshalRecord(writer io.Writer, rec *records.Rec) error
}

type RecordUnmarshaler interface {
    UnmarshalRecord(reader io.Reader) (*records.Rec, error)
}

