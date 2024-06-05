package main

import (
	"fmt"
	"sync"
)

// 数组栈
type ArrayStack struct {
	array []string //底层切片
	top   int      //栈顶指针
	lock  sync.Mutex
}

// 入栈
func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	//放入切片中，后进的元素放在
	stack.array = append(stack.array, v)
	stack.top++
}

// 出栈
func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 栈顶中元素为空
	if stack.top == -1 {
		panic("empty stack")
	}

	// 获取栈顶元素
	v := stack.array[stack.top]
	stack.top--

	// 创建新的数组，从原数组中复制元素
	newArray := make([]string, stack.top+1, stack.top+1)
	copy(newArray, stack.array)
	stack.array = newArray

	return v
}

// 获取栈顶元素
func (stack *ArrayStack) Peek() string {
	//栈顶以为空
	if stack.top == -1 {
		panic("empty stack")
	}

	//返回栈顶元素
	return stack.array[stack.top]
}

// 获取栈的大小
func (stack *ArrayStack) Size() int {
	return stack.top
}

// 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	if stack.top == -1 {
		return true
	}
	return false
}

func test() {
	arrayStack := new(ArrayStack)
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")
	fmt.Println("size:", arrayStack.Size())
	fmt.Println("size:", arrayStack.Size())
	arrayStack.Push("drag")
	fmt.Println("pop:", arrayStack.Pop())
}

func main() {
	test()
}
