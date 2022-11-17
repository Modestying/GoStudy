package reflect

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestObj(t *testing.T) {
	var o obj
	method := reflect.ValueOf(&o).MethodByName("Private")
	method.Call(nil) // stdout ...

	var h = http.Header{"k": {"v"}}
	clone := reflect.ValueOf(&h).MethodByName("Clone")
	fmt.Println(clone.Call(nil)[0]) // stdout map[k:[v]]
}
