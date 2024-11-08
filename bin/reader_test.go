package bin_test

import (
	"bytes"
	"testing"

	"github.com/bbfh-dev/go-bin/bin"
	"gotest.tools/assert"
)

func TestReader(test *testing.T) {
	var buffer bytes.Buffer
	assert.NilError(test, testWrite(&buffer))
	var a, b, c, d byte
	var str string
	var x uint16
	var y uint32
	var z uint64
	err := bin.NewReader(&buffer).
		ReadByte(&a).
		ReadByte(&b).
		ReadByte(&c).
		ReadByte(&d).
		ReadString(&str).
		ReadUint16(&x).
		ReadUint32(&y).
		ReadUint64(&z).
		Error()
	assert.NilError(test, err)
	assert.DeepEqual(test, a, byte(0))
	assert.DeepEqual(test, b, byte(1))
	assert.DeepEqual(test, c, byte(2))
	assert.DeepEqual(test, d, byte(3))
	assert.DeepEqual(test, str, "Hello World!")
	assert.DeepEqual(test, x, uint16(69))
	assert.DeepEqual(test, y, uint32(420))
	assert.DeepEqual(test, z, uint64(69420))
}

func TestReaderErr(test *testing.T) {
	var buffer bytes.Buffer
	var a byte
	if bin.NewReader(&buffer).ReadByte(&a).Error() == nil {
		test.Fatal("Must produce an error")
	}
}

func TestReaderPointer(test *testing.T) {
	var buffer bytes.Buffer
	_, err := buffer.WriteString("Hello World!\x00")
	assert.NilError(test, err)

	var str string
	assert.NilError(test, bin.NewReader(&buffer).ReadString(&str).Error())

	assert.DeepEqual(test, buffer.Bytes(), []byte{})
}
