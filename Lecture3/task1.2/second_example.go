package main

import (
	"fmt"
	"time"
)

// это второй пример для deadlock. В этом коде горутины вызывают sleepyGopher, но не отправляют значения в канал.
// Главная горутина все равно пытается получить значения из канала, что приведет к deadlock.
func main() {
	channel := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, channel)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-channel
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}

func sleepyGopher(id int, channel chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")
	// Не отправляем значение в канал
}
