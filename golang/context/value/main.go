package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctxVal := context.WithValue(ctx, "s", "xx")
	fmt.Println(ctxVal.Value("s"))

}
