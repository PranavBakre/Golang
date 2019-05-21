package Queue

import (
	"../node"
	"fmt"
)

const Max = 1000

type Queue struct {
	Data        [Max]*Node.Node
	Front, Rear int
}

func (q *Queue) Init() {
	q.Front = -1
	q.Rear = -1
}
func (q *Queue) Enqueue(d1 *Node.Node) {
	q.Rear++
	q.Data[q.Rear] = d1

}

func (q *Queue) Dequeue() *Node.Node {

	q.Front++
	temp := q.Data[q.Front]
	return temp
}

func (q *Queue) IsEmpty() bool {
	if q.Front == q.Rear {
		return true
	}
	return false
}

func (q *Queue) Display() {
	for i := q.Front + 1; i <= q.Rear; i++ {
		fmt.Printf("%s\t", q.Data[i].Data)
	}
	fmt.Println()
}
