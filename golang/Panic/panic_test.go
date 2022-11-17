package panic

import "testing"
import "fmt"

func TestPanicWithRecover(t *testing.T) {
	PanicWithRecover()
}

func TestPanicWithoutPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Msg:", r.(error).Error())
		}
	}()
	PanicWithoutRecover()
}
