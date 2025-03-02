package bin

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"strings"
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

//   - The size of text as a 32bit integer
//   - If the size is positive, we're dealing with a windows-1252 encoding, so
//     we don't need to do anything to get the number of bytes that the string
//     consumes (as windows-1252 is a 8bit encoding).
//   - If the size is negative, the string is encoded with UTF-16, so multiply
//     it by -2 to get the number of bytes needed to read the string.
//   - Consume the number of bytes determined, but drop the last letter (1 byte
//     for windows-1252, 2 for UTF-16) as this will be a null character which we
//     don't want.
func (reader *Reader) ReadReplayString(out *string) *Reader {
	var uSize uint32
	reader.ReadUint32(&uSize)
	size := int32(uSize)

	if size < 0 { // Encoded with UTF-16
		size *= -2
	}

	var str = make([]byte, size)
	_, err := reader.buffer.Read(str)
	reader.saveErr(err)

	*out = strings.TrimRight(string(str[:len(str)-1]), "\x00")
	return reader.saveErr(err)
}
