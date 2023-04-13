package time

import "time"

// Today return today start time and end time
func Today() (start, end time.Time) {
	return DayDeadline(time.Now())
}

// DayDeadline 获取t所在日期的当前开始与结束时间
func DayDeadline(t time.Time) (start, end time.Time) {
	t = t.Local()
	start = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return start, start.AddDate(0, 0, 1)
}
