package gtime

import (
	"fmt"
	"time"
)

const (
	SecondsPerMinute     int64 = 60
	SecondsPerHour       int64 = 60 * SecondsPerMinute
	SecondsPerDay        int64 = 24 * SecondsPerHour
	NanoSecondsPerSecond int64 = 1e9
)

type Caller func(count int) (interrupt bool)

func alarm(tickNanoSecond, unit int64, count int, call Caller) error {
	if call == nil {
		return fmt.Errorf("arg call(%v) is nil", call)
	}
	var distance int64
	for i := 0; i < count || count == -1; i++ {
		nanoSecond := time.Now().Local().UnixNano()
		distance = tickNanoSecond - (nanoSecond % unit)
		if distance == 0 {
			distance = unit
		} else {
			distance = (distance + unit) % unit
		}
		time.Sleep(time.Duration(distance))
		if call(i) {
			break
		}
	}
	return nil
}

//@arg:
//	timeFormat: string for 15:04:05
func AlarmDay(timeFormat string, count int, call Caller) error {
	tickTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2006-01-02 "+timeFormat, time.Local)
	if err != nil {
		return err
	}
	return alarm(tickTime.Local().UnixNano()%(SecondsPerDay*NanoSecondsPerSecond), SecondsPerDay*NanoSecondsPerSecond, count, call)
}
