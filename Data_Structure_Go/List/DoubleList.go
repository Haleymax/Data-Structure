package main

import "sync"

// 双端列表，双端队列
type DoubleList struct {
	head *ListNode  //指向链表头部
	tail *ListNode  //指向链表尾部
	len  int        //列表长度
	lock sync.Mutex //保证读写安全的并发锁
}

// 列表中的节点
type ListNode struct {
	pre   *ListNode //前驱节点
	next  *ListNode //后驱节点
	value string    //节点的值
}

// 获取节点的值
func (node *ListNode) GetValue() string {
	return node.value
}

// 获取节点前驱节点
func (node *ListNode) GetPre() *ListNode {
	return node.pre
}

// 获取节点的后驱节点
func (node *ListNode) GetNext() *ListNode {
	return node.next
}

// 判断是否存在后驱节点
func (node *ListNode) HasNext() bool {
	if node.next == nil {
		return false
	} else {
		return true
	}
}

// 判断列表是否存在前驱节点
func (node *ListNode) HasPre() bool {
	if node.pre == nil {
		return false
	} else {
		return true
	}
}

// 判断节点是否为空节点
func (node *ListNode) IsNil() bool {
	if node == nil {
		return false
	} else {
		return true
	}
}

// 返回列表长度
func (node *DoubleList) Len() int {
	return node.len
}

// 获取双端列表中第n个节点的值
func (node *DoubleList) Getvalue(n int) string {
	var r *ListNode
	r = node.head
	for i := 0; i < n; i++ {
		r = r.next
	}
	return r.value

}

// 插入元素，插入到双端列表的第n个位置
func (list *DoubleList) AddNodeFromHead(n int, v string) {
	//加上并发锁
	list.lock.Lock()
	defer list.lock.Unlock()

	//超出范围
	if n > list.len {
		panic("index out of range")
	}

	//创建工作指针
	node := list.head

	//定位到第n个节点
	for i := 1; i < n; i++ {
		node = node.next
	}

	//创建一个新的节点
	newNode := new(ListNode)
	newNode.value = v

	//如果列表为空则插入的节点为第一个节点
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		//检查前驱节点是否为空
		pre := node.pre
		if pre.IsNil() {
			//前驱节点为空那么新插入的节点就是双端列表的第一个节点
			newNode.next = node
			node.pre = newNode
			list.head = newNode
		} else {
			//处于中间节点，插入新加入的节点
			pre.next = newNode
			newNode.pre = pre
			node.pre = newNode
			newNode.next = node
		}
	}
	list.len++

}
