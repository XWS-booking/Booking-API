package model

import (
	"time"
)

type TimeInterval struct {
	From time.Time `bson:"from" json:"from"`
	To   time.Time `bson:"to" json:"to"`
}

func (timeInterval *TimeInterval) IsOverlapping(interval TimeInterval) bool {
	start := max(timeInterval.From, interval.From)
	end := min(timeInterval.To, interval.To)
	return !(start.After(end) || start.Equal(end))
}

func (timeInterval *TimeInterval) TryAppendInterval(interval TimeInterval) bool {
	if timeInterval.To.Equal(interval.From) {
		timeInterval.To = interval.To
		return true
	}
	return false
}

func (timeInterval *TimeInterval) GetDaysDifference() int32 {
	return int32(timeInterval.To.Sub(timeInterval.From).Hours() / 24)
}

func (timeInterval *TimeInterval) HasTimeIntervalInside(interval TimeInterval) bool {
	return (interval.From.After(timeInterval.From) || interval.From.Equal(timeInterval.From)) &&
		(interval.To.Before(timeInterval.To) || interval.To.Equal(timeInterval.To))
}

func max(time1 time.Time, time2 time.Time) time.Time {
	if time1.After(time2) {
		return time1
	}
	return time2
}

func min(time1 time.Time, time2 time.Time) time.Time {
	if time1.Before(time2) {
		return time1
	}
	return time2
}
