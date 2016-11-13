package gtime

import (
	"fmt"
	"testing"
	"time"
)

func TestAlarm(t *testing.T) {
	var tickS int64 = 3
	var unit int64 = 5

	err := alarm(tickS, unit, 1, nil)
	if err == nil {
		t.Errorf("err(%v) is nil", err)
	}

	ct := 3
	now1 := time.Now()
	err = alarm(tickS, unit, -1, func(count int) bool {
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
	err = alarm(1, 2, 3, func(count int) bool {
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

func TestAlarmDay(t *testing.T) {
	err := AlarmDay("fake format", 1, nil)
	if err == nil {
		t.Errorf("err(%v) is nil", err)
	}

	now := time.Now().Local()
	num, count := 0, 1
	hour, minute, second := now.Hour(), now.Minute(), now.Second()+2
	err = AlarmDay(fmt.Sprintf("%02d:%02d:%02d", hour, minute, second), count, func(count int) bool {
		num++
		h, m, s := time.Now().Local().Clock()
		fmt.Println("now:", h, m, s)
		if h != hour {
			t.Errorf("now hour(%v) not equal (%v)", h, hour)
		}
		if m != minute {
			t.Errorf("now minute(%v) not equal (%v)", m, minute)
		}
		if s != second {
			t.Errorf("now second(%v) not equal (%v)", s, second)
		}

		return false
	})
	if err != nil {
		t.Error(err)
		return
	}
	if num != count {
		t.Errorf("call num(%v) not equal (%v)", num, count)
	}
}
