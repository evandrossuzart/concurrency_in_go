package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	for a := 0; a < 100; a++ {
		go func() {
			var internalData int
			internalData++

			ch <- internalData
		}()

		// A extração do dado do canal atua como uma barreira de sincronização nativa
		data := <-ch

		if data == 0 {
			fmt.Printf("The value is %d.\n", data)
		} else {
			fmt.Printf("The value is %d.\n", data)
		}
	}
}

//for i in {1..10}; do go run 02_solucao_data_race.go; done
//for i in {1..10}; do go run 02_problema_data_race.go; done
