package graph

type (
	// Operations is an interface holding all vertices most common operations regardless its kind of storage
	Operations interface {
		BFS(start int) []int
		DFS(start int) []int
		Dijkstra(start int) []int
		TopSort() []int
		AddVertices(vertices []int)
		AddVertex(vertex int)
		AddEdges(edges [][]int)
	}
)

/*

func main() {
	edges := [][]int{

		{ 0, 7, -1},
		{ 0, 9, -1},
		{ 0, 11, -1},
		{ 7, 11, -1},
		{ 7, 6, -1},
		{ 7, 3, -1},
		{ 6, 5, -1},
		{ 3, 4, -1},
		{ 2, 3, -1},
		{ 2, 12, -1},
		{ 12, 8, -1},
		{ 8, 1, -1},
		{ 1, 10, -1},
		{ 10, 9, -1},
		{ 9, 8, -1},
	}

	myGraph := graph.LGraph{}
	myOtherGraph := graph.MGraph{}

	for i := 0; i < 15; i++ {
		myGraph.AddVertex(i)
		myOtherGraph.AddVertex(i)
	}

	myGraph.AddEdges(edges)
	myOtherGraph.AddEdges(edges)

	fmt.Println(myGraph.BFS(0))
	fmt.Println(myOtherGraph.DFS(0))

	//fmt.Println(myOtherGraph)

}
*/
