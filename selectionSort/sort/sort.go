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

	selectionSort(ar)

	fmt.Println(ar)
}

func selectionSort(ar []int) []int {
	for i := 0; i < len(ar)-1; i++ {
		var minIndex = i

		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			var temp = ar[i]
			ar[i] = ar[minIndex]
			ar[minIndex] = temp
		}
	}
	return ar
}
