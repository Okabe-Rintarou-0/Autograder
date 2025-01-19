package utils

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestSendIfNotFull(t *testing.T) {
	c.Convey("TestSendIfNotFull", t, func() {
		bufferSize := 2
		goRoutinesCnt := 5
		ch := make(chan int, bufferSize)
		fullCnt := 0
		for i := 0; i < goRoutinesCnt; i++ {
			if SendIfNotFull(ch, i) {
				fullCnt++
			}
		}
		c.So(fullCnt, c.ShouldEqual, goRoutinesCnt-bufferSize)
	})
}
