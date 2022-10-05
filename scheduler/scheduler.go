// Package scheduler provides a simple scheduler for Go
// it supports interval based scheduling
// it supports time based scheduling as well
package scheduler

import (
	"fmt"
	"time"
)

const (
	NanoSecond  Duration = "ns"    // Nanosecond identifier
	MicroSecond Duration = "µs"    // Microsecond identifier
	MilliSecond Duration = "ms"    // Millisecond identifier
	Second      Duration = "sec"   // Second identifier
	Minute      Duration = "min"   // Minute identifier
	Hour        Duration = "hr"    // Hour identifier
	Day         Duration = "day"   // Day identifier
	Week        Duration = "week"  // Week identifier
	Month       Duration = "month" // Month identifier
	Year        Duration = "year"  // Year identifier
)

type (

	// Duration is the unit of time for the scheduler
	//
	// it can be one of the following:
	// NanoSecond, MicroSecond, MilliSecond, Second, Minute, Hour, Day, Week, Month, Year
	Duration string

	// Interval is the scheduler that runs at a given interval
	// it can be used to run a task at a given interval
	//
	// see example for more details under examples/scheduler/simple/main.go
	Interval struct {
		frequency uint64
		unit      Duration
		nextRun   time.Time
		timezone  *time.Location
	}

	// Timed is the scheduler that runs at a given time
	Timed struct {
		Second uint8
		Minute uint8
		Hour   uint8
		Day    uint8
		Month  time.Month
		Year   uint32
	}

	// Scheduler is the interface that wraps the basic Schedule method.
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
// timezone is optional, if not provided, the scheduler will use the local timezone,
// if multiple are provided, only the 1st entry will be considered
// repeat if set as true, the scheduler will run every freq interval
//
// example:
//
//	// runs every second
//	// if the task takes 2 seconds to run, the next run will be 3 seconds after the previous run
//	i := scheduler.NewIntervalBased(1, scheduler.Second)
func NewIntervalBased(freq uint64, unit Duration, timezone ...*time.Location) *Interval {

	// default timezone is local
	if timezone == nil {
		timezone = []*time.Location{time.Local}
	}

	// create the scheduler based on interval
	return &Interval{
		frequency: freq,
		unit:      unit,
		timezone:  timezone[0],
	}
}

// Frequency returns the frequency of the scheduler in printable format (RFC3339)
func (i *Interval) Frequency() string {
	return fmt.Sprintf("%d * %s", i.frequency, i.unit)
}

// NextSchedule returns the next time the scheduler should run in printable format (RFC3339)
func (i *Interval) NextSchedule() string {
	return i.nextRun.String()
}

// Next sets the next run of the scheduler based on the frequency and unit at the timezone
// it will not run the scheduler itself, it will only set the next run, useful for debugging purposes.
//
// example:
//
//	// runs every second
//	i := scheduler.NewIntervalBased(1, scheduler.Second)
//	i.Next() // sets the next run to 1 second from now on i.nextRun field
//	i.NextSchedule() // returns the next run time in RFC3339 format
func (i *Interval) Next() {
	now := timeNow(i.timezone)

	// calculate the next run based on the frequency and unit
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
}

// Schedule waits for the next run of the scheduler based on the next run
// if the next run is in the past, it will get the next scheduled time based on interval
// if the next run is in the future, it will wait for the next run
//
// example:
//
//	// runs every second
//	// if the task takes 2 seconds to run, the next run will be 3 seconds after the previous run
//	i := scheduler.NewIntervalBased(1, scheduler.Second).Schedule()
func (i *Interval) Schedule() time.Time {
	if i.nextRun.IsZero() || time.Now().After(i.nextRun) {
		i.Next()
	}

	// wait for the next run, return the next run time once it's done
	t := <-time.After(time.Until(i.nextRun))
	return t
}
