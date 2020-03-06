package parse

import "time"

const timeTpl = "2006-01-02"

func Time2String(t time.Time) string {
	return t.Format(timeTpl)
}

func String2Time(s string) time.Time {
	loc, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation(timeTpl, s, loc)
	return t
}
