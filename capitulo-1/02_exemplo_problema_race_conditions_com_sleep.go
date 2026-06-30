package main

import (
	"fmt"
	"time"
)

func main() {

	var data int

	go func() {
		data++
	}()
	time.Sleep(1 * time.Second)
	if data == 0 {
		fmt.Println("The value is %d.", data)
	}
}
