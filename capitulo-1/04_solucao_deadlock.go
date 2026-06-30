package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	var a, b value

	printSum := func(v1, v2 *value) {
		defer wg.Done()

		// Trava, lê e destrava v1 independentemente
		v1.mu.Lock()
		val1 := v1.value
		v1.mu.Unlock()

		time.Sleep(2 * time.Second) // Simulação de latência

		// Trava, lê e destrava v2 independentemente
		v2.mu.Lock()
		val2 := v2.value
		v2.mu.Unlock()

		// A operação matemática e de I/O é feita fora da zona de lock
		fmt.Printf("sum=%d\n", val1+val2)
	}

	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
