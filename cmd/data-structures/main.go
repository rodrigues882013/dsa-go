package main

import (
	"data-structures/pkg/graph"
	"fmt"
)

//func main() {
//	//profits := []int{1, 6, 10, 16}
//	//weights := []int{1, 2, 3, 5}
//	//
//	//fmt.Println(dynamicprogramming.Knapsack(profits, weights, 7))
//	//fmt.Println(dynamicprogramming.Knapsack(profits, weights, 6))
//
//	//tree := trees.AVLTree{}
//	//
//	//tree.Insert(10)
//	//tree.Insert(4)
//	//tree.Insert(20)
//	//tree.Insert(90)
//	//tree.Insert(3)
//	//tree.Insert(50)
//	//tree.Insert(2)
//	//tree.Insert(40)
//	//tree.Insert(91)
//	//tree.Insert(1)
//
//
//
//}

func main() {
	edges := [][]int{

		{0, 7, 1},
		{0, 9, 3},
		{0, 11, 5},
		{7, 11, 1},
		{7, 6, 5},
		{7, 3, 7},
		{6, 5, 10},
		{3, 4, 20},
		{2, 3, 1},
		{2, 12, 5},
		{12, 8, 4},
		{8, 1, 18},
		{1, 10, 7},
		{10, 9, 6},
		{9, 8, 2},
	}

	myGraph := graph.LGraph{}
	myOtherGraph := graph.MGraph{}

	for i := 0; i < 15; i++ {
		myGraph.AddVertex(i)
		myOtherGraph.AddVertex(i)
	}

	myGraph.AddEdges(edges)
	myOtherGraph.AddEdges(edges)

	//fmt.Println(myGraph.BFS(0))
	//fmt.Println(myOtherGraph.DFS(0))
	fmt.Println(myGraph.Dijkstra(0))
	//fmt.Println(myOtherGraph)
}
