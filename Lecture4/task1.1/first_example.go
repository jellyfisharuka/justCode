package main

import (
	"fmt"
	"sync"
)

func main() {
	//в этом коде у меня возникает race condition так как несколько горутин одновременно изменяют срез data без синхронизации и аутпуты тоже бывают разными.
	var data []int
	var wg sync.WaitGroup
	//var m sync.Mutex используя мьютекс можно правильно синхронизировать и избегать race condition
	var numOf = 1000
	for i := 0; i < numOf; i++ {
		wg.Add(1)
		go func() {
			//m.Lock()
			data = append(data, 42)
			//m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Print(len(data))
}
