package main

// AVL树
type AVLTree struct {
	Root *AVLTreeNode //根节点
}

// AVL节点
type AVLTreeNode struct {
	Value  int64        //值
	Times  int64        //出现次数
	Height int64        //该节点作为树根节点，树的高度，用于计算平衡因子
	Left   *AVLTreeNode //左子树
	Right  *AVLTreeNode //右子树
}

// 初始化一个AVL树
func NewAVLTree() *AVLTree {
	return new(AVLTree)
}

// 更新节点树的高度
func (node *AVLTreeNode) UpdateHeight() {
	if node == nil {
		return
	}

	var leftHeight, rightHeight int64 = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}

	if node.Right != nil {
		rightHeight = node.Right.Height
	}

	//将最高的一边作为该树的树高
	maxHeight := leftHeight

	if rightHeight > leftHeight {
		maxHeight = rightHeight
	}

	node.Height = maxHeight + 1 //自身需要加一层
}

// 计算平衡因子
func (node *AVLTreeNode) BalanceFactor() int64 {
	var leftHeight, rightHeight int64 = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}
	return leftHeight - rightHeight
}
