package main

import (
	"fmt"
	"math/rand"
)

// Эта функция берет два отсортированных массива (левый и правый), сравнивает их элементы и постепенно создает
// новый отсортированный массив путем последовательного выбора минимального элемента среди двух.
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// mergeSort рекурсивно сортирует массив Эта функция рекурсивно разбивает входящий массив на меньшие части до тех пор,
// пока каждая часть не будет содержать максимум один элемент. После этого вызывает вспомогательную функцию merge()
//
//	для объединения этих частей обратно в отсортированном виде.
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}
	fmt.Println(ar)
	ar = mergeSort(ar)
	fmt.Println(ar)
}
