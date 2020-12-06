package bytify

// Int8ToBytes populates the given byte slice with the string representation of
// the given int8 value.
//
// WARNING: This method makes no attempt to verify that the input slice is not
// nil or even if it has the capacity to contain the given value.
//
// Returns the number of bytes written to the given buffer.
func Int8ToBytes(v int8, buf []byte) (out uint8) {
	var tmp uint8
	out = Int8StringSize(v)
	pos := int8(out) - 1
	min := int8(-1)

	if v < 0 {
		tmp = uint8(^v + 1)
		buf[0] = '-'
		min++
	} else {
		tmp = uint8(v)
	}

	for ; pos > min; pos-- {
		buf[pos] = tmp%base + '0'
		tmp /= base
	}

	return
}

// Int8ToByteSlice returns a new byte slice with the string representation of
// the given uint8 value.
func Int8ToByteSlice(v int8) (buf []byte) {
	sz := Int8StringSize(v)
	buf = make([]byte, sz)

	var tmp uint8
	pos := int8(sz) - 1
	min := int8(-1)

	if v < 0 {
		tmp = uint8(^v + 1)
		buf[0] = '-'
		min++
	} else {
		tmp = uint8(v)
	}

	for ; pos > min; pos-- {
		buf[pos] = tmp%base + '0'
		tmp /= base
	}

	return
}

// Int8ToBuf writes the the string representation of the given int8 value to
// the given writer.
//
// Returns the number of bytes written to the given buffer.
func Int8ToBuf(v int8, buf Writer) uint8 {
	tmp := Int8ToByteSlice(v)

	_, _ = buf.Write(tmp)

	return uint8(len(tmp))
}

// Int16ToBytes populates the given byte slice with the string representation of
// the given int16 value.
//
// WARNING: This method makes no attempt to verify that the input slice is not
// nil or even if it has the capacity to contain the given value.
//
// Returns the number of bytes written to the given buffer.
func Int16ToBytes(v int16, buf []byte) (out uint8) {
	var tmp uint16
	out = Int16StringSize(v)
	pos := int8(out) - 1
	min := int8(-1)

	if v < 0 {
		tmp = uint16(^v + 1)
		buf[0] = '-'
		min++
	} else {
		tmp = uint16(v)
	}

	for ; pos > min; pos-- {
		buf[pos] = byte(tmp%base) + '0'
		tmp /= base
	}

	return
}

// Int16ToByteSlice returns a new byte slice with the string representation of
// the given uint16 value.
func Int16ToByteSlice(v int16) (buf []byte) {
	sz := Int16StringSize(v)
	buf = make([]byte, sz)

	var tmp uint16
	pos := int8(sz) - 1
	min := int8(-1)

	if v < 0 {
		tmp = uint16(^v + 1)
		buf[0] = '-'
		min++
	} else {
		tmp = uint16(v)
	}

	for ; pos > min; pos-- {
		buf[pos] = byte(tmp%base) + '0'
		tmp /= base
	}

	return
}

// Int16ToBuf writes the the string representation of the given int16 value to
// the given writer.
//
// Returns the number of bytes written to the given buffer.
func Int16ToBuf(v int16, buf Writer) uint8 {
	tmp := Int16ToByteSlice(v)

	_, _ = buf.Write(tmp)

	return uint8(len(tmp))
}

// Int32ToBytes populates the given byte slice with the string representation of
// the given int32 value.
//
// WARNING: This method makes no attempt to verify that the input slice is not
// nil or even if it has the capacity to contain the given value.
//
// Returns the number of bytes written to the given buffer.
func Int32ToBytes(v int32, buf []byte) (out uint8) {
	var tmp uint32
	out = Int32StringSize(v)
	pos := int8(out) - 1
	min := int8(-1)

	if v < 0 {
		tmp = uint32(^v + 1)
		buf[0] = '-'
		min++
	} else {
		tmp = uint32(v)
	}

	for ; pos > min; pos-- {
		buf[pos] = byte(tmp%base) + '0'
		tmp /= base
	}

	return
}

// Int32ToByteSlice returns a new byte slice with the string representation of
// the given uint32 value.
func Int32ToByteSlice(v int32) (buf []byte) {
	sz := Int32StringSize(v)
	buf = make([]byte, sz)

	var tmp uint32
	pos := int8(sz) - 1
	min := int8(-1)

	if v < 0 {
		tmp = uint32(^v + 1)
		buf[0] = '-'
		min++
	} else {
		tmp = uint32(v)
	}

	for ; pos > min; pos-- {
		buf[pos] = byte(tmp%base) + '0'
		tmp /= base
	}

	return
}

// Int32ToBuf writes the the string representation of the given int32 value to
// the given writer.
//
// Returns the number of bytes written to the given buffer.
func Int32ToBuf(v int32, buf Writer) uint8 {
	tmp := Int32ToByteSlice(v)

	_, _ = buf.Write(tmp)

	return uint8(len(tmp))
}

// Int64ToBytes populates the given byte slice with the string representation of
// the given int64 value.
//
// WARNING 1: This method makes no attempt to verify that the input slice is not
// nil or even if it has the capacity to contain the given value.
//
// WARNING 2: This method makes no attempt to verify that the input value would
// not overflow it's type, iE
//
// Returns the number of bytes written to the given buffer.
func Int64ToBytes(v int64, buf []byte) (out uint8) {
	var tmp uint64
	out = Int64StringSize(v)
	pos := int8(out) - 1
	min := int8(-1)

	if v < 0 {
		tmp = uint64(^v + 1)
		buf[0] = '-'
		min++
	} else {
		tmp = uint64(v)
	}

	for ; pos > min; pos-- {
		buf[pos] = byte(tmp%base) + '0'
		tmp /= base
	}

	return
}

// Int64ToByteSlice returns a new byte slice with the string representation of
// the given uint64 value.
func Int64ToByteSlice(v int64) (buf []byte) {
	sz := Int64StringSize(v)
	buf = make([]byte, sz)

	var tmp uint64
	pos := int8(sz) - 1
	min := int8(-1)

	if v < 0 {
		tmp = uint64(^v + 1)
		buf[0] = '-'
		min++
	} else {
		tmp = uint64(v)
	}

	for ; pos > min; pos-- {
		buf[pos] = byte(tmp%base) + '0'
		tmp /= base
	}

	return
}

// Int64ToBuf writes the the string representation of the given int64 value to
// the given writer.
//
// Returns the number of bytes written to the given buffer.
func Int64ToBuf(v int64, buf Writer) uint8 {
	tmp := Int64ToByteSlice(v)

	_, _ = buf.Write(tmp)

	return uint8(len(tmp))
}
