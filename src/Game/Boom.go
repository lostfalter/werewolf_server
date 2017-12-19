package game

import (
	"strings"
)

// WaitForBoom waits wolves to boom
func WaitForBoom(boom chan bool) chan string {
	input := make(chan string)

	go func() {
		for {
			select {
			case msg := <-input:
				if strings.HasPrefix(msg, "boom") {
					boom <- true
					break
				} else if strings.HasPrefix(msg, "finish") {
					break
				}
			}
		}
	}()

	return input
}
