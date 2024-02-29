package goutil

import (
	"reflect"
	"strings"
)

type NilArgError struct {
	Parameter string
	Type reflect.Type
	Method string
}

func(err *NilArgError) Error() string {
	var builder strings.Builder
	builder.WriteString("Parameter")
	if len(err.Parameter) > 0 {
		builder.WriteString(" '")
		builder.WriteString(err.Parameter)
		builder.WriteRune('\'')
	}
	if err.Type != nil {
		builder.WriteString(" to method")
		if len(err.Method) > 0 {
			builder.WriteString(err.Type.String())
			builder.WriteRune('.')
			builder.WriteString(err.Method)
		}
	} else {
		builder.WriteString(" to function")
		if len(err.Method) > 0 {
			builder.WriteRune(' ')
			builder.WriteString(err.Method)
		}
	}
	builder.WriteString(" must not be nil, but was")
	return builder.String()
}

func NewNilArgError(param string, target any, method string) error {
	return &NilArgError {
		Parameter: param,
		Type: reflect.TypeOf(target),
		Method: method,
	}
}

type NilTargetError struct {
	Type reflect.Type
	Method string
}

func(err *NilTargetError) Error() string {
	var builder strings.Builder
	builder.WriteString("Target for method")
	if err.Type != nil {
		if len(err.Method) > 0 {
			builder.WriteString(err.Type.String())
			builder.WriteRune(' ')
			builder.WriteString(err.Method)
		} else {
			builder.WriteString(" of ")
			builder.WriteString(err.Type.String())
		}
	}
	builder.WriteString(" must not be nil, but was")
	return builder.String()
}

func NewNilTargetError(target any, method string) error {
	return &NilTargetError {
		Type: reflect.TypeOf(target),
		Method: method,
	}
}

type NilPropError struct {
	Property string
	Type reflect.Type
}

func(err *NilPropError) Error() string {
	var builder strings.Builder
	builder.WriteString("Property")
	if len(err.Property) > 0 {
		builder.WriteString(" '")
		builder.WriteString(err.Property)
		builder.WriteRune('\'')
	}
	if err.Type != nil {
		builder.WriteString(" of ")
		builder.WriteString(err.Type.String())
	}
	builder.WriteString(" must not be nil, but was")
	return builder.String()
}

func NewNilPropError(prop string, target any) error {
	return &NilPropError {
		Property: prop,
		Type: reflect.TypeOf(target),
	}
}
