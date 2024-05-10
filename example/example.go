package example

import (
	"fmt"

	"github.com/got-reload/demo/example2"
)

var (
	f = new(float64)
)

func F1() int {
	*f += 0.1
	fmt.Printf("f: %0.3f, sin %0.3f\n", *f, example2.Sin(*f))
	return 1
}
