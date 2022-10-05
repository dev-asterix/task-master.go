package scheduler

import (
	"fmt"
	"time"
)

const (
	NanoSecond  Duration = "ns"
	MicroSecond Duration = "µs"
	MilliSecond Duration = "ms"
	Second      Duration = "sec"
	Minute      Duration = "min"
	Hour        Duration = "hr"
	Day         Duration = "day"
	Week        Duration = "week"
	Month       Duration = "month"
	Year        Duration = "year"
)

type (
	Duration string

	Interval struct {
		frequency int64
		unit      Duration
		nextRun   time.Time
		timezone  *time.Location
	}

	Scheduler struct {
		Interval
	}
)

// timeNow returns the current time
// added for unit testing purpose
var timeNow = func(location *time.Location) time.Time {
	return time.Now().In(location)
}

// NewIntervalBased creates a new scheduler with the given interval
// timezone is optional, if not provided, the scheduler will use the local timezone, if multiple are provided, only the 1st entry will be considered
// repeat if set as true, the scheduler will run every freq interval
func NewIntervalBased(freq int64, unit Duration, timezone ...*time.Location) *Interval {
	if timezone == nil {
		timezone = []*time.Location{time.Local}
	}
	if freq <= 0 {
		freq = 1
	}

	return &Interval{
		frequency: freq,
		unit:      unit,
		timezone:  timezone[0],
	}
}

// Frequency returns the frequency of the scheduler
func (i *Interval) Frequency() string {
	return fmt.Sprintf("%d * %s", i.frequency, i.unit)
}

// Frequency returns the frequency of the scheduler
func (i *Interval) NextSchedule() string {
	return i.nextRun.String()
}

// Next returns the next time the scheduler should run
func (i *Interval) Next() *Interval {
	now := timeNow(i.timezone)

	switch i.unit {
	case NanoSecond:
		i.nextRun = now.Add(time.Duration(i.frequency) * time.Nanosecond)
	case MicroSecond:
		i.nextRun = now.Add(time.Duration(i.frequency) * time.Microsecond)
	case MilliSecond:
		i.nextRun = now.Add(time.Duration(i.frequency) * time.Millisecond)
	case Second:
		i.nextRun = now.Add(time.Duration(i.frequency) * time.Second)
	case Minute:
		i.nextRun = now.Add(time.Duration(i.frequency) * time.Minute)
	case Hour:
		i.nextRun = now.Add(time.Duration(i.frequency) * time.Hour)
	case Day:
		i.nextRun = now.AddDate(0, 0, int(i.frequency))
	case Week:
		i.nextRun = now.AddDate(0, 0, int(i.frequency*7))
	case Month:
		i.nextRun = now.AddDate(0, int(i.frequency), 0)
	case Year:
		i.nextRun = now.AddDate(int(i.frequency), 0, 0)
	}
	return i
}

// Schedule waits for the next run of the scheduler based on the next run
func (i *Interval) Schedule() time.Time {
	if i.nextRun.IsZero() || time.Now().After(i.nextRun) {
		i.Next()
	}
	t := <-time.After(time.Until(i.nextRun))
	return t
}
