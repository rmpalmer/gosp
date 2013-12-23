package records

import (
	"fmt"
	"bytes"
	"encoding/gob"
)

type Trace struct {
	headerSize	int
	dataSize	int
	header		[]int
	data		[]float64
}

type GobTrace struct {
}

func NewTrace(h, d int) *Trace {
	t := &Trace{h, d, nil, nil}
	t.header = make([]int, h)
	t.data = make([]float64, d)
	return t
}

func (h *Trace) HeaderSize() int {
	return h.headerSize
}

func (h *Trace) DataSize() int {
	return h.dataSize
}

func (h *Trace) Print() {
	fmt.Printf("trace h=%d d=%d\n",h.headerSize,h.dataSize)
}

func (h *Trace) Header() []int {
	return h.header
}

func (h *Trace) Data() []float64 {
	return h.data
}

func (h *Trace) GobEncode() ([]byte, error) {
	gobTrace := GobTrace { }
    var buffer bytes.Buffer
    encoder := gob.NewEncoder(&buffer)
    err := encoder.Encode(gobTrace)
    return buffer.Bytes(), err	
}

func (h *Trace) GobDecode(data []byte) error {
	return nil
}