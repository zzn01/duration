# duration
time.Duration string parser

## usage

```golang
package main

import (
	"fmt"
	"time"

	"github.com/zzn01/duration"
)

func main() {
	d := 2 * time.Hour
	d2, _ := duration.Parse(fmt.Sprintf("%s", d))
	fmt.Println(d == d2)
}
```
