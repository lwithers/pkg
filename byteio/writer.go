package byteio

import (
	"bufio"
	"io"
)

// Writer provides byte-oriented writing routines. It is satisfied by
// bufio.Writer and bytes.Buffer.
type Writer interface {
	io.Writer
	io.ByteWriter
	WriteRune(r rune) (n int, err error)
}

// NewWriter adapts any io.Writer into a byteio.Writer, possibly returning
// a new bufio.Writer.
func NewWriter(out io.Writer) Writer {
	if bout, ok := out.(Writer); ok {
		return bout
	}
	return bufio.NewWriter(out)
}

type flusher interface {
	Flush() error
}

// FlushIfNecessary tests whether out has a Flush method and if so calls it. If
// out does not have a Flush method this function silently does nothing.
func FlushIfNecessary(out io.Writer) error {
	if fout, ok := out.(flusher); ok {
		return fout.Flush()
	}
	return nil
}
