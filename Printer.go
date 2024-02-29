package goutil

import (
	"io"
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

type Printer interface {
	io.Writer
	io.ByteWriter
	RuneWriter
	PrintLike
	Longimetric[int]
	Grow(uint)
	WriteString(string) (int, error)
}

// by strings.Builder

type StringBuilderPrinter struct {
	builder *strings.Builder
}

func NewStringBuilderPrinter(builder *strings.Builder) Printer {
	return &StringBuilderPrinter {
		builder: builder,
	}
}

func(prn *StringBuilderPrinter) init() {
	if prn.builder == nil {
		prn.builder = &strings.Builder{}
	}
}

func(prn *StringBuilderPrinter) Write(bytes []byte) (int, error) {
	if prn == nil {
		return 0, NewNilTargetError(&StringBuilderPrinter{}, "Write")
	}
	prn.init()
	return prn.builder.Write(bytes)
}

func(prn *StringBuilderPrinter) WriteByte(b byte) error {
	if prn == nil {
		return NewNilTargetError(&StringBuilderPrinter{}, "WriteByte")
	}
	prn.init()
	return prn.builder.WriteByte(b)
}

func(prn *StringBuilderPrinter) WriteRune(r rune) (int, error) {
	if prn == nil {
		return 0, NewNilTargetError(&StringBuilderPrinter{}, "WriteRune")
	}
	prn.init()
	return prn.builder.WriteRune(r)
}

func(prn *StringBuilderPrinter) Print(args ...any) error {
	if prn == nil {
		return NewNilTargetError(&StringBuilderPrinter{}, "Print")
	}
	prn.init()
	_, err := prn.builder.WriteString(fmt.Sprint(args...))
	return err
}

func(prn *StringBuilderPrinter) Printf(format string, args ...any) error {
	if prn == nil {
		return NewNilTargetError(&StringBuilderPrinter{}, "Printf")
	}
	prn.init()
	_, err := prn.builder.WriteString(fmt.Sprintf(format, args...))
	return err
}

func(prn *StringBuilderPrinter) Println(args ...any) error {
	if prn == nil {
		return NewNilTargetError(&StringBuilderPrinter{}, "Println")
	}
	prn.init()
	_, err := prn.builder.WriteString(fmt.Sprintln(args...))
	return err
}

func(prn *StringBuilderPrinter) Len() int {
	if prn == nil {
		return 0
	}
	prn.init()
	return prn.builder.Len()
}

func(prn *StringBuilderPrinter) Grow(delta uint) {
	if prn == nil {
		return
	}
	prn.init()
	if delta > uint(math.MaxInt) {
		delta = uint(math.MaxInt)
	}
	prn.builder.Grow(int(delta))
}

func(prn *StringBuilderPrinter) WriteString(s string) (int, error) {
	if prn == nil {
		return 0, NewNilTargetError(&StringBuilderPrinter{}, "WriteString")
	}
	prn.init()
	return prn.builder.WriteString(s)
}

func(prn *StringBuilderPrinter) String() (string, error) {
	if prn == nil {
		return "", NewNilTargetError(&StringBuilderPrinter{}, "String")
	}
	prn.init()
	return prn.builder.String(), nil
}

var _ Printer = &StringBuilderPrinter{}

// by io.Writer

type WriterPrinter struct {
	writer io.Writer
	flags WriterPrinterFlags
	buffer []byte
	written int
}

type WriterPrinterFlags uint

const (
	WRPRNFL_PRIVATE_BUFFER WriterPrinterFlags = 1 << iota
)

func NewWriterPrinter(writer io.Writer, flags WriterPrinterFlags, written int) Printer {
	return &WriterPrinter {
		writer: writer,
		flags: flags,
		written: written,
	}
}

func(prn *WriterPrinter) init() {
	if prn.writer == nil {
		prn.writer = Discard{}
	}
}

func(prn *WriterPrinter) buf() []byte {
	if (prn.flags & WRPRNFL_PRIVATE_BUFFER) != 0 {
		return make([]byte, 4)
	}
	if prn.buffer == nil {
		prn.buffer = make([]byte, 4)
	}
	return prn.buffer
}

func(prn *WriterPrinter) advance(err error, count int) error {
	if count <= 0 {
		return err
	}
	nextWritten := prn.written + count
	if nextWritten < count {
		if err == nil {
			err = NewOverflowsIntPropError("written", prn, prn.written, count)
		}
	} else {
		prn.written = nextWritten
	}
	return err
}

func(prn *WriterPrinter) Write(bytes []byte) (count int, err error) {
	if prn == nil {
		err = NewNilTargetError(&WriterPrinter{}, "Write")
		return
	}
	prn.init()
	count, err = prn.writer.Write(bytes)
	err = prn.advance(err, count)
	return
}

func(prn *WriterPrinter) WriteByte(b byte) error {
	if prn == nil {
		return NewNilTargetError(&WriterPrinter{}, "WriteByte")
	}
	prn.init()
	buf := prn.buf()
	buf[0] = b
	count, err := prn.writer.Write(buf[0:1])
	err = prn.advance(err, count)
	return err
}

func(prn *WriterPrinter) WriteRune(r rune) (int, error) {
	if prn == nil {
		return 0, NewNilTargetError(&WriterPrinter{}, "WriteRune")
	}
	prn.init()
	buf := prn.buf()
	size := utf8.EncodeRune(buf, r)
	count, err := prn.writer.Write(buf[0:size])
	err = prn.advance(err, count)
	return count, err
}

func(prn *WriterPrinter) Print(args ...any) error {
	if prn == nil {
		return NewNilTargetError(&WriterPrinter{}, "Print")
	}
	prn.init()
	_, err := prn.writeString(fmt.Sprint(args...))
	return err
}

func(prn *WriterPrinter) Printf(format string, args ...any) error {
	if prn == nil {
		return NewNilTargetError(&WriterPrinter{}, "Printf")
	}
	prn.init()
	_, err := prn.writeString(fmt.Sprintf(format, args...))
	return err
}

func(prn *WriterPrinter) Println(args ...any) error {
	if prn == nil {
		return NewNilTargetError(&WriterPrinter{}, "Println")
	}
	prn.init()
	_, err := prn.writeString(fmt.Sprintln(args...))
	return err
}

func(prn *WriterPrinter) Len() int {
	if prn == nil {
		return 0
	}
	return prn.written
}

func(prn *WriterPrinter) Grow(uint) {}

func(prn *WriterPrinter) writeString(s string) (int, error) {
	count, err := prn.writer.Write([]byte(s))
	err = prn.advance(err, count)
	return count, err
}

func(prn *WriterPrinter) WriteString(s string) (int, error) {
	if prn == nil {
		return 0, NewNilTargetError(&WriterPrinter{}, "WriteString")
	}
	prn.init()
	return prn.writeString(s)
}

var _ Printer = &WriterPrinter{}
