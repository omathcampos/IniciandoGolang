package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("Olá Mundo") // Goroutine
		waitGroup.Done()      // -1 do contador
	}()

	go func() {
		escrever("Programando em Go!") // Goroutine
		waitGroup.Done()
	}()

	waitGroup.Wait() // 0 Ele serve para fazer com que as gourotines esperem e sejam todas as gourotines e funções executadas.
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
