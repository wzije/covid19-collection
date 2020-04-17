package utils

import (
	"log"
	"time"
)

type MyTime struct {
	now    time.Time
	offset int
}

func Time() MyTime {
	now := time.Now()
	_, offset := now.Zone()

	return MyTime{now: now, offset: offset}
}

func (t MyTime) Now() time.Time {
	return t.now
}

func (t MyTime) Offset() int {
	return t.offset
}

func (t MyTime) parse(time time.Time, offset int) time.Time {
	//return time - (offset * 60000)
	return time
}

func IsSameDate(t1 time.Time, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	if y1 == y2 && m1 == m2 && d1 == d2 {
		log.Print("Same Day")
		return true
	}

	return false
}
