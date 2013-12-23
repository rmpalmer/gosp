package operation

import (
	"records"
)

type Oper interface {
	Follow(c chan records.Rec)
}