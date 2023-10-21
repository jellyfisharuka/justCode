package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data    map[string]string
	rwMutex sync.RWMutex
)

func init() {
	data = make(map[string]string)
}

func readData(key string) string {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return data[key]
}

func writeData(key, value string) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	data[key] = value
}

func main() {
	// for Writing data
	go func() {
		for i := 0; i < 5; i++ {
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("value%d", i)
			writeData(key, value)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// for Reading data
	go func() {
		for i := 0; i < 5; i++ {
			key := fmt.Sprintf("key%d", i)
			result := readData(key)
			fmt.Printf("Read: %s => %s\n", key, result)
		}
	}()

	time.Sleep(time.Second * 2)
}
