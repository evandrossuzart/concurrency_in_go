package main

import (
	"fmt"
	"sync"
)

func main() {
	var data int
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Registra que há 1 goroutine para aguardar
	wg.Add(1)

	go func() {
		// Garante que o contador do WaitGroup seja decrementado ao final
		defer wg.Done()

		// Trava a memória para acesso exclusivo
		mu.Lock()
		data++
		mu.Unlock()
	}()

	// Bloqueia a goroutine main até que wg chegue a 0
	wg.Wait()

	// Trava novamente para leitura segura
	mu.Lock()
	if data == 1 {
		// Corrigido para fmt.Printf
		fmt.Printf("The value is %d.\n", data)
	}
	mu.Unlock()
}
