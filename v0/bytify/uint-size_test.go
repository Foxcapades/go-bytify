package bytify_test

import (
	"strconv"
	"testing"

	"github.com/foxcapades/go-bytify/v0/bytify"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUint8StringSize(t *testing.T) {
	Convey("Uint8StringSize()", t, func() {
		for _, test := range uint8s {
			Convey(
				"returns "+strconv.Itoa(int(test.size))+
					"when called with "+strconv.FormatUint(test.raw, 10),
				func() {
					So(bytify.Uint8StringSize(uint8(test.raw)), ShouldEqual, test.size)
				})
		}
	})
}

func TestUint16StringSize(t *testing.T) {
	Convey("Uint16StringSize()", t, func() {
		for _, test := range uint16s {
			Convey(
				"returns "+strconv.Itoa(int(test.size))+
					"when called with "+strconv.FormatUint(test.raw, 10),
				func() {
					So(bytify.Uint16StringSize(uint16(test.raw)), ShouldEqual, test.size)
				})
		}
	})
}

func TestUint32StringSize(t *testing.T) {
	Convey("Uint32StringSize()", t, func() {
		for _, test := range uint32s {
			Convey(
				"returns "+strconv.Itoa(int(test.size))+
					"when called with "+strconv.FormatUint(test.raw, 10),
				func() {
					So(bytify.Uint32StringSize(uint32(test.raw)), ShouldEqual, test.size)
				})
		}
	})
}

func TestUint64StringSize(t *testing.T) {
	Convey("Uint64StringSize()", t, func() {
		for _, test := range uint64s {
			Convey(
				"returns "+strconv.Itoa(int(test.size))+
					"when called with "+strconv.FormatUint(test.raw, 10),
				func() {
					So(bytify.Uint64StringSize(test.raw), ShouldEqual, test.size)
				})
		}
	})
}
