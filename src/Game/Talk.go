package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Talk for vote
func Talk(players []Player) {
	// just talk
	fmt.Println("Start Talk")

	SingleTalk()

}

func SingleTalk() {
	timer1 := time.NewTimer(time.Second * 115)
	// go func() {
	// 	<-timer1.C
	// 	fmt.Println("Timer 1 expired")

	// 	done
	// }()

	go func() {
		for true {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			fmt.Print(text)

			if strings.HasPrefix(text, "exit") {
				timer1.Stop()
				break
			}
		}
	}()

	<-timer1.C
	fmt.Println("Timer 1 expired")
}
