package bin_test

import (
	"bytes"
	"testing"

	"github.com/bbfh-dev/go-bin/bin"
	"gotest.tools/assert"
)

const writerStr = "\x00\x01\x02\x03" + "Hello World!\x00" + "\x00\x45" + "\x00\x00\x01\xa4" + "\x00\x00\x00\x00\x00\x01\x0f\x2c"

func testWrite(buffer *bytes.Buffer) error {
	return bin.NewWriter(buffer).
		Write(0, 1, 2, 3).
		WriteString("Hello World!").
		WriteUint16(69).
		WriteUint32(420).
		WriteUint64(69420).
		Error()
}

func TestWriter(test *testing.T) {
	var buffer bytes.Buffer
	assert.NilError(test, testWrite(&buffer))
	assert.DeepEqual(test, buffer.String(), writerStr)
}
