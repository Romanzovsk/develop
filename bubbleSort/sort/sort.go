package main

import (
	"fmt"
	"math/rand"
)

func init() {
	// необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}
	fmt.Println(ar)
	bubbleSort(ar)

	fmt.Println(ar)
}

func bubbleSort(ar []int) []int {
	var isDone = false

	for !isDone {
		isDone = true
		for i := 0; i < len(ar)-1; i++ {
			if ar[i] < ar[i+1] {
				ar[i+1], ar[i] = ar[i], ar[i+1]
				isDone = false
			}
		}
	}
	return ar
}
