package main

import (
	"fmt"
	"sync"
)

// MergeChannels объединяет данные из n каналов в один и возвращает его
func MergeChannels(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(ch <-chan int) {
		defer wg.Done()
		for val := range ch {
			out <- val
		}
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go func() {
		defer close(channel1)
		for i := 1; i <= 10; i++ {
			channel1 <- i
		}
	}()

	go func() {
		defer close(channel2)
		for i := 6; i <= 10; i++ {
			channel2 <- i
		}
	}()

	merge := MergeChannels(channel1, channel2)

	for value := range merge {
		fmt.Println(value)
	}
}
