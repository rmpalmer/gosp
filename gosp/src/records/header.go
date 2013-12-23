package records

import (
	"fmt"
)

type Header struct {
}

func (h *Header) HeaderSize() int {
	return 0
}

func (h *Header) DataSize() int {
	return 0
}

func (h *Header) Print() {
	fmt.Printf("header h=%d d=%d\n",0,0)
}