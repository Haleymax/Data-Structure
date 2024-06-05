package main

import (
	"fmt"
	"sync"
)

// 链表栈
type LinkStack struct {
	root *LinkNode //栈顶
	size int
	lock sync.Mutex //保护并发读写数据
}

type LinkNode struct {
	Next  *LinkNode //指向下一个节点
	Value string    //栈节点的元素
}

// 入栈
func (stack *LinkStack) Push(v string) {
	stack.lock.Lock() //加锁
	defer stack.lock.Unlock()

	//创建一个新的节点
	newNode := &LinkNode{Value: v}

	//将新节点插入到栈顶
	newNode.Next = stack.root
	stack.root = newNode

	stack.size++
}

// 出栈
func (stack *LinkStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	//栈为空的时候
	if stack.root == nil {
		panic("stack is empty")
	}

	//获取栈顶元素
	v := stack.root.Value

	//栈顶指针向后移动
	stack.root = stack.root.Next

	stack.size--

	return v
}

// 返回栈顶元素
func (stack *LinkStack) Peek() string {
	return stack.root.Value
}

// 获取栈的大小
func (stack *LinkStack) Size() int {
	return stack.size
}

// 判空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}

func test3() {
	// 创建一个新的链表栈
	linkStack := &LinkStack{}

	// 判空
	fmt.Println("Is stack empty?", linkStack.IsEmpty()) // 应该输出 true

	// 入栈
	linkStack.Push("apple")
	linkStack.Push("banana")
	linkStack.Push("cherry")

	// 判空
	fmt.Println("Is stack empty?", linkStack.IsEmpty()) // 应该输出 false

	// 出栈
	fmt.Println("Pop:", linkStack.Pop()) // 应该输出 cherry
	fmt.Println("Pop:", linkStack.Pop()) // 应该输出 banana

	// 入栈
	linkStack.Push("date")
	linkStack.Push("elderberry")

	// 出栈
	fmt.Println("Pop:", linkStack.Pop()) // 应该输出 elderberry

	// 获取栈顶元素
	fmt.Println("Peek:", linkStack.Peek()) // 应该输出 date

	// 获取栈大小
	fmt.Println("Stack size:", linkStack.Size()) // 应该输出 2
}

func main() {
	test3()
}
