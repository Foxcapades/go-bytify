package bytify

// Int8StringSize returns the number of characters required to represent the
// given value as a string.
func Int8StringSize(v int8) uint8 {
	if v < 0 {
		return nInt8StringSize(v)
	}

	return pInt8StringSize(v)
}

func nInt8StringSize(v int8) uint8 {
	switch true {
	case v < -99:
		return 4
	case v < -9:
		return 3
	default:
		return 2
	}
}

func pInt8StringSize(v int8) uint8 {
	switch true {
	case v > 99:
		return 3
	case v > 9:
		return 2
	default:
		return 1
	}
}

////////////////////////////////////////////////////////////////////////////////


// Int16StringSize returns the number of characters required to represent the
// given value as a string.
func Int16StringSize(v int16) uint8 {
	if v < 0 {
		return nInt16StringSize(v)
	}

	return pInt16StringSize(v)
}

func nInt16StringSize(v int16) uint8 {
	switch true {
	case v < -9_999:
		return 6
	case v < -999:
		return 5
	case v < -99:
		return 4
	case v < -9:
		return 3
	default:
		return 2
	}
}

func pInt16StringSize(v int16) uint8 {
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

////////////////////////////////////////////////////////////////////////////////

// Int32StringSize returns the number of characters required to represent the
// given value as a string.
func Int32StringSize(v int32) uint8 {
	if v < 0 {
		return nInt32StringSize(v)
	}

	return pInt32StringSize(v)
}

func nInt32StringSize(v int32) uint8 {
	switch true {
	case v < -999_999_999:
		return 11
	case v < -99_999_999:
		return 10
	case v < -9_999_999:
		return 9
	case v < -999_999:
		return 8
	case v < -99_999:
		return 7
	case v < -9_999:
		return 6
	case v < -999:
		return 5
	case v < -99:
		return 4
	case v < -9:
		return 3
	default:
		return 2
	}
}

func pInt32StringSize(v int32) uint8 {
	switch true {
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

////////////////////////////////////////////////////////////////////////////////

// Int64StringSize returns the number of characters required to represent the
// given value as a string.
func Int64StringSize(v int64) uint8 {
	if v < 0 {
		return nInt64StringSize(v)
	}

	return pInt64StringSize(v)
}

func nInt64StringSize(v int64) uint8 {
	switch true {
	case v < -999_999_999_999_999_999:
		return 20
	case v < -99_999_999_999_999_999:
		return 19
	case v < -9_999_999_999_999_999:
		return 18
	case v < -999_999_999_999_999:
		return 17
	case v < -99_999_999_999_999:
		return 16
	case v < -9_999_999_999_999:
		return 15
	case v < -999_999_999_999:
		return 14
	case v < -99_999_999_999:
		return 13
	case v < -9_999_999_999:
		return 12
	case v < -999_999_999:
		return 11
	case v < -99_999_999:
		return 10
	case v < -9_999_999:
		return 9
	case v < -999_999:
		return 8
	case v < -99_999:
		return 7
	case v < -9_999:
		return 6
	case v < -999:
		return 5
	case v < -99:
		return 4
	case v < -9:
		return 3
	default:
		return 2
	}
}

func pInt64StringSize(v int64) uint8 {
	switch true {
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
