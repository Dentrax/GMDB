package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestTrimSliceString(t *testing.T) {
	var testDatas = []struct {
		Input []string
		Title string
	}{
		{[]string{"The", "   100"}, "The+100"},
		{[]string{"Black   ", "Mirror"}, "Black+Mirror"},
		{[]string{"The   ", "   Rain"}, "The+Rain"},
		{[]string{"   Star   ", "   Wars   "}, "Star+Wars"},
		{[]string{"The", "Walking", "Dead"}, "The+Walking+Dead"},
	}

	for _, data := range testDatas {
		testName := strings.Join(data.Input, "")

		Convey("Testing arg input: ("+testName+")", t, func() {
			actual := TrimSliceString(data.Input)

			So(actual, ShouldNotBeBlank)
			So(actual, ShouldEqual, data.Title)
		})
	}
}
