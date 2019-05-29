package BinaryTree

import (
	"../node"
	"../queue"
	"../stack"
	"fmt"
)

type BST struct {
	Root *Node.Node
}

func (t *BST) Create() {
	var choice = 1
	var temp *Node.Node
	for choice == 1 {
		temp = &Node.Node{}
		fmt.Println("Enter the node value")
		fmt.Scanln(&temp.Data)
		t.Root = t.CreateR(t.Root, temp)
		fmt.Println("Continue? ")
		fmt.Scanln(&choice)
	}

}

func (t *BST) CreateR(temp *Node.Node, d1 *Node.Node) *Node.Node {
	if temp == nil {
		return d1
	}
	if temp.Data > d1.Data {
		temp.Left = t.CreateR(temp.Left, d1)
	} else if temp.Data < d1.Data {
		temp.Right = t.CreateR(temp.Right, d1)
	} else {
		fmt.Println("Error.Data Exists")

	}

	return temp
}

func (t *BST) CreateNR() {

	choice := 1
	if t.Root == nil {
		t.Root = &Node.Node{}
		fmt.Println("Enter the root value")
		fmt.Scanln(&t.Root.Data)

	}
	for choice == 1 {
		curr := t.Root
		temp := &Node.Node{}
		fmt.Println("Enter the node value")
		fmt.Scanln(&temp.Data)
		for {
			if curr.Data > temp.Data {
				if curr.Left == nil {
					curr.Left = temp
					break
				} else {
					curr = curr.Left
				}
			} else if curr.Data < temp.Data {
				if curr.Right == nil {
					curr.Right = temp
					break
				} else {
					curr = curr.Right
				}
			} else {
				fmt.Println("Error.Data exists")
				break
			}
		}
		fmt.Println("Continue ?")
		fmt.Scanln(&choice)

	}

}

func (t *BST) InNRdisplay() {
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

func (t *BST) BFS() {
	var q Queue.Queue
	var temp *Node.Node
	q.Init()
	q.Enqueue(t.Root)
	for !q.IsEmpty() {
		temp = q.Dequeue()

		fmt.Print(temp.Data, " ")
		if temp.Left != nil {
			q.Enqueue(temp.Left)

		}
		if temp.Right != nil {
			q.Enqueue(temp.Right)
		}
	}

}

func (t *BST) Height(temp *Node.Node) int {
	if temp == nil {
		return 0
	}

	h1 := t.Height(temp.Left)
	h2 := t.Height(temp.Right)
	if h1 > h2 {
		h1++
		return h1
	}
	h2++
	return h2

}

func (t *BST) Search(key string) (*Node.Node, *Node.Node) {
	temp := t.Root
	var parent *Node.Node = nil
	for temp != nil {
		if temp.Data == key {
			break
		} else if temp.Data > key {
			parent = temp
			temp = temp.Left
		} else if temp.Data < key {
			parent = temp
			temp = temp.Right
		}
	}
	if temp == nil {
		parent = nil
	}

	return temp, parent

}

func (t *BST) Delete() {
	var key string
	fmt.Println("Enter the key to be searched")
	fmt.Scanln(&key)
	var srch, parent = t.Search(key)
	if srch != nil {
		if srch == t.Root {
			if srch.Left == nil && srch.Right != nil {
				t.Root = t.Root.Right
				srch = nil
			} else if srch.Left != nil && srch.Right == nil {
				t.Root = t.Root.Left
				srch = nil
			} else if srch.Left != nil && srch.Right != nil {
				trvse := srch.Right
				for trvse.Left != nil {
					trvse = trvse.Left
				}
				trvse.Left = srch.Left
				t.Root = t.Root.Right
			}
		} else {
			if srch.Left == nil && srch.Right != nil {
				if parent.Left == srch {
					parent.Left = srch.Right
				} else if parent.Right == srch {
					parent.Right = srch.Right
				}
				srch.Left = nil
				srch.Right = nil
				srch = nil
			} else if srch.Left != nil && srch.Right == nil {
				if parent.Left == srch {
					parent.Left = srch.Left
				} else if parent.Right == srch {
					parent.Right = srch.Left
				}
			} else if srch.Left == nil && srch.Right == nil {
				srch = nil
			} else {

				trvse := srch.Right
				for trvse.Left != nil {
					trvse = trvse.Left
				}
				trvse.Left = srch.Left
				if parent.Left == srch {
					parent.Left = srch.Right
				} else if parent.Right == srch {
					parent.Right = srch.Right
				}
				srch = nil
			}
		}
	}
}

func (t *BST) CopyTreeR(t2 *BST) {

	t.Root = CopyTreeR2(t2.Root)
}

func CopyTreeR2(temp *Node.Node) *Node.Node {
	var t2 *Node.Node = nil
	if temp != nil {
		t2 = &Node.Node{}
		t2.Data = temp.Data
		t2.Left = CopyTreeR2(temp.Left)
		t2.Right = CopyTreeR2(temp.Right)

	}
	return t2
}

func (t *BST) CopyTreeNR(t2 *BST) {
	t.Root = &Node.Node{}
	t.Root.Data = t2.Root.Data

	temp1 := t.Root
	temp2 := t2.Root
	var s1 Stack.Stack
	var s2 Stack.Stack
	s1.Init()
	s2.Init()
	for {
		for temp2 != nil {
			if temp2.Left != nil {
				temp1.Left = &Node.Node{}
				temp1.Left.Data = temp2.Left.Data
			}
			if temp2.Right != nil {
				temp1.Right = &Node.Node{}
				temp1.Right.Data = temp2.Right.Data
			}
			s1.Push(temp1)
			s2.Push(temp2)

			temp1 = temp1.Left
			temp2 = temp2.Left

		}
		if s2.IsEmpty() {
			break
		} else {
			temp1 = s1.Pop()
			temp2 = s2.Pop()
			temp1 = temp1.Right
			temp2 = temp2.Right
		}

	}
}

func (t *BST) MirrorR(temp *Node.Node) {

	if temp != nil {
		t2 := temp.Left
		temp.Left = temp.Right
		temp.Right = t2
		t.MirrorR(temp.Left)
		t.MirrorR(temp.Right)
	}
}

func (t *BST) MirrorNR() {
	var q Queue.Queue
	q.Init()


	q.Enqueue(t.Root)

	for !q.IsEmpty() {
		temp:=q.Dequeue()
		temp.Left,temp.Right=temp.Right,temp.Left
		if temp.Left!=nil {
			q.Enqueue(temp.Left)
		}
		if temp.Right!=nil {
			q.Enqueue(temp.Right)
		}

	}
}
