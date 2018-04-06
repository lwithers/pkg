package byteio_test

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/lwithers/pkg/byteio"
)

// MockReader does not implement Reader, thus ensuring that NewReader must wrap
// it. The implementation returns an incrementing byte pattern.
type MockReader struct {
	pos uint8
}

func (r *MockReader) Read(buf []byte) (n int, err error) {
	for i := range buf {
		buf[i] = r.pos
		r.pos++
	}
	return len(buf), nil
}

// TestNewReaderBytesB checks that a bytes.Buffer is successfully transformed
// into a Reader.
func TestNewReaderBytesB(t *testing.T) {
	orig := bytes.NewBuffer(nil)
	bin := byteio.NewReader(orig)
	if act, ok := bin.(*bytes.Buffer); !ok {
		t.Errorf("Reader(%T) returned unexpected %T", orig, bin)
	} else if act != orig {
		t.Errorf("Reader(%p) returned unexpected %p", orig, act)
	}
}

// TestNewReaderBytesR checks that a bytes.Reader is successfully transformed
// into a Reader.
func TestNewReaderBytesR(t *testing.T) {
	orig := bytes.NewReader(nil)
	bin := byteio.NewReader(orig)
	if act, ok := bin.(*bytes.Reader); !ok {
		t.Errorf("Reader(%T) returned unexpected %T", orig, bin)
	} else if act != orig {
		t.Errorf("Reader(%p) returned unexpected %p", orig, act)
	}
}

// TestNewReaderBufio checks that a bufio.Reader is successfully transformed
// into a Reader.
func TestNewReaderBufio(t *testing.T) {
	orig := bufio.NewReader(os.Stdin)
	bin := byteio.NewReader(orig)
	if act, ok := bin.(*bufio.Reader); !ok {
		t.Errorf("Reader(%T) returned unexpected %T", orig, bin)
	} else if act != orig {
		t.Errorf("Reader(%p) returned unexpected %p", orig, act)
	}
}

// TestNewReader checks that an arbitrary io.Reader is successfully wrapped into
// a byteio.Reader.
func TestNewReader(t *testing.T) {
	orig := new(MockReader)
	bin := byteio.NewReader(orig)
	if _, ok := bin.(*bufio.Reader); !ok {
		t.Errorf("Reader(%T) did not wrap to bufio.Reader (got %T)",
			orig, bin)
	}
}
