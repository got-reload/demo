package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/got-reload/demo/example"
)

func main() {
	fmt.Printf("Press enter to call example.F1 and example2.F2 repeatedly\n")
	fmt.Printf("Enter s to stop\n")

	loop()

	r := bufio.NewReader(os.Stdin)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		if string(line) == "s" {
			return
		}
		loop()
	}
}

func loop() bool {
	fmt.Printf("example.F1: %d\n", example.F1())
	return false
}
