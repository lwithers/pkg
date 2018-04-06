package byteio

import "math"

// WriteUint16BE writes an unsigned uint16 in big-endian (network) byte order.
func WriteUint16BE(bout Writer, n uint16) error {
	if err := bout.WriteByte(byte(n >> 8)); err != nil {
		return err
	}
	return bout.WriteByte(byte(n))
}

// WriteInt16BE writes a signed int16 in big-endian (network) byte order.
func WriteInt16BE(bout Writer, i int16) error {
	return WriteUint16BE(bout, uint16(i))
}

// WriteUint32BE writes an unsigned uint32 in big-endian (network) byte order.
func WriteUint32BE(bout Writer, n uint32) error {
	if err := bout.WriteByte(byte(n >> 24)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 16)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 8)); err != nil {
		return err
	}
	return bout.WriteByte(byte(n))
}

// WriteInt32BE writes a signed int32 in big-endian byte order.
func WriteInt32BE(bout Writer, i int32) error {
	return WriteUint32BE(bout, uint32(i))
}

// WriteFloat32BE writes an IEEE-754 32-bit floating point number in big-endian
// (network) byte order.
func WriteFloat32BE(bout Writer, f float32) error {
	return WriteUint32BE(bout, math.Float32bits(f))
}

// WriteUint64BE writes an unsigned uint64 in big-endian (network) byte order.
func WriteUint64BE(bout Writer, n uint64) error {
	if err := bout.WriteByte(byte(n >> 56)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 48)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 40)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 32)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 24)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 16)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 8)); err != nil {
		return err
	}
	return bout.WriteByte(byte(n))
}

// WriteUint64BE writes a signed int64 in big-endian (network) byte order.
func WriteInt64BE(bout Writer, i int64) error {
	return WriteUint64BE(bout, uint64(i))
}

// WriteFloat32BE writes an IEEE-754 64-bit floating point number in big-endian
// (network) byte order.
func WriteFloat64BE(bout Writer, f float64) error {
	return WriteUint64BE(bout, math.Float64bits(f))
}

// WriteUint16LE writes an unsigned uint16 in little-endian byte order.
func WriteUint16LE(bout Writer, n uint16) error {
	if err := bout.WriteByte(byte(n)); err != nil {
		return err
	}
	return bout.WriteByte(byte(n >> 8))
}

// WriteInt16LE writes a signed int16 in little-endian byte order.
func WriteInt16LE(bout Writer, i int16) error {
	return WriteUint16LE(bout, uint16(i))
}

// WriteUint32LE writes an unsigned uint32 in little-endian byte order.
func WriteUint32LE(bout Writer, n uint32) error {
	if err := bout.WriteByte(byte(n)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 8)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 16)); err != nil {
		return err
	}
	return bout.WriteByte(byte(n >> 24))
}

// WriteInt32LE writes a signed int32 in little-endian byte order.
func WriteInt32LE(bout Writer, i int32) error {
	return WriteUint32LE(bout, uint32(i))
}

// WriteFloat32LE writes an IEEE-754 32-bit floating point number in
// little-endian byte order.
func WriteFloat32LE(bout Writer, f float32) error {
	return WriteUint32LE(bout, math.Float32bits(f))
}

// WriteUint64LE writes an unsigned uint64 in little-endian byte order.
func WriteUint64LE(bout Writer, n uint64) error {
	if err := bout.WriteByte(byte(n)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 8)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 16)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 24)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 32)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 40)); err != nil {
		return err
	}
	if err := bout.WriteByte(byte(n >> 48)); err != nil {
		return err
	}
	return bout.WriteByte(byte(n >> 56))
}

// WriteInt64LE writes a signed int64 in little-endian byte order.
func WriteInt64LE(bout Writer, i int64) error {
	return WriteUint64LE(bout, uint64(i))
}

// WriteFloat64LE writes an IEEE-754 64-bit floating point number in
// little-endian byte order.
func WriteFloat64LE(bout Writer, f float64) error {
	return WriteUint64LE(bout, math.Float64bits(f))
}
