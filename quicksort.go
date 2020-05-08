package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := os.Args[1:]
	numbers := make([]int, len(in))

	for i, n := range in {
		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("%s not a valid number! \n", n)
			os.Exit(1)
		}
		numbers[i] = number
	}
	fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int) []int {
	listSize := len(numbers)
	if listSize <= 1 {
		return numbers
	}

	n := make([]int, listSize)
	copy(n, numbers)
	pivotIndex := listSize / 2
	pivot := n[pivotIndex]
	n = append(n[:pivotIndex], n[pivotIndex+1:]...)
	smaller, greater := splitter(n, pivot)
	return append(append(quicksort(smaller), pivot), quicksort(greater)...)
}

func splitter(numbers []int, pivot int) (smaller []int, greater []int) {
	for _, v := range numbers {
		if v >= pivot {
			greater = append(greater, v)
		} else {
			smaller = append(smaller, v)
		}
	}
	return smaller, greater
}

