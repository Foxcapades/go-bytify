package internal

import (
	"github.com/foxcapades/go-bytify/v0/bytify"
	"strconv"
	"time"
)

func Uint8ToBytes(tv uint64, it int) time.Duration {
	v := uint8(tv)
	buffer := make([]byte, bytify.Uint64StringSize(tv))

	start := time.Now()
	for i := 0; i < it; i++ {
		bytify.Uint8ToBytes(v, buffer)
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

func Uint32ToBytes(tv uint64, it int) time.Duration {
	v := uint32(tv)
	buffer := make([]byte, bytify.Uint64StringSize(tv))

	start := time.Now()
	for i := 0; i < it; i++ {
		bytify.Uint32ToBytes(v, buffer)
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

func StdLibToBytes(tv uint64, it int) time.Duration {
	buffer := make([]byte, bytify.Uint64StringSize(tv))

	start := time.Now()
	for i := 0; i < it; i++ {
		strconv.AppendUint(buffer, tv, 10)
	}

	return time.Now().Sub(start)
}
