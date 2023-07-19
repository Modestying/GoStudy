package main

import "fmt"

func main() {
	testData := "好i"
	fmt.Println(len(testData))
	testStr := []rune(testData)
	fmt.Println(len(testStr))

	for i := 0; i < len(testStr)/2; i++ {
		testStr[i], testStr[len(testStr)-i-1] = testStr[len(testStr)-i-1], testStr[i]
	}
	fmt.Println(string(testStr))
}
