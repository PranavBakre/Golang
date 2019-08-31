package graph

import (
	"fmt"
	"sync"
)

type Node struct {
	Distance []int
	Next     []int
}

type Graph struct {
	NodeNumber      int
	Graph           []Node
	Update          []Node
	InitData        []Node
	IsUpdatePresent []bool
	Neighbour       [][]int
}

func (G *Graph) Init() {
	fmt.Println("Enter the number of nodes in the graph")
	fmt.Scanln(&G.NodeNumber)
	fmt.Println(G.NodeNumber)

	G.Graph = make([]Node, G.NodeNumber)
	G.InitData = make([]Node, G.NodeNumber)
	G.Update = make([]Node, G.NodeNumber)
	G.IsUpdatePresent = make([]bool, G.NodeNumber)
	G.Neighbour = make([][]int, G.NodeNumber)
	for i := 0; i < G.NodeNumber; i++ {
		G.Graph[i].Distance = make([]int, G.NodeNumber)
		G.Update[i].Distance = make([]int, G.NodeNumber)
		G.Neighbour[i] = make([]int, G.NodeNumber)
		G.InitData[i].Distance = make([]int, G.NodeNumber)
		G.Graph[i].Next = make([]int, G.NodeNumber)
		G.Update[i].Next = make([]int, G.NodeNumber)
		G.InitData[i].Next = make([]int, G.NodeNumber)
	}
	for i := 0; i < G.NodeNumber; i++ {
		for j := 0; j < G.NodeNumber; j++ {
			G.Graph[i].Distance[j] = -1
			G.Graph[i].Next[j] = -1
			G.Update[i].Distance[j] = 999
		}
	}

	G.Display()
}

func (G *Graph) AddData() {
	fmt.Println("Enter the Distance between the nodes. 99 for infinity")
	for i := 0; i < G.NodeNumber; i++ {

		for j := 0; j < G.NodeNumber; j++ {

			if i == j {
				G.Graph[i].Distance[j] = 0
				G.Graph[j].Distance[i] = 0
			}

			if G.Graph[i].Distance[j] == -1 {
				fmt.Println("Node", i, "and", j, ": ")
				fmt.Scanln(&G.Graph[i].Distance[j])
				G.Graph[j].Distance[i] = G.Graph[i].Distance[j]
				if G.Graph[i].Distance[j] != 999 {
					G.Neighbour[i][j] = 1
					G.Neighbour[j][i] = G.Neighbour[i][j]
					G.Graph[i].Next[j] = j
					G.Graph[j].Next[i] = i
				}
			}
			G.InitData[i].Distance[j] = G.Graph[i].Distance[j]
			G.InitData[i].Next[j] = G.Graph[i].Next[j]
		}
	}

	G.Display()
}

func (G *Graph) Display() {
	for i := 0; i < G.NodeNumber; i++ {
		fmt.Println(G.Graph[i])
	}
}

func (G *Graph) UpdateData(Node int) {
	for i := 0; i < G.NodeNumber; i++ {
		if i != Node && G.Graph[Node].Distance[i] > G.Update[Node].Distance[i] {
			G.Graph[Node].Distance[i] = G.Update[Node].Distance[i]
			G.Graph[Node].Next[i] = G.Update[Node].Next[i]
		}
	}
	G.IsUpdatePresent[Node] = false
}

func (G *Graph) SendData(wg *sync.WaitGroup, Sender int, Node int) {
	for i := 0; i < G.NodeNumber; i++ {
		if i != Node {
			v := G.Graph[Sender].Distance[i] + G.Graph[Sender].Distance[Node]
			if v < G.Update[Node].Distance[i] {
				G.Update[Node].Distance[i] = v
				G.Update[Node].Next[i] = Sender
				G.IsUpdatePresent[Node] = true
			}
		}
	}
	wg.Done()
}

func (G *Graph) Run() {
	var k int
	var wg sync.WaitGroup
	for j := 0; j < G.NodeNumber; j++ {
		if G.Neighbour[0][j] == 1 {
			wg.Add(1)
			go G.SendData(&wg, 0, j)

		}
	}
	wg.Wait()
	for i := 0; i < G.NodeNumber; i = (i + 1) % G.NodeNumber {
		if G.IsUpdatePresent[i] {
			G.UpdateData(i)

			for j := 0; j < G.NodeNumber; j++ {
				if G.Neighbour[i][j] == 1 {
					wg.Add(1)
					G.SendData(&wg, i, j)

				}
			}
			wg.Wait()
		}

		for k = 0; k < G.NodeNumber && G.IsUpdatePresent[k] == false; k++ {
		}
		if k == G.NodeNumber && i == G.NodeNumber-1 {
			break
		}
	}
	G.Display()

}

func (G *Graph) ChangeData() {
	var i, j, c int
	fmt.Println("Enter the nodes and new cost")
	fmt.Scanln(&i, &j, &c)
	G.Graph = G.InitData

	G.Graph[i].Distance[j] = c
	G.Graph[j].Distance[i] = c
	if c != 999 {
		G.Graph[i].Next[j] = j
		G.Graph[j].Next[i] = i
	}
	for ii := 0; ii < G.NodeNumber; ii++ {
		for jj := 0; jj < G.NodeNumber; jj++ {
			G.Update[ii].Distance[jj] = 999
		}
	}
}
