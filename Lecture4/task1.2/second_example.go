package main

import (
	"fmt"
	"time"
)

func main() {
	//тут у нас возникает race condition так как две горутины работают асинхронно и могут получить доступ к x одновременно. Аутпут каждый раз может быть разным.
	//кратко говоря можем получить и четные, и нечетные числа.
	san := 0
	for {

		go func() {
			san++
		}()

		go func() {
			if san%2 == 0 {
				time.Sleep(1 * time.Millisecond) //я добавила тайм слип, чтобы заметить эффектов race condition.
				fmt.Println(san)
			}
		}()
	} //чтобы не было race condition, можно использовать мьютексы, каналы
}
