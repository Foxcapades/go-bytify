package bytify_test

import (
	"fmt"
	"os"

	"github.com/foxcapades/go-bytify/v0/bytify"
)

func ExampleUint64ToBytes() {
	val := uint64(0xFFFF_FFFF_FFFF_FFFF)
	buf := make([]byte, bytify.Uint64StringSize(val))

	bytify.Uint64ToBytes(val, buf)

	fmt.Println(string(buf))
	// Output: 18446744073709551615
}

func ExampleUint64ToBuf() {
	val := uint64(0xFFFF_FFFF_FFFF_FFFF)
	buf := os.Stdout

	bytify.Uint64ToBuf(val, buf)
	// Output: 18446744073709551615
}

func ExampleUint32ToBytes() {
	val := uint32(0xFFFF_FFFF)
	buf := make([]byte, bytify.Uint32StringSize(val))

	bytify.Uint32ToBytes(val, buf)

	fmt.Println(string(buf))
	// Output: 4294967295
}

func ExampleUint32ToBuf() {
	val := uint32(0xFFFF_FFFF)
	buf := os.Stdout

	bytify.Uint32ToBuf(val, buf)
	// Output: 4294967295
}

func ExampleUint16ToBytes() {
	val := uint16(0xFFFF)
	buf := make([]byte, bytify.Uint16StringSize(val))

	bytify.Uint16ToBytes(val, buf)

	fmt.Println(string(buf))
	// Output: 65535
}

func ExampleUint16ToBuf() {
	val := uint16(0xFFFF)
	buf := os.Stdout

	bytify.Uint16ToBuf(val, buf)
	// Output: 65535
}

func ExampleUint8ToBytes() {
	val := uint8(0xFF)
	buf := make([]byte, bytify.Uint8StringSize(val))

	bytify.Uint8ToBytes(val, buf)

	fmt.Println(string(buf))
	// Output: 255
}

func ExampleUint8ToBuf() {
	val := uint8(0xFF)
	buf := os.Stdout

	bytify.Uint8ToBuf(val, buf)
	// Output: 255
}
