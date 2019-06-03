package main

import "fmt"

type Node struct {
	Data   string
	Lchild *Node
	Rchild *Node
}

func main() {
	// root := buildBinaryTree()
	// fmt.Println("preOrlder")
	// preOderRecur(root)
	// fmt.Println()
	// fmt.Println("inOrlder")
	// inOrderRecur(root)
	// fmt.Println()
	// fmt.Println("posOrder")
	// posOrderRecur(root)
	btree := buildBinaryTree2()
	fmt.Println("preOrlder")
	preOderRecur1(btree, 0)

}

func buildBinaryTree() (root *Node) {
	root = &Node{Data: "root"}
	root.Lchild = &Node{Data: "A"}
	root.Rchild = &Node{Data: "B"}
	root.Lchild.Lchild = &Node{Data: "C"}
	root.Lchild.Rchild = &Node{Data: "D"}
	root.Rchild.Lchild = &Node{Data: "E"}
	root.Rchild.Rchild = &Node{Data: "F"}
	return
}

func buildBinaryTree2() (binaryTree []string) {
	binaryTree = make([]string, 7, 7)
	binaryTree[0] = "root"
	binaryTree[1] = "A"
	binaryTree[2] = "B"
	binaryTree[3] = "C"
	binaryTree[4] = "D"
	binaryTree[5] = "E"
	binaryTree[6] = "F"
	return
}

func preOderRecur(root *Node) {
	fmt.Printf(root.Data)
	fmt.Printf(" ")
	if root.Lchild != nil {
		preOderRecur(root.Lchild)
	}

	if root.Rchild != nil {
		preOderRecur(root.Rchild)
	}
}

func preOderRecur1(btree []string, i int) {
	btreeLen := len(btree)
	if i < btreeLen {
		fmt.Printf(btree[i] + " ")
	}

	if 2*i+1 < btreeLen {
		preOderRecur1(btree, 2*i+1)
	}

	if 2*i+2 < btreeLen {
		preOderRecur1(btree, 2*i+2)
	}
}

func inOrderRecur(root *Node) {
	if root.Lchild != nil {
		inOrderRecur(root.Lchild)
	}
	fmt.Printf(root.Data + " ")
	if root.Rchild != nil {
		inOrderRecur(root.Rchild)
	}
}

func posOrderRecur(root *Node) {
	if root.Lchild != nil {
		posOrderRecur(root.Lchild)
	}
	if root.Rchild != nil {
		posOrderRecur(root.Rchild)
	}
	fmt.Printf(root.Data + " ")
}
