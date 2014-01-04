package records

import (

)

type Trace struct {
	Header []int
	Data []float64
}

func NewTrace(hlen, dlen int) *Trace {
	t := &Trace{make([]int, hlen), make([]float64, dlen)}
	return t
}

func (t *Trace) Rectyp() int {
	return 4095
}

func (t *Trace) Hlen() int {
	return len(t.Header)
}

func (t *Trace) Dlen() int {
	return len(t.Data)
}
