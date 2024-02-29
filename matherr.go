package goutil

import (
	"fmt"
	"strings"
	"reflect"
)

type OverflowsIntError struct {
	Meaning string
	Old int
	Delta int
}

func(err *OverflowsIntError) Error() string {
	var builder strings.Builder
	if len(err.Meaning) > 0 {
		builder.WriteString(err.Meaning)
	} else if err.Delta < 0 {
		builder.WriteString("Subtraction")
	} else {
		builder.WriteString("Addition")
	}
	if err.Delta < 0 {
		builder.WriteString(" underflows ")
	} else {
		builder.WriteString(" overflows ")
	}
	builder.WriteString(" int: ")
	builder.WriteString(fmt.Sprintf("%d + %d == %d", err.Old, err.Delta, err.Old + err.Delta))
	return builder.String()
}

func NewOverflowsIntError(meaning string, old int, delta int) error {
	return &OverflowsIntError {
		Meaning: meaning,
		Old: old,
		Delta: delta,
	}
}

func NewOverflowsIntPropError(prop string, target any, old int, delta int) error {
	var builder strings.Builder
	builder.WriteString("Property")
	if len(prop) > 0 {
		builder.WriteString(" '")
		builder.WriteString(prop)
		builder.WriteRune('\'')
	}
	ty := reflect.TypeOf(target)
	if ty != nil {
		builder.WriteString(" of ")
		builder.WriteString(ty.String())
	}
	return &OverflowsIntError {
		Meaning: builder.String(),
		Old: old,
		Delta: delta,
	}
}
