package bytify_test

import (
	"bytes"
	"strconv"
	"testing"

	. "github.com/foxcapades/go-bytify/v0/bytify"
	. "github.com/smartystreets/goconvey/convey"
)

type test struct {
	raw  uint64
	size uint8
	val  string
}

var (
	uint8s = []test{
		{0, 1, "0"},
		{1, 1, "1"},
		{10, 2, "10"},
		{100, 3, "100"},
	}

	uint16s = append(uint8s,
		test{256, 3, "256"},
		test{1_000, 4, "1000"},
		test{10_000, 5, "10000"},
	)

	uint32s = append(uint16s,
		test{65536, 5, "65536"},
		test{100_000, 6, "100000"},
		test{1_000_000, 7, "1000000"},
		test{10_000_000, 8, "10000000"},
		test{100_000_000, 9, "100000000"},
		test{1_000_000_000, 10, "1000000000"},
	)
	uint64s = append(uint32s,
		test{4_294_967_296, 10, "4294967296"},
		test{10_000_000_000, 11, "10000000000"},
		test{100_000_000_000, 12, "100000000000"},
		test{1_000_000_000_000, 13, "1000000000000"},
		test{10_000_000_000_000, 14, "10000000000000"},
		test{100_000_000_000_000, 15, "100000000000000"},
		test{1_000_000_000_000_000, 16, "1000000000000000"},
		test{10_000_000_000_000_000, 17, "10000000000000000"},
		test{100_000_000_000_000_000, 18, "100000000000000000"},
		test{1_000_000_000_000_000_000, 19, "1000000000000000000"},
		test{10_000_000_000_000_000_000, 20, "10000000000000000000"},
	)
)

func TestUint8ToBytes(t *testing.T) {
	Convey("Uint8ToBytes", t, func() {
		for _, test := range uint8s {
			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatUint(test.raw, 10)
				buffer := make([]byte, test.size)
				Convey("when given uint8("+stdForm+")", func() {
					Uint8ToBytes(uint8(test.raw), buffer)

					So(buffer, ShouldResemble, []byte(test.val))
				})
			})
		}
	})
}

func TestUint16ToBytes(t *testing.T) {
	Convey("Uint16ToBytes", t, func() {
		for _, test := range uint16s {
			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatUint(test.raw, 10)
				buffer := make([]byte, test.size)
				Convey("when given uint16("+stdForm+")", func() {
					Uint16ToBytes(uint16(test.raw), buffer)

					So(buffer, ShouldResemble, []byte(test.val))
				})
			})
		}
	})
}


func TestUint32ToBytes(t *testing.T) {
	Convey("Uint32ToBytes", t, func() {
		for _, test := range uint32s {
			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatUint(test.raw, 10)
				buffer := make([]byte, test.size)
				Convey("when given uint32("+stdForm+")", func() {
					Uint32ToBytes(uint32(test.raw), buffer)

					So(buffer, ShouldResemble, []byte(test.val))
				})
			})
		}
	})
}


func TestUint64ToBytes(t *testing.T) {
	Convey("Uint64ToBytes", t, func() {
		for _, test := range uint64s {
			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatUint(test.raw, 10)
				buffer := make([]byte, test.size)
				Convey("when given uint64("+stdForm+")", func() {
					Uint64ToBytes(test.raw, buffer)

					So(buffer, ShouldResemble, []byte(test.val))
				})
			})
		}
	})
}


func TestUint8ToBuf(t *testing.T) {
	Convey("Uint8ToBuf", t, func() {
		for _, test := range uint8s {
			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatUint(test.raw, 10)
				buffer := new(bytes.Buffer)
				Convey("when given uint8("+stdForm+")", func() {
					Uint8ToBuf(uint8(test.raw), buffer)

					So(buffer.Bytes(), ShouldResemble, []byte(test.val))
				})
			})
		}
	})
}

func TestUint16ToBuf(t *testing.T) {
	Convey("Uint16ToBuf", t, func() {
		for _, test := range uint16s {
			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatUint(test.raw, 10)
				buffer := new(bytes.Buffer)
				Convey("when given uint16("+stdForm+")", func() {
					Uint16ToBuf(uint16(test.raw), buffer)

					So(buffer.Bytes(), ShouldResemble, []byte(test.val))
				})
			})
		}
	})
}


func TestUint32ToBuf(t *testing.T) {
	Convey("Uint32ToBuf", t, func() {
		for _, test := range uint32s {

			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatUint(test.raw, 10)
				buffer := new(bytes.Buffer)
				Convey("when given uint32("+stdForm+")", func() {
					Uint32ToBuf(uint32(test.raw), buffer)

					So(buffer.Bytes(), ShouldResemble, []byte(test.val))
				})
			})
		}
	})
}


func TestUint64ToBuf(t *testing.T) {
	Convey("Uint64ToBuf", t, func() {
		for _, test := range uint64s {
			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatUint(test.raw, 10)
				buffer := new(bytes.Buffer)
				Convey("when given uint64("+stdForm+")", func() {
					Uint64ToBuf(test.raw, buffer)

					So(buffer.Bytes(), ShouldResemble, []byte(test.val))
				})
			})
		}
	})
}
