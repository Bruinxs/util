package util

import (
	"fmt"
	"time"
)

func Alarm(delay, unit int64, count int, call func(count int) (interrupt bool)) error {
	if call == nil {
		return fmt.Errorf("arg call(%v) is nil", call)
	}
	var distance int64
	for i := 0; i < count || count == -1; i++ {
		nanoSecond := time.Now().Local().UnixNano()
		distance = delay - (nanoSecond % unit)
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
