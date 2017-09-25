package algorithms

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAlgorithms(t *testing.T) {

	Convey("Given a list of transactions", t, func() {

		Convey("For validate that pencil is subset of [pencil beats]", func() {
			ss := Subset([]string{"pencil", "beatss"}, []string{"iphone", "pencil", "beatss", "power bank"})
			Convey("The output should be true", func() {
				So(ss, ShouldEqual, true)
			})
		})

	})
}
