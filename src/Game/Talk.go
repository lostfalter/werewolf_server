package game

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

// Talk for vote
func Talk(players []Player) {
	// just talk
	fmt.Println("Start Talk")
	for _, p := range players {
		temp := make([]Player, 0)
		temp = append(temp, p)
		r := strings.NewReader(generateRandomTalkProcedure(temp))

		fmt.Printf("Player %d starts to talk    ", p.id)

		singleTalk(r)
	}
}

func singleTalk(r io.Reader) {
	timer1 := time.NewTimer(time.Second * 15)

	exit := make(chan bool)
	boom := make(chan bool)
	var im InputManager

	im.Register(checkExit(exit))
	im.Register(WaitForBoom(boom))

	finish := im.WaitForInput(r)

	select {
	case <-exit:
		fmt.Println("user end talk!")
	case <-boom:
		fmt.Println("somebody boom!")
	case <-timer1.C:
		fmt.Println("talk time expired")
	}

	finish <- true
}

func checkExit(exit chan bool) chan string {
	input := make(chan string)

	go func() {
		for {
			select {
			case msg := <-input:
				if strings.HasPrefix(msg, "exit") {
					exit <- true
					break
				} else if strings.HasPrefix(msg, "finish") {
					break
				}
			}
		}
	}()

	return input
}

func generateRandomTalkProcedure(players []Player) string {
	var result string
	behaviors := [2]string{"exit", "boom"}

	rand.Seed(time.Now().UnixNano())

	for range players {
		result += behaviors[rand.Intn(len(behaviors))] + "\n"
	}

	return result
}
