package reflect

import (
	"fmt"
	"runtime"
)

type obj int

func (o *obj) Private() {
	fmt.Println("private func LOL.", runtime.GOOS, runtime.GOARCH)
}
