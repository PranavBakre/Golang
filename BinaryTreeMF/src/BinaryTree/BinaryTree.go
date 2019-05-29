package BinaryTree

import (
	"../node"
	"fmt"
	"../queue"
	"../stack"
)

type Tree struct {
	Root *Node.Node
}

func (t *Tree) CreateNR() {
	choice := 1
	var lor string
	if t.Root == nil {
		t.Root = &Node.Node{}
		fmt.Println("Enter the root value")
		fmt.Scanln(&t.Root.Data)

	}
	for choice == 1 {
		curr := t.Root
		temp := &Node.Node{}
		fmt.Println("Enter the node data")
		fmt.Scanln(&temp.Data)

		for {

			fmt.Println("Enter the position:  L/R ?")
			fmt.Scanln(&lor)

			if lor == "l" || lor == "L" {
				if curr.Left == nil {
					curr.Left = temp
					break
				} else {
					curr = curr.Left
				}
			} else if lor == "r" || lor == "R" {
				if curr.Right == nil {
					curr.Right = temp
					break
				} else {
					curr = curr.Right
				}
			}

		}
		fmt.Println("Continue?")
		fmt.Scanln(&choice)
	}
}

func (t *Tree) CreateR() {
	var new string
	if t.Root != nil {
		fmt.Println("Create a new Tree?")
		fmt.Scanln(new)
		if new == "y" || new == "Y" {
			//T.erase()
		}
	}

	t.Root = CreateR1(t.Root)
}

func CreateR1(temp *Node.Node) *Node.Node {
	var pos string
	if temp == nil {
		var d1 = &Node.Node{}
		fmt.Println("Enter the value of the node")
		fmt.Scanln(&d1.Data)
		temp = d1

	}
	fmt.Println("Add the node on the left ?")
	fmt.Scanln(&pos)
	if pos == "l" || pos == "L" {
		temp.Left = CreateR1(temp.Left)
	}

	fmt.Println("Add the node on the right ?")
	fmt.Scanln(&pos)
	if pos == "r" || pos == "R" {
		temp.Right = CreateR1(temp.Right)
	}

	return temp
}

func (t *Tree) Predisplay(temp *Node.Node) {
	if temp != nil {
		fmt.Print(temp.Data, " ")
		t.Predisplay(temp.Left)
		t.Predisplay(temp.Right)
	}
}

func (t *Tree) Indisplay(temp *Node.Node) {
	if temp != nil {

		t.Indisplay(temp.Left)
		fmt.Print(temp.Data, " ")
		t.Indisplay(temp.Right)
	}
}

func (t *Tree) Postdisplay(temp *Node.Node) {
	if temp != nil {

		t.Postdisplay(temp.Left)
		t.Postdisplay(temp.Right)
		fmt.Print(temp.Data, " ")
	}
}

func (t *Tree) PreNRdisplay() {
	temp := t.Root
	var s Stack.Stack
	s.Init()
	for {
		for temp != nil {

			s.Push(temp)
			fmt.Print(temp.Data, " ")
			temp = temp.Left
		}

		if s.IsEmpty() {
			break
		} else {
			temp = s.Pop()
			if temp != nil {
				temp = temp.Right
			}
		}
	}
}

func (t *Tree) InNRdisplay() {
	temp := t.Root
	var s Stack.Stack
	s.Init()
	for {
		for temp != nil {

			s.Push(temp)

			temp = temp.Left
		}

		if s.IsEmpty() {
			break
		} else {
			temp = s.Pop()

			if temp != nil {
				fmt.Print(temp.Data, " ")
				temp = temp.Right
			}
		}
	}
}

func (t *Tree) PostNRdisplay() {
	temp := t.Root
	fmt.Println()
	var s Stack.Stack
	s.Init()
	for {
		for temp != nil {
			s.Push(temp)
			temp = temp.Left
		}

/*		if s.Data[s.Top].Right == nil {
			temp = s.Pop()
			fmt.Print(temp.Data, " ")
		}*/
		for !s.IsEmpty()  {
			if temp == s.Data[s.Top].Right {
			temp = s.Pop()
			fmt.Print(temp.Data, " ")
			}else {
				break
			}

		}
		if s.IsEmpty() {
			return

		}

			temp = s.Data[s.Top].Right

	}
}

func (t *Tree)BFS() {
	var q Queue.Queue
	var temp *Node.Node
	q.Init()
	q.Enqueue(t.Root)
	for !q.IsEmpty() {
		temp=q.Dequeue()

		fmt.Print(temp.Data," ")
		if temp.Left!=nil {
			q.Enqueue(temp.Left)
			
		}
		if temp.Right!=nil {
			q.Enqueue(temp.Right)
		}
	}
	
}

func (t *Tree ) Erase() {
	t.Erase1(t.Root)
	t.Root=nil
}

func (t *Tree) Erase1(temp *Node.Node) {
	if temp!=nil {
		t.Erase1(temp.Left)
		t.Erase1(temp.Right)
		temp=nil
	}
}


