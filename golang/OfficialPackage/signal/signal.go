package signal

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func ElegantClose(wait *sync.WaitGroup) {
	c := make(chan os.Signal)
	signal.Notify(c)
	for data := range c {
		fmt.Println("Receive Signal :", data.String())
		switch data.String() {
		case "interrupt":
			close(c)
			break
		default:
			fmt.Println("default")
			break
		}
	}

	if _, ok := <-c; !ok {
		fmt.Println("Success Close Signal Channel")
	}

	wait.Done()
}
