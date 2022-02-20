package main

import (
	"data-structures/pkg/sortalgorithm"
	"fmt"
)

func te(arr *[]int) {
	(*arr)[1] = 10
}

func main() {
	arr := []int{7, 6, 10, 5, 9, 2, 1, 15, 7}
	sortalgorithm.QuickSort(arr)
	fmt.Println(arr)

}
