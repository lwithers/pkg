package byteio_test

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math"
	"testing"

	"github.com/lwithers/pkg/byteio"
)

// TestReadUint ensures correct operation of the ReadUint variants by verifying
// read data integrity against encoding/binary.
func TestReadUint(t *testing.T) {
	checkErr := func(err error) {
		if err != nil {
			t.Fatalf("unexpected I/O error: %v", err)
		}
	}

	// prepare a buffer of test data
	vals := []uint64{0, 1, 0x8000, 0xFFFF, 0x81020304050607}
	buf := bytes.NewBuffer(nil)
	for _, val := range vals {
		checkErr(binary.Write(buf, binary.BigEndian, uint16(val)))
		checkErr(binary.Write(buf, binary.BigEndian, uint32(val)))
		checkErr(binary.Write(buf, binary.BigEndian, uint64(val)))
		checkErr(binary.Write(buf, binary.LittleEndian, uint16(val)))
		checkErr(binary.Write(buf, binary.LittleEndian, uint32(val)))
		checkErr(binary.Write(buf, binary.LittleEndian, uint64(val)))
	}

	// validate that we read back the expected values
	for _, exp64 := range vals {
		var (
			act16, exp16 uint16 = 0, uint16(exp64)
			act32, exp32 uint32 = 0, uint32(exp64)
			act64        uint64
			err          error
		)

		act16, err = byteio.ReadUint16BE(buf)
		checkErr(err)
		if act16 != exp16 {
			t.Errorf("ReadUint16BE: act %X ≠ exp %X", act16, exp16)
		}

		act32, err = byteio.ReadUint32BE(buf)
		checkErr(err)
		if act32 != exp32 {
			t.Errorf("ReadUint32BE: act %X ≠ exp %X", act32, exp32)
		}

		act64, err = byteio.ReadUint64BE(buf)
		checkErr(err)
		if act64 != exp64 {
			t.Errorf("ReadUint64BE: act %X ≠ exp %X", act64, exp64)
		}

		act16, err = byteio.ReadUint16LE(buf)
		checkErr(err)
		if act16 != exp16 {
			t.Errorf("ReadUint16LE: act %X ≠ exp %X", act16, exp16)
		}

		act32, err = byteio.ReadUint32LE(buf)
		checkErr(err)
		if act32 != exp32 {
			t.Errorf("ReadUint32LE: act %X ≠ exp %X", act32, exp32)
		}

		act64, err = byteio.ReadUint64LE(buf)
		checkErr(err)
		if act64 != exp64 {
			t.Errorf("ReadUint64LE: act %X ≠ exp %X", act64, exp64)
		}
	}
}

// TestReadInt ensures correct operation of the ReadInt variants by comparing
// operation against encoding/binary.
func TestReadInt(t *testing.T) {
	checkErr := func(err error) {
		if err != nil {
			t.Fatalf("read error: %v", err)
		}
	}

	// prepare a buffer of test data using encoding/binary
	vals := []int64{-1, 0, 1, -0x8000, 0xFFFF, 0x01020304050607}
	buf := bytes.NewBuffer(nil)
	for _, val := range vals {
		checkErr(binary.Write(buf, binary.BigEndian, int16(val)))
		checkErr(binary.Write(buf, binary.BigEndian, int32(val)))
		checkErr(binary.Write(buf, binary.BigEndian, int64(val)))
		checkErr(binary.Write(buf, binary.LittleEndian, int16(val)))
		checkErr(binary.Write(buf, binary.LittleEndian, int32(val)))
		checkErr(binary.Write(buf, binary.LittleEndian, int64(val)))
	}

	// validate that we read back the expected values
	for _, exp64 := range vals {
		var (
			v16, exp16 int16 = 0, int16(exp64)
			v32, exp32 int32 = 0, int32(exp64)
			v64        int64
			err        error
		)

		v16, err = byteio.ReadInt16BE(buf)
		checkErr(err)
		if v16 != exp16 {
			t.Errorf("ReadInt16BE: act %X ≠ exp %X", v16, exp16)
		}

		v32, err = byteio.ReadInt32BE(buf)
		checkErr(err)
		if v32 != exp32 {
			t.Errorf("ReadInt32BE: act %X ≠ exp %X", v32, exp32)
		}

		v64, err = byteio.ReadInt64BE(buf)
		checkErr(err)
		if v64 != exp64 {
			t.Errorf("ReadInt64BE: act %X ≠ exp %X", v64, exp64)
		}

		v16, err = byteio.ReadInt16LE(buf)
		checkErr(err)
		if v16 != exp16 {
			t.Errorf("ReadInt16LE: act %X ≠ exp %X", v16, exp16)
		}

		v32, err = byteio.ReadInt32LE(buf)
		checkErr(err)
		if v32 != exp32 {
			t.Errorf("ReadInt32LE: act %X ≠ exp %X", v32, exp32)
		}

		v64, err = byteio.ReadInt64LE(buf)
		checkErr(err)
		if v64 != exp64 {
			t.Errorf("ReadInt64LE: act %X ≠ exp %X", v64, exp64)
		}
	}
}

// TestReadFloat validates operation of the byteio.ReadFloat variants against
// values written by encoding/binary.
func TestReadFloat(t *testing.T) {
	vals := []float64{
		0, 1, -1,
		math.MaxFloat32, math.SmallestNonzeroFloat32,
		math.Inf(1), math.Inf(-1), math.NaN(),
	}

	buf := bytes.NewBuffer(nil)

	for _, val := range vals {
		binary.Write(buf, binary.BigEndian, val)
		binary.Write(buf, binary.BigEndian, float32(val))
		binary.Write(buf, binary.LittleEndian, val)
		binary.Write(buf, binary.LittleEndian, float32(val))
	}

	for _, exp64 := range vals {
		exp32 := float32(exp64)

		if act64, err := byteio.ReadFloat64BE(buf); err != nil {
			t.Fatalf("read error: %v", err)
		} else if act64 != exp64 && !math.IsNaN(act64) && !math.IsNaN(exp64) {
			t.Errorf("act %f (0x%X) ≠ exp %f (0x%X)",
				act64, math.Float64bits(act64),
				exp64, math.Float64bits(exp64))
		}
		if act32, err := byteio.ReadFloat32BE(buf); err != nil {
			t.Fatalf("read error: %v", err)
		} else if act32 != exp32 && !math.IsNaN(float64(act32)) && !math.IsNaN(exp64) {
			t.Errorf("act %f (0x%X) ≠ exp %f (0x%X)",
				act32, math.Float32bits(act32),
				exp32, math.Float32bits(exp32))
		}

		if act64, err := byteio.ReadFloat64LE(buf); err != nil {
			t.Fatalf("read error: %v", err)
		} else if act64 != exp64 && !math.IsNaN(act64) && !math.IsNaN(exp64) {
			t.Errorf("act %f (0x%X) ≠ exp %f (0x%X)",
				act64, math.Float64bits(act64),
				exp64, math.Float64bits(exp64))
		}
		if act32, err := byteio.ReadFloat32LE(buf); err != nil {
			t.Fatalf("read error: %v", err)
		} else if act32 != exp32 && !math.IsNaN(float64(act32)) && !math.IsNaN(exp64) {
			t.Errorf("act %f (0x%X) ≠ exp %f (0x%X)",
				act32, math.Float32bits(act32),
				exp32, math.Float32bits(exp32))
		}
	}
}

// TestReadIntEOF ensures that io.EOF is propagated correctly when a ReadInt
// variant is called at the end of a buffer (i.e. 0 bytes to read). It only
// tests the Uint variants since the Int and Float variants are wrappers over
// Uint.
func TestReadIntEOF(t *testing.T) {
	bin := bytes.NewBuffer(nil)
	check := func(fn string, f func() error) {
		err := f()
		switch err {
		case nil:
			t.Errorf("%s: did not return expected error", fn)
		case io.EOF:
			// OK
		default:
			t.Errorf("%s: returned unexpected error %v", fn, err)
		}
	}

	check("ReadUint16LE", func() error { _, err := byteio.ReadUint16LE(bin); return err })
	check("ReadUint16BE", func() error { _, err := byteio.ReadUint16BE(bin); return err })
	check("ReadUint32LE", func() error { _, err := byteio.ReadUint32LE(bin); return err })
	check("ReadUint32BE", func() error { _, err := byteio.ReadUint32BE(bin); return err })
	check("ReadUint64LE", func() error { _, err := byteio.ReadUint64LE(bin); return err })
	check("ReadUint64BE", func() error { _, err := byteio.ReadUint64BE(bin); return err })

}

// TestReadIntShort ensures that io.ErrUnexpectedEOF is returned when a ReadInt
// variant is called without sufficient data, but not immediately at the end of
// the buffer (i.e. > 0 bytes to read, but not sufficient for data type).
func TestReadIntShort(t *testing.T) {
	check := func(fn string, sz int, f func(bin byteio.Reader) error) {
		for i := 1; i < sz; i++ {
			bin := bytes.NewBuffer(make([]byte, i))
			err := f(bin)
			switch err {
			case nil:
				t.Errorf("%s/%d: did not return expected error",
					fn, i)
			case io.ErrUnexpectedEOF:
				// OK
			case io.EOF:
				t.Errorf("%s/%d: returned incorrect io.EOF",
					fn, i)
			default:
				t.Errorf("%s/%d: returned unexpected error %v",
					fn, i, err)
			}
		}
	}

	check("ReadUint16LE", 2, func(bin byteio.Reader) error { _, err := byteio.ReadUint16LE(bin); return err })
	check("ReadUint16BE", 2, func(bin byteio.Reader) error { _, err := byteio.ReadUint16BE(bin); return err })
	check("ReadUint32LE", 4, func(bin byteio.Reader) error { _, err := byteio.ReadUint32LE(bin); return err })
	check("ReadUint32BE", 4, func(bin byteio.Reader) error { _, err := byteio.ReadUint32BE(bin); return err })
	check("ReadUint64LE", 8, func(bin byteio.Reader) error { _, err := byteio.ReadUint64LE(bin); return err })
	check("ReadUint64BE", 8, func(bin byteio.Reader) error { _, err := byteio.ReadUint64BE(bin); return err })
}

// ErrAbortReader is returned by AbortReader when the limit is reached.
var ErrAbortReader = errors.New("aborted read")

// AbortReader will abort after a certain number of bytes have been read with a
// non-io.EOF error. This can be used to test error handling on partial reads.
// It does not implement byteio.Reader, to ensure that errors propagate
// correctly through the implicit bufio wrapper.
type AbortReader struct {
	when, cur int
}

func (ar *AbortReader) Read(buf []byte) (n int, err error) {
	for i := range buf {
		if ar.when == ar.cur {
			return i, ErrAbortReader
		}
		buf[i] = 0x55
		ar.cur++
	}
	return len(buf), nil
}

// TestReadIntErr ensures that errors encountered during reads are propagated
// correctly.
func TestReadIntErr(t *testing.T) {
	check := func(name string, size int, fn func(bin byteio.Reader) error) {
		for i := 0; i < size; i++ {
			err := fn(byteio.NewReader(&AbortReader{when: i}))
			switch err {
			case nil:
				if i < size {
					t.Errorf("%s: expected err after %d "+
						"bytes", name, i)
				}
			case ErrAbortReader:
				if i == size {
					t.Errorf("%s: unexpected err after %d "+
						"bytes", name, i)
				}
			default:
				t.Errorf("%s: unexpected err %v", name, err)
			}
		}
	}

	check("ReadUint16BE", 2, func(bin byteio.Reader) error { _, err := byteio.ReadUint16BE(bin); return err })
	check("ReadUint32BE", 4, func(bin byteio.Reader) error { _, err := byteio.ReadUint32BE(bin); return err })
	check("ReadUint64BE", 8, func(bin byteio.Reader) error { _, err := byteio.ReadUint64BE(bin); return err })
	check("ReadUint16LE", 2, func(bin byteio.Reader) error { _, err := byteio.ReadUint16LE(bin); return err })
	check("ReadUint32LE", 4, func(bin byteio.Reader) error { _, err := byteio.ReadUint32LE(bin); return err })
	check("ReadUint64LE", 8, func(bin byteio.Reader) error { _, err := byteio.ReadUint64LE(bin); return err })
}

// BenchmarkReadUint32BE is a simple benchmark for reading 32-bit integers.
func BenchmarkReadUint32BE(b *testing.B) {
	bin := byteio.NewReader(new(MockReader))
	for i := 0; i < b.N; i++ {
		_, _ = byteio.ReadUint32BE(bin)
	}
}

// BenchmarkReadUint32BE is a simple benchmark for reading 64-bit integers.
func BenchmarkReadUint64BE(b *testing.B) {
	bin := byteio.NewReader(new(MockReader))
	for i := 0; i < b.N; i++ {
		_, _ = byteio.ReadUint64BE(bin)
	}
}
