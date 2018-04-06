package byteio_test

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/lwithers/pkg/byteio"
)

// MockWriter discards data sent to it. It also implements flusher.
type MockWriter struct {
	sawFlush bool
}

func (w *MockWriter) Write(buf []byte) (n int, err error) {
	w.sawFlush = false
	return len(buf), nil
}

func (w *MockWriter) Flush() error {
	w.sawFlush = true
	return nil
}

// TestNewWriterBytesB checks that a bytes.Buffer is successfully transformed
// into a Writer.
func TestNewWriterBytesB(t *testing.T) {
	orig := bytes.NewBuffer(nil)
	bin := byteio.NewWriter(orig)
	if act, ok := bin.(*bytes.Buffer); !ok {
		t.Errorf("Writer(%T) returned unexpected %T", orig, bin)
	} else if act != orig {
		t.Errorf("Writer(%p) returned unexpected %p", orig, act)
	}
}

// TestNewWriterBufio checks that a bufio.Writer is successfully transformed
// into a Writer.
func TestNewWriterBufio(t *testing.T) {
	orig := bufio.NewWriter(os.Stdin)
	bin := byteio.NewWriter(orig)
	if act, ok := bin.(*bufio.Writer); !ok {
		t.Errorf("Writer(%T) returned unexpected %T", orig, bin)
	} else if act != orig {
		t.Errorf("Writer(%p) returned unexpected %p", orig, act)
	}
}

// TestNewWriter checks that an arbitrary io.Writer is successfully wrapped into
// a byteio.Writer.
func TestNewWriter(t *testing.T) {
	orig := new(MockWriter)
	bin := byteio.NewWriter(orig)
	if _, ok := bin.(*bufio.Writer); !ok {
		t.Errorf("Writer(%T) did not wrap to bufio.Writer (got %T)",
			orig, bin)
	}
}

// TestFlushIfNecessary checks that no error is returned when called on
// something which does not require flush, and that Flush() is indeed called
// if it is presnet.
func TestFlushIfNecessary(t *testing.T) {
	noflush := bytes.NewBuffer(nil)
	if err := byteio.FlushIfNecessary(noflush); err != nil {
		t.Errorf("unexpected error while not flushing: %v", err)
	}

	w := new(MockWriter)
	if err := byteio.FlushIfNecessary(w); err != nil {
		t.Errorf("unexpected error while flushing: %v", err)
	}
	if !w.sawFlush {
		t.Error("Flush() was not called")
	}
}
