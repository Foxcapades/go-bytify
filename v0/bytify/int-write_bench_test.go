package bytify_test

import (
	"github.com/foxcapades/go-bytify/v0/bytify"
	"math"
	"strconv"
	"testing"
)

var tmp []byte

func BenchmarkInt8ToByteSlice_Zero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int8ToByteSlice(0)
	}
}
func BenchmarkInt8ToByteSlice_MinInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int8ToByteSlice(math.MinInt8)
	}
}
func BenchmarkInt8ToByteSlice_MaxInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int8ToByteSlice(math.MaxInt8)
	}
}

////////////////////////////////////////////////////////////////////////////////

func BenchmarkInt16ToByteSlice_Zero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int16ToByteSlice(0)
	}
}
func BenchmarkInt16ToByteSlice_MinInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int16ToByteSlice(math.MinInt8)
	}
}
func BenchmarkInt16ToByteSlice_MaxInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int16ToByteSlice(math.MaxInt8)
	}
}
func BenchmarkInt16ToByteSlice_MinInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int16ToByteSlice(math.MinInt16)
	}
}
func BenchmarkInt16ToByteSlice_MaxInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int16ToByteSlice(math.MaxInt16)
	}
}

////////////////////////////////////////////////////////////////////////////////

func BenchmarkInt32ToByteSlice_Zero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int32ToByteSlice(0)
	}
}
func BenchmarkInt32ToByteSlice_MinInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int32ToByteSlice(math.MinInt8)
	}
}
func BenchmarkInt32ToByteSlice_MaxInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int32ToByteSlice(math.MaxInt8)
	}
}
func BenchmarkInt32ToByteSlice_MinInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int32ToByteSlice(math.MinInt16)
	}
}
func BenchmarkInt32ToByteSlice_MaxInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int32ToByteSlice(math.MaxInt16)
	}
}
func BenchmarkInt32ToByteSlice_MinInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int32ToByteSlice(math.MinInt32)
	}
}
func BenchmarkInt32ToByteSlice_MaxInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int32ToByteSlice(math.MaxInt32)
	}
}

////////////////////////////////////////////////////////////////////////////////

func BenchmarkInt64ToByteSlice_Zero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int64ToByteSlice(0)
	}
}
func BenchmarkInt64ToByteSlice_MinInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int64ToByteSlice(math.MinInt8)
	}
}
func BenchmarkInt64ToByteSlice_MaxInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int64ToByteSlice(math.MaxInt8)
	}
}
func BenchmarkInt64ToByteSlice_MinInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int64ToByteSlice(math.MinInt16)
	}
}
func BenchmarkInt64ToByteSlice_MaxInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int64ToByteSlice(math.MaxInt16)
	}
}
func BenchmarkInt64ToByteSlice_MinInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int64ToByteSlice(math.MinInt32)
	}
}
func BenchmarkInt64ToByteSlice_MaxInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int64ToByteSlice(math.MaxInt32)
	}
}
func BenchmarkInt64ToByteSlice_MinInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int64ToByteSlice(math.MinInt64)
	}
}
func BenchmarkInt64ToByteSlice_MaxInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = bytify.Int64ToByteSlice(math.MaxInt64)
	}
}

////////////////////////////////////////////////////////////////////////////////

func BenchmarkAppendInt_Zero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = strconv.AppendInt(nil, 0, 10)
	}
}
func BenchmarkAppendInt_MinInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = strconv.AppendInt(nil, math.MinInt8, 10)
	}
}
func BenchmarkAppendInt_MaxInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = strconv.AppendInt(nil, math.MaxInt8, 10)
	}
}
func BenchmarkAppendInt_MinInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = strconv.AppendInt(nil, math.MinInt16, 10)
	}
}
func BenchmarkAppendInt_MaxInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = strconv.AppendInt(nil, math.MaxInt16, 10)
	}
}
func BenchmarkAppendInt_MinInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = strconv.AppendInt(nil, math.MinInt32, 10)
	}
}
func BenchmarkAppendInt_MaxInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = strconv.AppendInt(nil, math.MaxInt32, 10)
	}
}
func BenchmarkAppendInt_MinInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = strconv.AppendInt(nil, math.MinInt64, 10)
	}
}
func BenchmarkAppendInt_MaxInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp = strconv.AppendInt(nil, math.MaxInt64, 10)
	}
}
