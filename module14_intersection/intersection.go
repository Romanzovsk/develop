package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция для чтения чисел с клавиатуры
func readIntFromInput(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	num, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		panic(err)
	}
	return num
}

// Функция для ввода массива строк
func readArray(size int) []string {
	var array []string
	for i := 0; i < size; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Введите элемент %d: ", i+1)
		element, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		array = append(array, strings.TrimSpace(element))
	}
	return array
}

// Основная логика программы
func main() {
	// Чтение размеров массивов
	sizeA := readIntFromInput("Введите размер первого массива: ")
	sizeB := readIntFromInput("Введите размер второго массива: ")

	// Чтение элементов массивов
	arrA := readArray(sizeA)
	arrB := readArray(sizeB)

	// Использование map для нахождения общего множества
	commonElementsMap := make(map[string]bool)
	result := make([]string, 0)

	// Проходим по первому массиву и добавляем каждый элемент в map
	for _, element := range arrA {
		commonElementsMap[element] = true
	}

	// Проходим по второму массиву и проверяем наличие в первом
	for _, element := range arrB {
		if commonElementsMap[element] {
			result = append(result, element)
		}
	}

	// Вывод результата
	fmt.Println("Общие элементы:")
	if len(result) > 0 {
		for _, elem := range result {
			fmt.Println(elem)
		}
	} else {
		fmt.Println("Нет общих элементов.")
	}
}
