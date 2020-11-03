package bytify

const (
	base = 10

	maxU16 uint32 = 0x0000_FFFF
	maxU32 uint64 = 0x0000_0000_FFFF_FFFF
)

// Uint8ToBytes populates the given byte slice with the string representation of
// the given uint8 value.
//
// WARNING: This method makes no attempt to verify that the input slice is not
// nil or even if it has the capacity to contain the given value.
//
// Returns the number of bytes written to the given buffer.
func Uint8ToBytes(v uint8, buf []byte) uint8 {
	pos := uint8StartPos(v)
	cur := v

	for ; pos > -1; pos-- {
		buf[pos] = cur%base + '0'
		cur /= base
	}

	return uint8(pos + 1)
}

// Uint8ToBuf writes the the string representation of the given uint8 value to
// the given writer.
//
// Returns the number of bytes written to the given buffer.
func Uint8ToBuf(v uint8, buf Writer) uint8 {
	sz := Uint8StringSize(v)
	tmp := make([]byte, sz)
	cur := v

	for pos := int8(sz) - 1; pos > -1; pos-- {
		tmp[pos] = cur%base + '0'
		cur /= base
	}

	_, _ = buf.Write(tmp)

	return sz
}

// Uint16ToBytes populates the given byte slice with the string representation
// of the given uint16 value.
//
// WARNING: This method makes no attempt to verify that the input slice is not
// nil or even if it has the capacity to contain the given value.
//
// Returns the number of bytes written to the given buffer.
func Uint16ToBytes(v uint16, buf []byte) uint8 {
	pos := uint16StartPos(v)
	cur := v

	for ; pos > -1; pos-- {
		buf[pos] = byte(cur%base) + '0'
		cur /= base
	}

	return uint8(pos + 1)
}

// Uint16ToBuf writes the the string representation of the given uint16 value to
// the given writer.
//
// Returns the number of bytes written to the given buffer.
func Uint16ToBuf(v uint16, buf Writer) uint8 {
	sz := Uint16StringSize(v)
	tmp := make([]byte, sz)

	cur := v

	for pos := int8(sz) - 1; pos > -1; pos-- {
		tmp[pos] = byte(cur%base) + '0'
		cur /= base
	}

	_, _ = buf.Write(tmp)

	return sz
}

// Uint32ToBytes populates the given byte slice with the string representation
// of the given uint32 value.
//
// WARNING: This method makes no attempt to verify that the input slice is not
// nil or even if it has the capacity to contain the given value.
//
// Returns the number of bytes written to the given buffer. // TODO: given 66000 it needed 6 spaces?
func Uint32ToBytes(v uint32, buf []byte) uint8 {
	pos := uint32StartPos(v)
	cur := v

	for ; pos > -1; pos-- {
		buf[pos] = byte(cur%base) + '0'
		cur /= base
	}

	return uint8(pos + 1)
}

// Uint32ToBuf writes the the string representation of the given uint32 value to
// the given writer.
//
// Returns the number of bytes written to the given buffer.
func Uint32ToBuf(v uint32, buf Writer) uint8 {
	sz := Uint32StringSize(v)
	tmp := make([]byte, sz)
	cur := v

	for pos := int8(sz) - 1; pos > -1; pos-- {
		tmp[pos] = byte(cur%base) + '0'
		cur /= base
	}

	_, _ = buf.Write(tmp)

	return sz
}

// Uint64ToBytes populates the given byte slice with the string representation
// of the given uint64 value.
//
// WARNING: This method makes no attempt to verify that the input slice is not
// nil or even if it has the capacity to contain the given value.
//
// Returns the number of bytes written to the given buffer.
func Uint64ToBytes(v uint64, buf []byte) uint8 {
	pos := uint64StartPos(v)
	cur := v

	for ; pos > -1; pos-- {
		buf[pos] = byte(cur%base) + '0'
		cur /= base
	}

	return uint8(pos + 1)
}

// Uint64ToBuf writes the the string representation of the given uint32 value to
// the given writer.
//
// Returns the number of bytes written to the given buffer.
func Uint64ToBuf(v uint64, buf Writer) uint8 {
	sz := Uint64StringSize(v)
	tmp := make([]byte, sz)
	cur := v

	for pos := int8(sz) - 1; pos > -1; pos-- {
		tmp[pos] = byte(cur%base) + '0'
		cur /= base
	}

	_, _ = buf.Write(tmp)

	return sz
}

func uint8StartPos(v uint8) int8 {
	return int8(Uint8StringSize(v)) - 1
}

func uint16StartPos(v uint16) int8 {
	return int8(Uint16StringSize(v)) - 1
}

func uint32StartPos(v uint32) int8 {
	return int8(Uint32StringSize(v)) - 1
}

func uint64StartPos(v uint64) int8 {
	return int8(Uint64StringSize(v)) - 1
}
