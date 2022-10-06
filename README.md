# task-master.go

customizable task execution system and library written in go with no dependancy

[![CodeQL](https://github.com/dev-asterix/task-master.go/actions/workflows/codeql.yml/badge.svg)](https://github.com/dev-asterix/task-master.go/actions/workflows/codeql.yml)

---

## Usage

```bash
go get github.com/dev-asterix/task-master.go
```

### simple scheduler based on time interval

```go
package main

import (
    "github.com/dev-asterix/task-master.go/scheduler"
)

func main() {
    // create a new scheduler
    s := scheduler.NewIntervalBased(10, scheduler.Second)

    for {

        // run tasks

        s.Schedule() // calculate next schedule and wait for it
        // runs after scheduled interval (here, 10 seconds)
    }
}

```

_Code.Share.Prosper_
