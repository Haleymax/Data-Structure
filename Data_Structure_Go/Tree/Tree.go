package main

import "fmt"

// 二叉树
type TreeNode struct {
	Data  string    //节点数据域
	Left  *TreeNode //左子树
	Right *TreeNode //右子树
}

// 先序遍历
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	fmt.Println(tree.Data)
	//先序遍历是将左子树遍历完之后再遍历右子树
	PreOrder(tree.Left)
	PreOrder(tree.Right)
}

// 中序遍历
func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	//中序遍历是先遍历左子树，遍历完再遍历自己的节点，最后再遍历右子树
	MidOrder(tree.Left)
	fmt.Println(tree.Data)
	MidOrder(tree.Right)
}

// 后续遍历
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	//后序遍历，将左右子树都遍历完了之后再读自己
	MidOrder(tree.Left)
	MidOrder(tree.Right)
	fmt.Println(tree.Data)
}

func main9() {
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}
	t.Right.Left = &TreeNode{Data: "D"}
	t.Right.Right = &TreeNode{Data: "E"}
	t.Right.Left.Left = &TreeNode{Data: "F"}
	t.Right.Left.Right = &TreeNode{Data: "G"}

	fmt.Println("先序遍历：")
	PreOrder(t)
	fmt.Println("中序遍历：")
	MidOrder(t)
	fmt.Println("后序遍历：")
	PostOrder(t)

}
