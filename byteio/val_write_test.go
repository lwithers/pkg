package byteio_test

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"testing"
	"unicode/utf8"

	"github.com/lwithers/pkg/byteio"
)

// TestWriteUint ensures correct operation of the WriteUint variants by
// verifying written data integrity against encoding/binary.
func TestWriteUint(t *testing.T) {
	checkErr := func(err error) {
		if err != nil {
			t.Fatalf("unexpected I/O error: %v", err)
		}
	}

	// prepare a buffer of test data
	vals := []uint64{0, 1, 0x8000, 0xFFFF, 0x81828384858687}
	buf := bytes.NewBuffer(nil)
	for _, val := range vals {
		checkErr(byteio.WriteUint16BE(buf, uint16(val)))
		checkErr(byteio.WriteUint32BE(buf, uint32(val)))
		checkErr(byteio.WriteUint64BE(buf, val))
		checkErr(byteio.WriteUint16LE(buf, uint16(val)))
		checkErr(byteio.WriteUint32LE(buf, uint32(val)))
		checkErr(byteio.WriteUint64LE(buf, val))
	}

	// verify its integrity
	for _, exp64 := range vals {
		var (
			act16, exp16 uint16 = 0, uint16(exp64)
			act32, exp32 uint32 = 0, uint32(exp64)
			act64        uint64
		)

		checkErr(binary.Read(buf, binary.BigEndian, &act16))
		if act16 != exp16 {
			t.Errorf("WriteUint16BE: act %X ≠ exp %X", act16, exp16)
		}

		checkErr(binary.Read(buf, binary.BigEndian, &act32))
		if act32 != exp32 {
			t.Errorf("WriteUint32BE: act %X ≠ exp %X", act32, exp32)
		}

		checkErr(binary.Read(buf, binary.BigEndian, &act64))
		if act64 != exp64 {
			t.Errorf("WriteUint64BE: act %X ≠ exp %X", act64, exp64)
		}

		checkErr(binary.Read(buf, binary.LittleEndian, &act16))
		if act16 != exp16 {
			t.Errorf("WriteUint16LE: act %X ≠ exp %X", act16, exp16)
		}

		checkErr(binary.Read(buf, binary.LittleEndian, &act32))
		if act32 != exp32 {
			t.Errorf("WriteUint32LE: act %X ≠ exp %X", act32, exp32)
		}

		checkErr(binary.Read(buf, binary.LittleEndian, &act64))
		if act64 != exp64 {
			t.Errorf("WriteUint64LE: act %X ≠ exp %X", act64, exp64)
		}
	}
}

// TestWriteInt ensures correct operation of the WriteInt variants by verifying
// written data integrity against encoding/binary.
func TestWriteInt(t *testing.T) {
	checkErr := func(err error) {
		if err != nil {
			t.Fatalf("unexpected I/O error: %v", err)
		}
	}

	// prepare a buffer of test data
	vals := []int64{-1, 0, 1, -0x8000, 0xFFFF, 0x01020304050607}
	buf := bytes.NewBuffer(nil)
	for _, val := range vals {
		checkErr(byteio.WriteInt16BE(buf, int16(val)))
		checkErr(byteio.WriteInt32BE(buf, int32(val)))
		checkErr(byteio.WriteInt64BE(buf, val))
		checkErr(byteio.WriteInt16LE(buf, int16(val)))
		checkErr(byteio.WriteInt32LE(buf, int32(val)))
		checkErr(byteio.WriteInt64LE(buf, val))
	}

	// verify its integrity
	for _, exp64 := range vals {
		var (
			act16, exp16 int16 = 0, int16(exp64)
			act32, exp32 int32 = 0, int32(exp64)
			act64        int64
		)

		checkErr(binary.Read(buf, binary.BigEndian, &act16))
		if act16 != exp16 {
			t.Errorf("WriteInt16BE: act %X ≠ exp %X", act16, exp16)
		}

		checkErr(binary.Read(buf, binary.BigEndian, &act32))
		if act32 != exp32 {
			t.Errorf("WriteInt32BE: act %X ≠ exp %X", act32, exp32)
		}

		checkErr(binary.Read(buf, binary.BigEndian, &act64))
		if act64 != exp64 {
			t.Errorf("WriteInt64BE: act %X ≠ exp %X", act64, exp64)
		}

		checkErr(binary.Read(buf, binary.LittleEndian, &act16))
		if act16 != exp16 {
			t.Errorf("WriteInt16LE: act %X ≠ exp %X", act16, exp16)
		}

		checkErr(binary.Read(buf, binary.LittleEndian, &act32))
		if act32 != exp32 {
			t.Errorf("WriteInt32LE: act %X ≠ exp %X", act32, exp32)
		}

		checkErr(binary.Read(buf, binary.LittleEndian, &act64))
		if act64 != exp64 {
			t.Errorf("WriteInt64LE: act %X ≠ exp %X", act64, exp64)
		}
	}
}

// TestWriteFloat verifies that byteio-written floating point numbers can be
// read back correctly using encoding/binary.
func TestWriteFloat(t *testing.T) {
	vals := []float64{
		0, 1, -1,
		math.MaxFloat32, math.SmallestNonzeroFloat32,
		math.Inf(1), math.Inf(-1), math.NaN(),
	}
	checkErr := func(err error) {
		if err != nil {
			t.Fatalf("unexpected I/O error: %v", err)
		}
	}

	// prepare buffer of encoded data
	buf := bytes.NewBuffer(nil)
	for _, val := range vals {
		checkErr(byteio.WriteFloat32BE(buf, float32(val)))
		checkErr(byteio.WriteFloat32LE(buf, float32(val)))
		checkErr(byteio.WriteFloat64BE(buf, val))
		checkErr(byteio.WriteFloat64LE(buf, val))
	}

	// verify integrity using encoding/binary
	for _, exp64 := range vals {
		var (
			act32, exp32 float32 = 0, float32(exp64)
			act64        float64
		)

		checkErr(binary.Read(buf, binary.BigEndian, &act32))
		if act32 != exp32 && !math.IsNaN(float64(act32)) && !math.IsNaN(exp64) {
			t.Errorf("act %f (0x%X) ≠ exp %f (0x%X)",
				act32, math.Float32bits(act32),
				exp32, math.Float32bits(exp32))
		}
		checkErr(binary.Read(buf, binary.LittleEndian, &act32))
		if act32 != exp32 && !math.IsNaN(float64(act32)) && !math.IsNaN(exp64) {
			t.Errorf("act %f (0x%X) ≠ exp %f (0x%X)",
				act32, math.Float32bits(act32),
				exp32, math.Float32bits(exp32))
		}

		checkErr(binary.Read(buf, binary.BigEndian, &act64))
		if act64 != exp64 && !math.IsNaN(act64) && !math.IsNaN(exp64) {
			t.Errorf("act %f (0x%X) ≠ exp %f (0x%X)",
				act64, math.Float64bits(act64),
				exp64, math.Float64bits(exp64))
		}
		checkErr(binary.Read(buf, binary.LittleEndian, &act64))
		if act64 != exp64 && !math.IsNaN(act64) && !math.IsNaN(exp64) {
			t.Errorf("act %f (0x%X) ≠ exp %f (0x%X)",
				act64, math.Float64bits(act64),
				exp64, math.Float64bits(exp64))
		}
	}
}

// ErrAbortWriter is returned by AbortWriter when the limit is reached.
var ErrAbortWriter = errors.New("aborted write")

// AbortWriter will abort after a certain number of bytes have been written. It
// implements byteio.Writer since we do not want it to be wrapped by a
// bufio.Writer, which would defer errors until a flush operation.
type AbortWriter struct {
	when, cur int
}

func (aw *AbortWriter) WriteByte(b byte) error {
	if aw.when == aw.cur {
		return ErrAbortWriter
	}
	aw.cur++
	return nil
}

func (aw *AbortWriter) WriteRune(r rune) (int, error) {
	len := utf8.RuneLen(r)
	for i := 0; i < len; i++ {
		if err := aw.WriteByte(' '); err != nil {
			return i, err
		}
	}
	return len, nil
}

func (aw *AbortWriter) Write(buf []byte) (int, error) {
	for i, b := range buf {
		if err := aw.WriteByte(b); err != nil {
			return i, err
		}
	}
	return len(buf), nil
}

// TestWriteIntErr ensures that an error is correctly returned when writing an
// integer and the underlying writer reports an error.
func TestWriteIntErr(t *testing.T) {
	check := func(name string, size int, fn func(bout byteio.Writer) error) {
		for i := 0; i <= size; i++ {
			err := fn(&AbortWriter{when: i})
			switch err {
			case nil:
				if i < size {
					t.Errorf("%s: expected err after %d "+
						"bytes", name, i)
				}
			case ErrAbortWriter:
				if i == size {
					t.Errorf("%s: unexpected err after %d "+
						"bytes", name, i)
				}
			default:
				t.Errorf("%s: unexpected err %v", name, err)
			}
		}
	}

	check("WriteUint16BE", 2, func(bout byteio.Writer) error { return byteio.WriteUint16BE(bout, 0) })
	check("WriteUint32BE", 4, func(bout byteio.Writer) error { return byteio.WriteUint32BE(bout, 0) })
	check("WriteUint64BE", 8, func(bout byteio.Writer) error { return byteio.WriteUint64BE(bout, 0) })
	check("WriteUint16LE", 2, func(bout byteio.Writer) error { return byteio.WriteUint16LE(bout, 0) })
	check("WriteUint32LE", 4, func(bout byteio.Writer) error { return byteio.WriteUint32LE(bout, 0) })
	check("WriteUint64LE", 8, func(bout byteio.Writer) error { return byteio.WriteUint64LE(bout, 0) })
}
