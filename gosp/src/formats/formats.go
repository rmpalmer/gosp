package formats

import (
	"io"
	"records"
)

type RecordMarshaler interface {
	InitFile(writer io.Writer) error
	ValidateFile(reader io.Reader) error
	MarshalRecord(rec records.Record) error
	UnmarshalRecord() (records.Record, error)
}

