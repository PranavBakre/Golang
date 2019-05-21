package Stack

import (
	"../node"
	"fmt"
)

const Max = 1000

type Stack struct {
	Data [Max]*Node.Node
	Top  int
}

func (s *Stack) Init() {
	s.Top = -1
}

func (s *Stack) Push(d1 *Node.Node) {
	s.Top++
	s.Data[s.Top] = d1
}

func (s *Stack) Pop() *Node.Node {
	temp := s.Data[s.Top]
	s.Top--
	return temp
}

func (s *Stack) IsEmpty() bool {
	if s.Top == -1 {
		return true
	}
	return false
}

func (s *Stack) Display() {
	for i := s.Top; i > -1; i-- {
		fmt.Println(s.Data[i].Data)
	}
}
