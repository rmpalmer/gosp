package records

import (
	"fmt"
	"bytes"
	"encoding/gob"
)

type Eod struct {
	header		[]int
	data		[]float64	
}

func (h *Eod) HeaderSize() int {
	return 0
}

func (h *Eod) DataSize() int {
	return 0
}

func (h *Eod) Print() {
	fmt.Printf("End-of-data\n")
}

func (h *Eod) Header() []int {
	return h.header
}

func (h *Eod) Data() []float64 {
	return h.data
}

func (h *Eod) GobEncode() ([]byte, error) {
    var buffer bytes.Buffer
    encoder := gob.NewEncoder(&buffer)
    err := encoder.Encode(h)
    return buffer.Bytes(), err	
}

func (h *Eod) GobDecode(data []byte) error {
	return nil
}