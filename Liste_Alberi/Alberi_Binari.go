package main

import (
	"fmt"
)

type bitreeNode struct {
	left  *bitreeNode
	right *bitreeNode
	val   int
}

type bitree struct {
	root *bitreeNode
}

func inorder(node *bitreeNode) {
	if node != nil {
		inorder(node.left)
		fmt.Println(node.val)
		inorder(node.right)
	}
}

func preorder(node *bitreeNode) {
	if node != nil {
		fmt.Println(node.val)
		preorder(node.left)
		preorder(node.right)
	}
}

func postorder(node *bitreeNode) {
	if node != nil {
		postorder(node.left)
		postorder(node.right)
		fmt.Println(node.val)
	}
}

func newNode(value int) *bitreeNode {
	node := new(bitreeNode)
	node.val = value
	node.left = nil
	node.right = nil
	return node
}

func arr2tree(a []int, i int) (root *bitreeNode) {
	if i < len(a) {
		root = newNode(a[i])
		root.left = arr2tree(a, 2*i+1)
		root.right = arr2tree(a, 2*i+2)
	}
	return
}

//78 [54-, 90 [19, 95]-], 21 [16 [5,-],19 [56,43]]]

func stampaAlbero(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val, "")
	if node.left == nil && node.right == nil {
		return
	}
	fmt.Print(" [")
	if node.left != nil {
		stampaAlbero(node.left)
	} else {
		fmt.Print("-")
	}
	fmt.Print(", ")
	if node.right != nil {
		stampaAlbero(node.right)
	} else {
		fmt.Print("-")
	}
	fmt.Print("]")
}

func StampaSommario(node *bitreeNode, spaces int) {
	if node == nil {
		return
	}

	// Stampa l'asterisco e il valore del nodo
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	fmt.Print("*")
	fmt.Println(node.val)

	// Stampa i figli con l'indentazione corretta
	if node.left != nil {
		StampaSommario(node.left, spaces+1)
	}
	if node.right != nil {
		StampaSommario(node.right, spaces+1)
	}

	// Stampa un asterisco solo se il nodo ha un solo figlio
	if (node.left == nil) != (node.right == nil) {
		for i := 0; i < spaces; i++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	}
}

func main() {
	tree := new(bitree)
	tree.root = newNode(78)
	tree.root.left = newNode(54)
	tree.root.right = newNode(21)
	tree.root.left.right = newNode(90)
	tree.root.left.right.right = newNode(95)
	tree.root.left.right.left = newNode(19)
	tree.root.right.left = newNode(16)
	tree.root.right.left.left = newNode(5)
	tree.root.right.right = newNode(19)
	tree.root.right.right.right = newNode(43)
	tree.root.right.right.left = newNode(56)

	preorder(tree.root)
	fmt.Println()
	postorder(tree.root)
	fmt.Println()
	inorder(tree.root)

	StampaSommario(tree.root, 0)
	stampaAlbero(tree.root)

	fmt.Println()

	a := []int{69, 89, 28, 39, 66, 44, 12, 2, 71}
	b := arr2tree(a, 0)
	stampaAlbero(b)

}
