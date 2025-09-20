package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Целевое значение счётчика
const targetCounterValue int64 = 1000

// Размер шага каждого прироста
const step int64 = 1

// Число работающих горутин
const numWorkers = 10

func main() {
	var counter int64 = 0
	var wg sync.WaitGroup

	// Вычисляем необходимое количество шагов каждой горутиной
	iterationsPerWorker := targetCounterValue / step / int64(numWorkers)

	// Функционал рабочей горутины
	increment := func(workerID int) {
		defer wg.Done()

		// Каждая горутина увеличивает счётчик указанное число раз
		for j := 0; j < int(iterationsPerWorker); j++ {
			atomic.AddInt64(&counter, step)
		}
	}

	// Запускаем нужное количество горутин
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go increment(i + 1)
	}

	// Ждем завершение всех горутин
	wg.Wait()

	// Выводим итоговое значение счётчика
	fmt.Println("Финальное значение:", counter)
}
