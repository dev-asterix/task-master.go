package main

import (
	"fmt"

	"github.com/dev-asterix/task-master.go/scheduler"
)

func main() {
	intervalScheduler := scheduler.NewIntervalBased(10, scheduler.Second)

	for {

		// perform the actions required here if action is to be performed 1st and then wait for the next interval

		scheduleWaitDoneAt := intervalScheduler.Schedule()

		// perform the actions required here if action is to be performed 2nd and then run the next interval

		fmt.Println("scheduler running", intervalScheduler.NextSchedule(), scheduleWaitDoneAt.String())
	}
}
