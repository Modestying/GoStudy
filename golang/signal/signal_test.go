package signal_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestSignal(t *testing.T) {
	wait := &sync.WaitGroup{}
	wait.Add(1)
	go ElegantClose(&wait)
	wait.Wait()
	fmt.Println("service Exit")
}
