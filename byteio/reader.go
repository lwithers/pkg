package byteio

import (
	"bufio"
	"io"
)

// Reader provides byte-oriented reading routines. It is satisfied by
// bufio.Reader, bytes.Reader and bytes.Buffer.
type Reader interface {
	io.Reader
	io.ByteReader
	io.RuneReader
}

// NewReader adapts any io.Reader into a byteio.Reader, possibly returning
// a new bufio.Reader.
func NewReader(in io.Reader) Reader {
	if bin, ok := in.(Reader); ok {
		return bin
	}
	return bufio.NewReader(in)
}
