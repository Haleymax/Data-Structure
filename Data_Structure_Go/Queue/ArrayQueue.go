package main

import (
	"fmt"
	"sync"
)

// 用数组实现队列
type ArrayQueue struct {
	array []string
	size  int
	front int        //队头索引
	rear  int        //队尾索引
	lock  sync.Mutex //并发锁保证读写数据的安全
}

// 初始化队列(需要传入队列的大小)
func (queue *ArrayQueue) Init(capacity int) {
	queue.array = make([]string, capacity)
	queue.size = 0
	queue.front = 0
	queue.rear = 0

}

// 入队
func (queue *ArrayQueue) Enqueue(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 检查队列是否已满
	if queue.rear == len(queue.array)-1 {
		panic("Queue is full")
	}

	//将元素放入队列中
	queue.rear++ //队尾入队
	queue.array[queue.rear] = v
	queue.size++
}

// 出队
func (queue *ArrayQueue) Dequeue() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.front == queue.rear {
		panic("Queue is empty")
	}

	v := queue.array[queue.front] //从对头出队
	queue.front++
	queue.size--

	return v
}

// 打印队列
func (queue *ArrayQueue) Show() {
	for i := queue.front; i <= queue.front+queue.size; i++ {
		fmt.Println(queue.array[i])
	}
}

func test() {
	var q ArrayQueue
	q.Init(10)
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")
	q.Enqueue("d")
	//q.Show()

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Show()
}

func main2() {
	test()
}
