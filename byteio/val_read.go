package byteio

import (
	"io"
	"math"
)

// ReadUint16BE reads an unsigned uint16 in big-endian (network) byte order.
func ReadUint16BE(bin Reader) (uint16, error) {
	var b0, b1 byte
	var err error
	if b0, err = bin.ReadByte(); err != nil {
		// allow io.EOF to propagate normally before first read
		return 0, err
	}
	if b1, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	return uint16(b0)<<8 | uint16(b1), nil
}

// ReadInt16BE reads a signed int16 in big-endian (network) byte order.
func ReadInt16BE(bin Reader) (int16, error) {
	n, err := ReadUint16BE(bin)
	return int16(n), err
}

// ReadUint32BE reads an unsigned uint32 in big-endian (network) byte order.
func ReadUint32BE(bin Reader) (uint32, error) {
	var b0, b1, b2, b3 byte
	var err error
	if b0, err = bin.ReadByte(); err != nil {
		return 0, err
	}
	if b1, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b2, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b3, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	return uint32(b0)<<24 | uint32(b1)<<16 |
		uint32(b2)<<8 | uint32(b3), nil
}

// ReadInt32BE reads a signed int32 in big-endian (network) byte order.
func ReadInt32BE(bin Reader) (int32, error) {
	n, err := ReadUint32BE(bin)
	return int32(n), err
}

// ReadFloat32BE reads an IEEE-754 32-bit floating point number in big-endian
// (network) byte order.
func ReadFloat32BE(bin Reader) (float32, error) {
	n, err := ReadUint32BE(bin)
	return math.Float32frombits(n), err
}

// ReadUint64BE reads an unsigned uint64 in big-endian (network) byte order.
func ReadUint64BE(bin Reader) (uint64, error) {
	var b0, b1, b2, b3, b4, b5, b6, b7 byte
	var err error
	if b0, err = bin.ReadByte(); err != nil {
		return 0, err
	}
	if b1, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b2, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b3, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b4, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b5, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b6, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b7, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	return uint64(b0)<<56 | uint64(b1)<<48 |
		uint64(b2)<<40 | uint64(b3)<<32 |
		uint64(b4)<<24 | uint64(b5)<<16 |
		uint64(b6)<<8 | uint64(b7), nil
}

// ReadInt64BE reads a signed int64 in big-endian (network) byte order.
func ReadInt64BE(bin Reader) (int64, error) {
	n, err := ReadUint64BE(bin)
	return int64(n), err
}

// ReadFloat64BE reads an IEEE-754 64-bit floating point number in big-endian
// (network) byte order.
func ReadFloat64BE(bin Reader) (float64, error) {
	n, err := ReadUint64BE(bin)
	return math.Float64frombits(n), err
}

// ReadUint16LE reads an unsigned uint16 in little-endian byte order.
func ReadUint16LE(bin Reader) (uint16, error) {
	var b0, b1 byte
	var err error
	if b1, err = bin.ReadByte(); err != nil {
		return 0, err
	}
	if b0, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	return uint16(b0)<<8 | uint16(b1), nil
}

// ReadUint16LE reads a signed int16 in little-endian byte order.
func ReadInt16LE(bin Reader) (int16, error) {
	n, err := ReadUint16LE(bin)
	return int16(n), err
}

// ReadUint32LE reads an unsigned uint32 in little-endian byte order.
func ReadUint32LE(bin Reader) (uint32, error) {
	var b0, b1, b2, b3 byte
	var err error
	if b3, err = bin.ReadByte(); err != nil {
		return 0, err
	}
	if b2, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b1, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b0, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	return uint32(b0)<<24 | uint32(b1)<<16 |
		uint32(b2)<<8 | uint32(b3), nil
}

// ReadUint32LE reads a signed int32 in little-endian byte order.
func ReadInt32LE(bin Reader) (int32, error) {
	n, err := ReadUint32LE(bin)
	return int32(n), err
}

// ReadFloat32LE reads an IEEE-754 32-bit floating point number in big-endian
// (network) byte order.
func ReadFloat32LE(bin Reader) (float32, error) {
	n, err := ReadUint32LE(bin)
	return math.Float32frombits(n), err
}

// ReadUint64LE reads an unsigned uint64 in little-endian byte order.
func ReadUint64LE(bin Reader) (uint64, error) {
	var b0, b1, b2, b3, b4, b5, b6, b7 byte
	var err error
	if b7, err = bin.ReadByte(); err != nil {
		return 0, err
	}
	if b6, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b5, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b4, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b3, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b2, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b1, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	if b0, err = bin.ReadByte(); err != nil {
		if err == io.EOF {
			return 0, io.ErrUnexpectedEOF
		}
		return 0, err
	}
	return uint64(b0)<<56 | uint64(b1)<<48 |
		uint64(b2)<<40 | uint64(b3)<<32 |
		uint64(b4)<<24 | uint64(b5)<<16 |
		uint64(b6)<<8 | uint64(b7), nil
}

// ReadInt64LE reads a signed int64 in little-endian byte order.
func ReadInt64LE(bin Reader) (int64, error) {
	n, err := ReadUint64LE(bin)
	return int64(n), err
}

// ReadFloat64LE reads an IEEE-754 64-bit floating point number in big-endian
// (network) byte order.
func ReadFloat64LE(bin Reader) (float64, error) {
	n, err := ReadUint64LE(bin)
	return math.Float64frombits(n), err
}
