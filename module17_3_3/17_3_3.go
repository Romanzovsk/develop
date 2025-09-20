package main

import (
	"fmt"
	"sync"
)

const step int64 = 1
const finalCounterValue int64 = 1000

func main() {
	var counter int64 = 0
	var mutex sync.Mutex
	var cv sync.Cond
	cv.L = &mutex

	done := make(chan struct{}) // Канал для оповещения главной горутины о достижении цели

	increment := func() {
		mutex.Lock()
		defer mutex.Unlock()

		current := counter
		newVal := current + step
		counter = newVal

		if newVal >= finalCounterValue {
			close(done) // Закрываем канал, означая конец процесса
		}
	}

	// Запустим N горутин, чтобы увеличивать счётчик параллельно
	for i := 0; i < 1000; i++ {
		go increment()
	}

	// Ждём закрытия канала
	<-done

	// Завершаем ожидание и выводим результат
	fmt.Println(counter)
}
