package main

import (
	"fmt"
	"sync"
)

// 使用链表实现队列
type LinkQueue struct {
	front *LinkNode //队头指针
	rear  *LinkNode //队尾指针
	size  int       //队列大小
	lock  sync.Mutex
}

// 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value string
}

// 初始化队列
func (q *LinkQueue) Init() {
	q.front = nil
	q.rear = nil
	q.size = 0
}

// 入队
func (q *LinkQueue) Enqueue(value string) {
	q.lock.Lock()
	defer q.lock.Unlock()

	//创建新加入的节点
	NewNode := &LinkNode{Value: value}
	NewNode.Next = nil

	//判断队列是否为空
	if q.front == nil && q.rear == nil {
		q.front = NewNode
		q.rear = NewNode

	} else {
		q.rear.Next = NewNode
		q.rear = q.rear.Next
	}
	q.size++
}

// 出队
func (q *LinkQueue) Dequeue() string {
	q.lock.Lock()
	defer q.lock.Unlock()

	//判断队列是否为空
	if q.size == 0 {
		panic("Queue is empty")
	} else {
		value := q.front.Value
		q.front = q.front.Next
		q.size--
		//在golang语言中不需要手动释放节点，如果没有再被使用了，系统会自动释放

		if q.size == 0 {
			q.front = nil
			q.rear = nil
		}
		return value
	}
}

// 打印函数
func (q *LinkQueue) Show() {
	for n := q.front; n != nil; n = n.Next {
		fmt.Println(n.Value)
	}
}

func test2() {
	var q LinkQueue
	q.Init()
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")
	q.Enqueue("d")
	//q.Show()

	q.Dequeue()
	q.Dequeue()
	q.Show()

}
func main3() {
	test2()
}
