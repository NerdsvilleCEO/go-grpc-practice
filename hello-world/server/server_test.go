package server

import (
	"net"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Given the server is started with defaults", t, func() {
		ListenRPC()
		Convey("When I hit port 8000 it should respond", func() {
			_, err := net.Dial("tcp", "localhost:8000")
			So(err, ShouldEqual, nil)
		})
	})
}
