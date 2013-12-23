package records

import (

)

type Rec interface {
	HeaderSize() int
	DataSize() int
	Header() []int
	Data() []float64
	Print()
	GobEncode() ([]byte, error)
	GobDecode([]byte) error 
}

