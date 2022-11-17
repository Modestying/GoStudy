package dd

import (
	"reflect"
	"testing"
)

func BenchmarkDB_Get(b *testing.B) {

	db := new(DB)
	m := map[string]func(){
		"Get":  db.Get,
		"Get1": db.Get1,
		"Get2": db.Get2,
		"Get3": db.Get3,
		"Get4": db.Get4,
	}
	b.Run("reflect_bench", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflect.ValueOf(db).MethodByName("Get").Call(nil)
		}
	})

	b.Run("map_bench", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			method, _ := m["Get"]
			method()
		}
	})
}
