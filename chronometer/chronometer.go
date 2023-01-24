package chronometer

import (
	"fmt"
	"math"
	"time"
)

type _Status int64

const (
	Idle _Status = iota
	Running
	Stopped
)

var status _Status
var startedAt time.Time
var pausedAt time.Time

func Start() {
	if status == Idle {
		startedAt = time.Now()
		status = Running
	}
}

func Resume() {
	if status == Stopped {
		startedAt = startedAtWithPause()
		pausedAt = time.Time{}
		status = Running
	}
}

func Stop() {
	if status == Running {
		pausedAt = time.Now()
		status = Stopped
	}
}

func Reset() {
	startedAt = time.Time{}
	pausedAt = time.Time{}
	status = Idle
}

func Ellapsed() string {
	if (startedAt != time.Time{}) {
		ellapsed := time.Now().Sub(startedAtWithPause())
		fmt.Println(ellapsed)
		ellapsedMilliseconds := math.Floor(math.Abs(math.Mod(float64(ellapsed.Milliseconds())/100, 10)))
		ellapsedSeconds := math.Floor(math.Abs(math.Mod(ellapsed.Seconds(), 60)))
		ellapsedMinutes := math.Floor(math.Abs(math.Mod(ellapsed.Minutes(), 60)))
		if ellapsedMinutes == 0 {
			return fmt.Sprintf("%02.0f:%01.0f", ellapsedSeconds, ellapsedMilliseconds)
		}
		ellapsedHours := math.Floor(math.Abs(math.Mod(ellapsed.Hours(), 24)))
		if ellapsedHours == 0 {
			return fmt.Sprintf("%02.0f:%02.0f:%01.0f", ellapsedMinutes, ellapsedSeconds, ellapsedMilliseconds)
		}
		return fmt.Sprintf("%02.0f:%02.0f:%02.0f:%01.0f", ellapsedHours, ellapsedMinutes, ellapsedSeconds, ellapsedMilliseconds)
	} else {
		return "00:0"
	}
}

func Status() _Status {
	return status
}

func startedAtWithPause() time.Time {
	if (pausedAt != time.Time{}) {
		ellapsed := startedAt.Sub(pausedAt)
		return time.Now().Add(ellapsed)
	} else {
		return startedAt
	}
}
