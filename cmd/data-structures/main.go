package main

import (
	"data-structures/pkg/dynamicprogramming"
	"fmt"
)

func main() {
	profits := []int{1, 6, 10, 16}
	weights := []int{1, 2, 3, 5}

	fmt.Println(dynamicprogramming.Knapsack(profits, weights, 7))
	fmt.Println(dynamicprogramming.Knapsack(profits, weights, 6))

}
