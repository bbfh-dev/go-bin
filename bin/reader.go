package bin

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
)

type Reader struct {
	buffer *bufio.Reader
	errors []error
}

func NewBufioReader(reader *bufio.Reader) *Reader {
	return &Reader{
		buffer: reader,
		errors: []error{},
	}
}

// Shorthand for creating a bufio.Reader from io.Reader
func NewReader(reader io.Reader) *Reader {
	return &Reader{
		buffer: bufio.NewReader(reader),
		errors: []error{},
	}
}

func (reader *Reader) Error() error {
	return errors.Join(reader.errors...)
}

func (reader *Reader) saveErr(err error) *Reader {
	if err != nil {
		reader.errors = append(reader.errors, err)
	}
	return reader
}

func (reader *Reader) ReadByte(out *byte) *Reader {
	var err error
	*out, err = reader.buffer.ReadByte()
	return reader.saveErr(err)
}

func (reader *Reader) ReadUint16(out *uint16) *Reader {
	var data = make([]byte, 2)
	_, err := reader.buffer.Read(data)
	*out = binary.BigEndian.Uint16(data)
	return reader.saveErr(err)
}

func (reader *Reader) ReadUint32(out *uint32) *Reader {
	var data = make([]byte, 4)
	_, err := reader.buffer.Read(data)
	*out = binary.BigEndian.Uint32(data)
	return reader.saveErr(err)
}

func (reader *Reader) ReadUint64(out *uint64) *Reader {
	var data = make([]byte, 8)
	_, err := reader.buffer.Read(data)
	*out = binary.BigEndian.Uint64(data)
	return reader.saveErr(err)
}

func (reader *Reader) ReadString(out *string) *Reader {
	str, err := reader.buffer.ReadString('\x00')
	*out = str[:len(str)-1]
	return reader.saveErr(err)
}

func (reader *Reader) ReadRawString(out *string) *Reader {
	var err error
	*out, err = reader.buffer.ReadString('\x00')
	return reader.saveErr(err)
}
