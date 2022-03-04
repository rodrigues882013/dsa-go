package main

import "data-structures/pkg/trees"

func main() {
	//profits := []int{1, 6, 10, 16}
	//weights := []int{1, 2, 3, 5}
	//
	//fmt.Println(dynamicprogramming.Knapsack(profits, weights, 7))
	//fmt.Println(dynamicprogramming.Knapsack(profits, weights, 6))

	tree := trees.AVLTree{}

	tree.Insert(10)
	tree.Insert(4)
	tree.Insert(20)
	tree.Insert(90)
	tree.Insert(3)
	tree.Insert(50)
	tree.Insert(2)
	tree.Insert(40)
	tree.Insert(91)
	tree.Insert(1)


}
