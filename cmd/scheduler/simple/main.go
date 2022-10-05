package main

import (
	"fmt"
	"time"

	"github.com/dev-asterix/task-master.go/scheduler"
)

func main() {
	i := scheduler.NewIntervalBased(5, scheduler.Second)

	for {

		// perform the actions required here if action is to be performed 1st and then wait for the next interval

		i.Schedule()

		// perform the actions required here if action is to be performed 2nd and then run the next interval

		fmt.Println("scheduler running", i.NextSchedule(), time.Now().String())
	}
}
