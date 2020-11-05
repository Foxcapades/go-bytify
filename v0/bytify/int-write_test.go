package bytify_test

import (
	"bytes"
	"strconv"
	"testing"

	. "github.com/foxcapades/go-bytify/v0/bytify"
	. "github.com/smartystreets/goconvey/convey"
)

type itest struct {
	raw  int64
	size uint8
	val  string
}

var (
	int8s = []itest{
		{-128, 4, "-128"},
		{-10, 3, "-10"},
		{-1, 2, "-1"},
		{0, 1, "0"},
		{1, 1, "1"},
		{10, 2, "10"},
		{127, 3, "127"},
	}

	int16s = append(int8s,
		itest{-32_768, 6, "-32768"},
		itest{-1_000, 5, "-1000"},
		itest{-256, 4, "-256"},
		itest{256, 3, "256"},
		itest{1_000, 4, "1000"},
		itest{32_767, 5, "32767"},
	)

	int32s = append(int16s,
		itest{-2_147_483_648, 11, "-2147483648"},
		itest{-100_000_000, 10, "-100000000"},
		itest{-10_000_000, 9, "-10000000"},
		itest{-1_000_000, 8, "-1000000"},
		itest{-100_000, 7, "-100000"},
		itest{-65536, 6, "-65536"},
		itest{65536, 5, "65536"},
		itest{100_000, 6, "100000"},
		itest{1_000_000, 7, "1000000"},
		itest{10_000_000, 8, "10000000"},
		itest{100_000_000, 9, "100000000"},
		itest{2_147_483_647, 10, "2147483647"},
	)

	int64s = append(int32s,
		itest{-9_223_372_036_854_775_808, 20, "-9223372036854775808"},
		itest{-100_000_000_000_000_000, 19, "-100000000000000000"},
		itest{-10_000_000_000_000_000, 18, "-10000000000000000"},
		itest{-1_000_000_000_000_000, 17, "-1000000000000000"},
		itest{-100_000_000_000_000, 16, "-100000000000000"},
		itest{-10_000_000_000_000, 15, "-10000000000000"},
		itest{-1_000_000_000_000, 14, "-1000000000000"},
		itest{-100_000_000_000, 13, "-100000000000"},
		itest{-10_000_000_000, 12, "-10000000000"},
		itest{-4_294_967_296, 11, "-4294967296"},
		itest{4_294_967_296, 10, "4294967296"},
		itest{10_000_000_000, 11, "10000000000"},
		itest{100_000_000_000, 12, "100000000000"},
		itest{1_000_000_000_000, 13, "1000000000000"},
		itest{10_000_000_000_000, 14, "10000000000000"},
		itest{100_000_000_000_000, 15, "100000000000000"},
		itest{1_000_000_000_000_000, 16, "1000000000000000"},
		itest{10_000_000_000_000_000, 17, "10000000000000000"},
		itest{100_000_000_000_000_000, 18, "100000000000000000"},
		itest{9_223_372_036_854_775_807, 19, "9223372036854775807"},
	)
)

func TestInt8ToBytes(t *testing.T) {
	Convey("Int8ToBytes", t, func() {
		for _, test := range int8s {
			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatInt(test.raw, 10)
				buffer := make([]byte, test.size)
				Convey("when given int8("+stdForm+")", func() {
					Int8ToBytes(int8(test.raw), buffer)

					So(string(buffer), ShouldResemble, test.val)
				})
			})
		}
	})
}

func TestInt16ToBytes(t *testing.T) {
	Convey("Int16ToBytes", t, func() {
		for _, test := range int16s {
			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatInt(test.raw, 10)
				buffer := make([]byte, test.size)
				Convey("when given int16("+stdForm+")", func() {
					Int16ToBytes(int16(test.raw), buffer)

					So(string(buffer), ShouldResemble, test.val)
				})
			})
		}
	})
}


func TestInt32ToBytes(t *testing.T) {
	Convey("Int32ToBytes", t, func() {
		for _, test := range int32s {
			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatInt(test.raw, 10)
				buffer := make([]byte, test.size)
				Convey("when given int32("+stdForm+")", func() {
					Int32ToBytes(int32(test.raw), buffer)

					So(string(buffer), ShouldResemble, test.val)
				})
			})
		}
	})
}


func TestInt64ToBytes(t *testing.T) {
	Convey("Int64ToBytes", t, func() {
		for _, test := range int64s {
			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatInt(test.raw, 10)
				buffer := make([]byte, test.size)
				Convey("when given int64("+stdForm+")", func() {
					Int64ToBytes(test.raw, buffer)

					So(string(buffer), ShouldResemble, test.val)
				})
			})
		}
	})
}


func TestInt8ToBuf(t *testing.T) {
	Convey("Int8ToBuf", t, func() {
		for _, test := range int8s {
			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatInt(test.raw, 10)
				buffer := new(bytes.Buffer)
				Convey("when given int8("+stdForm+")", func() {
					Int8ToBuf(int8(test.raw), buffer)

					So(buffer.String(), ShouldResemble, test.val)
				})
			})
		}
	})
}

func TestInt16ToBuf(t *testing.T) {
	Convey("Int16ToBuf", t, func() {
		for _, test := range int16s {
			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatInt(test.raw, 10)
				buffer := new(bytes.Buffer)
				Convey("when given int16("+stdForm+")", func() {
					Int16ToBuf(int16(test.raw), buffer)

					So(buffer.String(), ShouldResemble, test.val)
				})
			})
		}
	})
}


func TestInt32ToBuf(t *testing.T) {
	Convey("Int32ToBuf", t, func() {
		for _, test := range int32s {

			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatInt(test.raw, 10)
				buffer := new(bytes.Buffer)
				Convey("when given int32("+stdForm+")", func() {
					Int32ToBuf(int32(test.raw), buffer)

					So(buffer.String(), ShouldResemble, test.val)
				})
			})
		}
	})
}


func TestInt64ToBuf(t *testing.T) {
	Convey("Int64ToBuf", t, func() {
		for _, test := range int64s {
			Convey(`returns "` + test.val + `"`, func() {
				stdForm := strconv.FormatInt(test.raw, 10)
				buffer := new(bytes.Buffer)
				Convey("when given int64("+stdForm+")", func() {
					Int64ToBuf(test.raw, buffer)

					So(buffer.String(), ShouldResemble, test.val)
				})
			})
		}
	})
}

func TestInt8ToByteSlice(t *testing.T) {
	Convey("Int8ToByteSlice", t, func() {
		for i := 127; i > -128; i-- {
			So(string(Int8ToByteSlice(int8(i))), ShouldResemble, strconv.Itoa(i))
		}
	})
}

func TestInt16ToByteSlice(t *testing.T) {
	Convey("Int16ToByteSlice", t, func() {
		for i := 32767; i > -32768; i -= 257 {
			So(string(Int16ToByteSlice(int16(i))), ShouldResemble, strconv.Itoa(i))
		}
	})
}

func TestInt32ToByteSlice(t *testing.T) {
	Convey("Int32ToByteSlice", t, func() {
		for i := 2147483647; i > -2147483648; i -= 16843009 {
			So(string(Int32ToByteSlice(int32(i))), ShouldResemble, strconv.Itoa(i))
		}
	})
}

func TestInt64ToByteSlice(t *testing.T) {
	Convey("Int64ToByteSlice", t, func() {
		for i := int64(9223372036854775807); i > -9223372036854775808; i -= int64(72340172838076673) {
			So(string(Int64ToByteSlice(i)), ShouldResemble, strconv.FormatInt(i, 10))
		}
	})
}
