import (
	"sync"
)

var mu sync.Mutex
var myMap = make(map[string]int)

func writeToMap(key string, value int) {
	mu.Lock()
	myMap[key] = value
	mu.Unlock()
}

func main() {
	// Запускаем несколько горутин, которые пытаются одновременно записать в карту
	for i := 0; i < 10; i++ {
		go writeToMap(fmt.Sprintf("key%d", i), i)
	}

	// Ждем завершения всех горутин
	// ...

	// Далее можно безопасно читать и изменять карту
}
