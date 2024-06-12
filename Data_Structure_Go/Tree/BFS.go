package main

import (
	"fmt"
	"sync"
)

//二叉树的层次遍历（广度优先搜索）

// 二叉树结构体
type Treenode struct {
	Data  string    //数据域
	Left  *Treenode //左子树
	Right *Treenode //右子树
}

// 链表节点
type LinkNode struct {
	Data *Treenode //用于存储二叉树的节点
	Next *LinkNode
}

// 用链表实现队列，广度优先搜索依靠的就是队列
type LinkQueue struct {
	root *LinkNode //队列起点
	size int
	lock sync.Mutex //并发锁
}

// 入队
func (queue *LinkQueue) Add(v *Treenode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	//如果队列为空，那么新增加节点
	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Data = v
	} else {
		//将新加入的节点加入到队尾
		newNode := new(LinkNode)
		newNode.Data = v

		//获取队尾节点指针
		tail := queue.root
		for tail.Next != nil {
			tail = tail.Next
		}

		//将新节点放在链表尾部
		tail.Next = newNode
	}
	queue.size++
}

// 出队
func (queue *LinkQueue) Remove() *Treenode {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	//队列中元素为空
	if queue.size == 0 {
		panic("The queue is empty")
	}

	//顶部元素要出队
	Head := queue.root
	v := Head.Data
	queue.root = queue.root.Next

	queue.size--
	return v
}

// 层次遍历
func LayerOrder(treenode *Treenode) {
	if treenode == nil {
		panic("This tree is empty,please tray again")
	}

	//创建队列,层次遍历利用队列实现
	queue := new(LinkQueue)

	//先将根节点入队
	queue.Add(treenode)

	//直到队列为空
	for queue.size > 0 {
		elem := queue.Remove()

		fmt.Println(elem.Data)
		if elem.Left != nil {
			queue.Add(elem.Left)
		}
		if elem.Right != nil {
			queue.Add(elem.Right)
		}
	}
}

func main4() {
	t := &Treenode{Data: "A"}
	t.Left = &Treenode{Data: "B"}
	t.Right = &Treenode{Data: "C"}
	t.Left.Left = &Treenode{Data: "D"}
	t.Left.Right = &Treenode{Data: "E"}
	t.Right.Left = &Treenode{Data: "F"}
	fmt.Println("\n层次排序")
	LayerOrder(t)
}
