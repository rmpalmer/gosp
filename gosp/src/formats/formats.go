package formats

import (
	"io"
	"records"
)

// rmp but I want these methods to take pointer receivers.

type RecordMarshaler interface {
	InitFile(writer io.Writer) error
	ValidateFile(reader io.Reader) error
	MarshalTrace(trace *records.Trace) error
	UnmarshalTrace() (*records.Trace, error)
}

