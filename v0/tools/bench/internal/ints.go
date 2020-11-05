package internal

import (
	"strconv"
	"time"

	"github.com/foxcapades/go-bytify/v0/bytify"
)

func Int8ToBytes(tv int64, it int) time.Duration {
	v := int8(tv)
	buffer := make([]byte, bytify.Int64StringSize(tv))

	start := time.Now()
	for i := 0; i < it; i++ {
		bytify.Int8ToBytes(v, buffer)
	}
	return time.Now().Sub(start)
}

func Int8ToByteSlice(tv int64, it int) time.Duration {
	var buffer []byte
	v := int8(tv)

	start := time.Now()
	for i := 0; i < it; i++ {
		buffer = bytify.Int8ToByteSlice(v)
	}
	dummy(buffer)

	return time.Now().Sub(start)
}

func Int16ToBytes(tv int64, it int) time.Duration {
	v := int16(tv)
	buffer := make([]byte, bytify.Int64StringSize(tv))

	start := time.Now()
	for i := 0; i < it; i++ {
		bytify.Int16ToBytes(v, buffer)
	}
	return time.Now().Sub(start)
}

func Int16ToByteSlice(tv int64, it int) time.Duration {
	var buffer []byte
	v := int16(tv)

	start := time.Now()
	for i := 0; i < it; i++ {
		buffer = bytify.Int16ToByteSlice(v)
	}
	dummy(buffer)

	return time.Now().Sub(start)
}

func Int32ToBytes(tv int64, it int) time.Duration {
	v := int32(tv)
	buffer := make([]byte, bytify.Int64StringSize(tv))

	start := time.Now()
	for i := 0; i < it; i++ {
		bytify.Int32ToBytes(v, buffer)
	}
	return time.Now().Sub(start)
}

func Int32ToByteSlice(tv int64, it int) time.Duration {
	var buffer []byte
	v := int32(tv)

	start := time.Now()
	for i := 0; i < it; i++ {
		buffer = bytify.Int32ToByteSlice(v)
	}
	dummy(buffer)

	return time.Now().Sub(start)
}

func Int64ToBytes(tv int64, it int) time.Duration {
	buffer := make([]byte, bytify.Int64StringSize(tv))

	start := time.Now()
	for i := 0; i < it; i++ {
		bytify.Int64ToBytes(tv, buffer)
	}
	return time.Now().Sub(start)
}

func Int64ToByteSlice(tv int64, it int) time.Duration {
	var buffer []byte

	start := time.Now()
	for i := 0; i < it; i++ {
		buffer = bytify.Int64ToByteSlice(tv)
	}
	dummy(buffer)

	return time.Now().Sub(start)
}

func StdLibIntToBytes(tv int64, it int) time.Duration {
	buffer := make([]byte, bytify.Int64StringSize(tv))

	start := time.Now()
	for i := 0; i < it; i++ {
		strconv.AppendInt(buffer, tv, 10)
	}

	return time.Now().Sub(start)
}

func StdLibIntToByteSlice(tv int64, it int) time.Duration {
	var buffer []byte

	start := time.Now()
	for i := 0; i < it; i++ {
		buffer = strconv.AppendInt(nil, tv, 10)
	}
	dummy(buffer)

	return time.Now().Sub(start)
}
