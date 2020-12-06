package internal

import (
	"strconv"
	"time"

	"github.com/foxcapades/go-bytify/v0/bytify"
)

var buffer []byte

func Uint8ToBytes(tv uint64, it int) time.Duration {
	v := uint8(tv)
	buffer := make([]byte, bytify.Uint64StringSize(tv))

	start := time.Now()
	for i := 0; i < it; i++ {
		bytify.Uint8ToBytes(v, buffer)
	}
	return time.Now().Sub(start)
}

func Uint8ToByteSlice(tv uint64, it int) time.Duration {
	v := uint8(tv)

	start := time.Now()
	for i := 0; i < it; i++ {
		buffer = bytify.Uint8ToByteSlice(v)
	}

	return time.Now().Sub(start)
}

func Uint16ToBytes(tv uint64, it int) time.Duration {
	v := uint16(tv)
	buffer := make([]byte, bytify.Uint64StringSize(tv))

	start := time.Now()
	for i := 0; i < it; i++ {
		bytify.Uint16ToBytes(v, buffer)
	}
	return time.Now().Sub(start)
}

func Uint16ToByteSlice(tv uint64, it int) time.Duration {
	v := uint16(tv)

	start := time.Now()
	for i := 0; i < it; i++ {
		buffer = bytify.Uint16ToByteSlice(v)
	}

	return time.Now().Sub(start)
}

func Uint32ToBytes(tv uint64, it int) time.Duration {
	v := uint32(tv)
	buffer := make([]byte, bytify.Uint64StringSize(tv))

	start := time.Now()
	for i := 0; i < it; i++ {
		bytify.Uint32ToBytes(v, buffer)
	}
	return time.Now().Sub(start)
}

func Uint32ToByteSlice(tv uint64, it int) time.Duration {
	v := uint32(tv)

	start := time.Now()
	for i := 0; i < it; i++ {
		buffer = bytify.Uint32ToByteSlice(v)
	}

	return time.Now().Sub(start)
}

func Uint64ToBytes(tv uint64, it int) time.Duration {
	buffer := make([]byte, bytify.Uint64StringSize(tv))

	start := time.Now()
	for i := 0; i < it; i++ {
		bytify.Uint64ToBytes(tv, buffer)
	}
	return time.Now().Sub(start)
}

func Uint64ToByteSlice(tv uint64, it int) time.Duration {
	start := time.Now()
	for i := 0; i < it; i++ {
		buffer = bytify.Uint64ToByteSlice(tv)
	}

	return time.Now().Sub(start)
}

func StdLibUintToBytes(tv uint64, it int) time.Duration {
	buffer := make([]byte, bytify.Uint64StringSize(tv))

	start := time.Now()
	for i := 0; i < it; i++ {
		strconv.AppendUint(buffer, tv, 10)
	}

	return time.Now().Sub(start)
}

func StdLibUintToByteSlice(tv uint64, it int) time.Duration {
	start := time.Now()
	for i := 0; i < it; i++ {
		buffer = strconv.AppendUint(nil, tv, 10)
	}

	return time.Now().Sub(start)
}
