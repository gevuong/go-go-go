package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewDeck(t *testing.T) {

	Convey("Given an array of slice string", t, func() {
		d := newDeck()

		Convey("deck length should should equal 12", func() {
			So(len(d), ShouldEqual, 12)
		})

		Convey("first card should be Ace of Diamonds", func() {
			So(d[0], ShouldEqual, "Ace of Diamonds")
		})

		Convey("last card should be Three of Spades", func() {
			So(d[len(d)-1], ShouldEqual, "Three of Spades")
		})
	})
}
