package main

import (
	"./src/BinaryTree"
	"fmt"
	"os"
	"syscall"
)

func main() {

	fd, err := os.Open("test.txt")
	if err != nil {
		os.Create("test.txt")
		syscall.Exec("/bin/gnome-terminal", []string{"gnome-terminal", "--", "go", "run", "BinaryTree.go"}, os.Environ())
	} else {
		fd.Close()
		os.Remove("test.txt")
	}
	var bst BinaryTree.BST
	var key, exit string
	bst.CreateNR()
	bst.InNRdisplay()
	fmt.Println()
	fmt.Println(bst.Height(bst.Root))
	fmt.Println("Enter the key")
	fmt.Scanln(&key)
	node, _ := bst.Search(key)
	if node != nil {
		fmt.Println(node.Data)
	}else {
		fmt.Println(key," not found\n")
	}
	bst.Delete()
	fmt.Println("\nInorder display")
	bst.InNRdisplay()
	fmt.Println("\nBFS display\n")
	bst.BFS()

	var bst2 BinaryTree.BST
	bst2.CopyTreeNR(&bst)
	fmt.Println()
	fmt.Println()
	fmt.Println("Displaying Copied Tree")
	bst2.InNRdisplay()
	fmt.Println("\nMirror\nInorder display")
	bst2.MirrorNR()
	bst2.InNRdisplay()
	fmt.Println()
	fmt.Println("Press Enter to exit\n")
	fmt.Scanln(&exit)

}
