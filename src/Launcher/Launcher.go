package main

import (
	"fmt"

	"../Game"
)

func main() {
	fmt.Println("Game started!")

	good := 0
	bad := 0

	for i := 0; i < 100; i++ {
		if game.Start() {
			good++
		} else {
			bad++
		}
	}

	fmt.Printf("good rate = %.3f, bad rate = %.3f\n",
		float64(good)/float64(good+bad),
		float64(bad)/float64(good+bad))

	fmt.Println("Game finished!")
}
