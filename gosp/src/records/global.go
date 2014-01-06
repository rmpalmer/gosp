package records

import (

)

const (
	GlobalType = 255
)

type Global struct {
	Tmin int
	Tmax int
	Dt   int
}

func (g *Global) Rectyp() int {
	return GlobalType
}

func NewGlobal(tmin, tmax, dt int) *Global {
	g := &Global{tmin, tmax, dt}
	return g
}
