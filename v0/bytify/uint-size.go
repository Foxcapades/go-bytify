package bytify


// Uint8StringSize returns the number of characters required to represent the
// given value as a string.
func Uint8StringSize(v uint8) uint8 {
	switch true {
	case v > 99:
		return 3
	case v > 9:
		return 2
	default:
		return 1
	}
}

// Uint16StringSize returns the number of characters required to represent the
// given value as a string.
func Uint16StringSize(v uint16) uint8 {
	switch true {
	case v > 9_999:
		return 5
	case v > 999:
		return 4
	case v > 99:
		return 3
	case v > 9:
		return 2
	default:
		return 1
	}
}

// Uint32StringSize returns the number of characters required to represent the
// given value as a string.
func Uint32StringSize(v uint32) uint8 {
	switch true {
	case v <= maxU16:
		return Uint16StringSize(uint16(v))
	case v > 999_999_999:
		return 10
	case v > 99_999_999:
		return 9
	case v > 9_999_999:
		return 8
	case v > 999_999:
		return 7
	case v > 99_999:
		return 6
	default:
		return 5
	}
}

// Uint64StringSize returns the number of characters required to represent the
// given value as a string.
func Uint64StringSize(v uint64) uint8 {
	switch true {
	case v <= maxU32:
		return Uint32StringSize(uint32(v))
	case v > 9_999_999_999_999_999_999:
		return 20
	case v > 999_999_999_999_999_999:
		return 19
	case v > 99_999_999_999_999_999:
		return 18
	case v > 9_999_999_999_999_999:
		return 17
	case v > 999_999_999_999_999:
		return 16
	case v > 99_999_999_999_999:
		return 15
	case v > 9_999_999_999_999:
		return 14
	case v > 999_999_999_999:
		return 13
	case v > 99_999_999_999:
		return 12
	case v > 9_999_999_999:
		return 11
	default:
		return 10
	}
}
