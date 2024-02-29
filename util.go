package goutil

import (
	"golang.org/x/exp/constraints"
)

type Longimetric[LengthT constraints.Integer] interface {
	Len() LengthT
}
