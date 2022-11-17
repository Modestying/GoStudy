package signal_test

import (
	"fmt"
	"sync"
	"testing"
	"github.com/Modestying/GoStudy/golang/signal"

)

func TestSignal(t *testing.T) {
	wait := &sync.WaitGroup{}
	wait.Add(1)
	go ElegantClose(&wait)
	wait.Wait()
	fmt.Println("Service Exit")
}
