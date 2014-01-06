package records

import (

)

type Global struct {
	Tmin int
	Tmax int
	Dt   int
}

func (g *Global) Rectyp() int {
	return 255
}

func NewGlobal(tmin, tmax, dt int) *Global {
	g := &Global{tmin, tmax, dt}
	return g
}
