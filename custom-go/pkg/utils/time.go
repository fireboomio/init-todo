package utils

import "time"

func TodayBegin() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func TodayEnd() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()).Add(-1 * time.Nanosecond)
}

func CurrentDateTime() string {
	return time.Now().Format(time.RFC3339)
}
