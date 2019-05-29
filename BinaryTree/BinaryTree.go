package main

import (
	"fmt"
	"time"
	"./src/BinarySearchTree"
)

func main() {
	var bst BinarySearchTree.BST
	var key string
	bst.CreateNR()
	bst.InNRdisplay()
	fmt.Println()
	fmt.Println(bst.Height(bst.Root))
	fmt.Println("Enter the key")
	fmt.Scanln(&key)
	node, _ := bst.Search(key)
	if node != nil {
		fmt.Println(node.Data)
	}
	bst.Delete()
	bst.InNRdisplay()
	fmt.Println()
	bst.BFS()

	var bst2 BinarySearchTree.BST
	bst2.CopyTreeNR(&bst)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	bst2.InNRdisplay()
	fmt.Println()
	bst2.MirrorNR()
	bst2.InNRdisplay()
	fmt.Println()
	time.Sleep(5*time.Second)
}
