package main

import (  
    "fmt"
    "time"
)

func main() {  
    arr := []int{101, 95, 10, 188, 100}
    max := arr[0]
    for _, v := range arr {
        if v > max {
            max = v
        }
    }
    fmt.Printf("Max element is %d\n", max)
    var s string
    for {
      t1 := time.Now()
      fmt.Scan(&s)
      t2 := time.Now()
      if s == "q"{
        break
      }
      fmt.Println(t2.Sub(t1).Seconds())
    }
}
