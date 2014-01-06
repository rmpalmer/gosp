package formats

import (
	"io"
	"records"
)

type RecordMarshaler interface {
	InitFile(writer io.Writer) error
	ValidateFile(reader io.Reader) error
	//MarshalTrace(trace *records.Trace) error
	//UnmarshalTrace() (*records.Trace, error)
	//MarshalGlobal(global *records.Global) error
	//UnmarshalGlobal() (*records.Global, error)
	MarshalRecord(rec records.Record) error
	UnmarshalRecord() (records.Record, error)
}

