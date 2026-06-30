package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	var sharedState int // Simulando um recurso real que precise de proteção
	const runtime = 1 * time.Second

	// Função genérica e otimizada para os workers
	worker := func(name string, iterations int, sleepTime time.Duration) {
		defer wg.Done()
		var count int

		for begin := time.Now(); time.Since(begin) <= runtime; {
			for i := 0; i < iterations; i++ {
				// O Lock protege APENAS a mutação de estado (microssegundos)
				sharedLock.Lock()
				sharedState++
				sharedLock.Unlock()

				// O processamento pesado / bloqueio (time.Sleep)
				// é movido para FORA da seção crítica.
				time.Sleep(sleepTime)
			}
			count++
		}
		fmt.Printf("%s worker executed %v loops.\n", name, count)
	}

	wg.Add(2)
	// Ambos concorrem de forma muito mais limpa, reduzindo a contenção da trava
	go worker("Greedy-like", 1, 3*time.Nanosecond)
	go worker("Polite-like", 3, 1*time.Nanosecond)

	wg.Wait()
	fmt.Printf("Final shared state: %d\n", sharedState)
}
