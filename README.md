A monotonic clock and timer for Go. **Currently available only on Linux**, but I would like to expand it to support at least Windows and OS X in the near future.

#### Check out the GoDoc at https://godoc.org/github.com/pavben/monoclock

## Example

```go
package main

import (
	"fmt"
	"github.com/pavben/monoclock"
)

func main() {
	timer := monoclock.NewMonoTimer()

	// do something here that takes some time
	for i := 0; i < 5000; i++ {
		fmt.Printf("wheeeeee\n")
	}

	fmt.Printf("Spamming your console took %dms\n", timer.Value())
}
```
