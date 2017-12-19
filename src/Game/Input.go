package game

import (
	"bufio"
	"fmt"
	"io"
	"sync"
)

type InputManager struct {
	listeners         []chan string
	singletonInstLock sync.Mutex
}

func (im *InputManager) Notify(s string) {
	//im.singletonInstLock.Lock()
	for i := range im.listeners {
		im.listeners[i] <- s
	}
}

func (im *InputManager) Register(c chan string) {
	//im.singletonInstLock.Lock()
	im.listeners = append(im.listeners, c)
}

// GetInput retrieve input and put it in channel
func GetInput(r io.Reader) (chan string, chan bool) {

	input := make(chan string)
	finish := make(chan bool)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			readLine := s.Text()
			input <- readLine

			select {
			case <-finish:
				fmt.Println("input finished!")
				break
			default:
			}
		}
	}()

	return input, finish
}

func (im *InputManager) WaitForInput(r io.Reader) chan bool {
	finish := make(chan bool)

	ic, f2 := GetInput(r)

	go func() {
		for {
			select {
			case msg := <-ic:
				im.Notify(msg)
			case <-finish:
				im.Notify("finish")
				f2 <- true
				break
			}
		}

	}()

	return finish
}
