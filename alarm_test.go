package util

import (
	"fmt"
	"testing"
	"time"
)

func TestAlarm(t *testing.T) {
	var tickS int64 = 3
	var unit int64 = 5

	err := Alarm(tickS, unit, 1, nil)
	if err == nil {
		t.Errorf("err(%v) is nil", err)
	}

	ct := 3
	now1 := time.Now()
	err = Alarm(tickS, unit, -1, func(count int) bool {
		fmt.Println("count:", count)
		if count == ct {
			return true
		}
		if count > ct {
			t.Errorf("count(%v) > %v", count, ct)
		}
		return false
	})
	if err != nil {
		t.Error(err)
	}
	now2 := time.Now()
	if now2.UnixNano()-now1.UnixNano() <= (int64(ct)-1)*unit {
		t.Errorf("second distance(%v) not equal (%v)", now2.UnixNano()-now1.UnixNano(), (int64(ct)-1)*unit)
	}

	num := 0
	err = Alarm(1, 2, 3, func(count int) bool {
		num++
		return false
	})
	if err != nil {
		t.Error(err)
	}
	if num != 3 {
		t.Errorf("call num(%v) not equal (%v)", num, 3)
	}
}
