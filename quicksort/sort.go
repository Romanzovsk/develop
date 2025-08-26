package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left := 0
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	arr[0], arr[pivotIndex] = arr[pivotIndex], arr[0]

	for i := range arr {
		if arr[i] < pivot {
			left++
			arr[left], arr[i] = arr[i], arr[left]
		}
	}

	arr[0], arr[left] = arr[left], arr[0]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}
	fmt.Println(ar)

	quickSort(ar)

	fmt.Println(ar)
}
