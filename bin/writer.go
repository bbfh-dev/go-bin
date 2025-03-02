package bin

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// Builder pattern for io.Writer.
//
// Use .Error() in the end of the chain to gather all errors that have occured
type Writer struct {
	buffer io.Writer
	errors []error
	count  int
}

func NewWriter(writer io.Writer) *Writer {
	return &Writer{
		buffer: writer,
		errors: []error{},
		count:  0,
	}
}

func (writer *Writer) Error() error {
	return errors.Join(writer.errors...)
}

func (writer *Writer) saveErr(err error) *Writer {
	writer.count++
	if err != nil {
		writer.errors = append(writer.errors, fmt.Errorf("(call %d) %w", writer.count, err))
	}
	return writer
}

func (writer *Writer) Write(data ...byte) *Writer {
	_, err := writer.buffer.Write(data)
	return writer.saveErr(err)
}

func (writer *Writer) WriteUint16(value uint16) *Writer {
	var data = make([]byte, 2)
	binary.BigEndian.PutUint16(data, value)
	_, err := writer.buffer.Write(data)
	return writer.saveErr(err)
}

func (writer *Writer) WriteUint32(value uint32) *Writer {
	var data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, value)
	_, err := writer.buffer.Write(data)
	return writer.saveErr(err)
}

func (writer *Writer) WriteUint64(value uint64) *Writer {
	var data = make([]byte, 8)
	binary.BigEndian.PutUint64(data, value)
	_, err := writer.buffer.Write(data)
	return writer.saveErr(err)
}

func (writer *Writer) WriteString(str string) *Writer {
	_, err := writer.buffer.Write(append([]byte(str), '\x00'))
	return writer.saveErr(err)
}

type Writable interface {
	Bytes() []byte
}

func (writer *Writer) WriteFixedSlice(values []Writable) *Writer {
	writer.WriteUint64(uint64(len(values)))
	switch len(values) {
	case 0:
		return writer
	case 1:
		item := values[0].Bytes()
		writer.WriteUint64(uint64(len(item)))
		writer.Write(item...)
	}
	return writer
}
