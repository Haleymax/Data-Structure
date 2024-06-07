package main

import (
	"fmt"
	"sync"
)

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

// 从尾部开始插入一个元素
func (list *DoubleList) AddNodeFromTail(n int, v string) {
	//加入并发锁保证安全
	list.lock.Lock()
	defer list.lock.Unlock()

	//索引超出范围
	if n > list.len {
		panic("index out of range")
	}

	//找出尾节点
	node := list.tail

	//向前遍历获取插入的位置
	for i := 0; i <= n; i++ {
		node = node.pre
	}

	//创建需要插入的新节点
	newNode := new(ListNode)
	newNode.value = v

	//如果插入的那个列表正好是表头
	if node.IsNil() {
		list.head = newNode
	} else {
		//保存后驱节点
		next := node.next

		//如果该后驱节点为空那么他也将是这个列表的队尾

		if next.IsNil() {

			//将新节点接入到旧节点的后面
			node.next = newNode
			newNode.pre = node
		} else {
			//插入新节点
			newNode.pre = node
			node.next = newNode

			newNode.next = next
			next.pre = newNode
		}
	}
	list.len++
}

// 从列表头部开始查找第n个节点
func (list *DoubleList) IndexFormHead(n int) *ListNode {

	//超出队列直接返回空指针
	if n >= list.len {
		return nil
	}

	//获取头部节点
	node := list.head

	for i := 0; i <= n; i++ {
		node = node.next
	}
	return node
}

// 从尾部开始往前查找第n个位置的节点
func (list *DoubleList) IndexFormTail(n int) *ListNode {
	//超出链表的长度返回空值
	if n >= list.len {
		return nil
	}

	//获取尾部节点
	node := list.tail

	//往前遍历获取前面第n个节点
	for i := 0; i <= n; i++ {
		node = node.pre
	}
	return node
}

// 从头部开始移除某个节点，并返回这个节点
func (list *DoubleList) PopFromHead(n int) *ListNode {
	//操作链表需要加入并发锁保证安全
	list.lock.Lock()
	defer list.lock.Unlock()

	//确保索引没有超出队列范围,若超出无法获取和删除返回一个空值
	if n >= list.len {
		return nil
	}

	//头部指针
	node := list.head

	//取到需要操作的那个节点位置
	for i := 0; i <= n; i++ {
		node = node.next
	}

	//取下移除节点的前驱指针和后驱指针防止断链
	pre := node.pre
	next := node.next

	//判断节点是否为列表中唯一的元素
	if pre.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	} else if pre.IsNil() {
		//移除的是头结点的话需要将头结点的下一个节点设置为头结点
		list.head = next
		next.pre = nil
	} else if next.IsNil() {
		//如果移除的是尾节点的话需要将尾节点的前一个节点设置成尾节点
		list.tail = pre
		pre.next = nil
	}

	list.len--
	return node
}

// 从尾部开始移除前面第n个元素
func (list *DoubleList) PopFromTail(n int) *ListNode {
	//对链表进行操作需要加上并发锁
	list.lock.Lock()
	defer list.lock.Unlock()

	if n >= list.len {
		return nil
	}

	//获取尾节点
	node := list.tail

	//向前操作找到需要插入的节点
	for i := 0; i <= n; i++ {
		node = node.pre
	}

	//操作节点的前驱和后继
	pre := node.pre
	next := node.next

	// 如果前驱和后驱都为nil，那么移除的节点为链表唯一节点
	if pre.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	} else if pre.IsNil() {
		// 表示移除的是头部节点，那么下一个节点成为头节点
		list.head = next
		next.pre = nil
	} else if next.IsNil() {
		// 表示移除的是尾部节点，那么上一个节点成为尾节点
		list.tail = pre
		pre.next = nil
	} else {
		// 移除的是中间节点
		pre.next = next
		next.pre = pre
	}
	// 节点减一
	list.len = list.len - 1
	return node

}

func test() {
	list := new(DoubleList)
	// 在列表头部插入新元素
	list.AddNodeFromHead(0, "I")
	list.AddNodeFromHead(0, "love")
	list.AddNodeFromHead(0, "you")
	// 在列表尾部插入新元素
	list.AddNodeFromTail(0, "may")
	list.AddNodeFromTail(0, "happy")
	list.AddNodeFromTail(list.Len()-1, "begin second")
	list.AddNodeFromHead(list.Len()-1, "end second")
	// 正常遍历，比较慢，因为内部会遍历拿到值返回
	for i := 0; i < list.Len(); i++ {
		// 从头部开始索引
		node := list.head
		// 节点为空不可能，因为list.Len()使得索引不会越界
		if !node.IsNil() {
			fmt.Println(node.GetValue())
		}
	}
	// 元素一个个 POP 出来
	for {
		node := list.PopFromHead(0)
		if node.IsNil() {
			// 没有元素了，直接返回
			break
		}
		fmt.Println(node.GetValue())
	}
	fmt.Println("----------")
	fmt.Println("len", list.Len())
}

func main() {
	test()
}
