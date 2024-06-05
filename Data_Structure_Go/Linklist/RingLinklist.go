package main

import "fmt"

// 这是一个循环链表的实现
type Ring struct {
	next, prev *Ring       //分别代表指向前驱和后继的指针
	Value      interface{} //数据域
}

// 初始化循环列表
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// 创建N个节点的循环链表
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r

	for i := 0; i < n; i++ {
		p.next = &Ring{prev: p} //创建新的结构体
		p = p.next
	}
	p.next = r
	r.prev = p
	return r

}

// 获取下一个节点
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init() //循环链表中的下一个节点若为空表示不合法，如果不合法就初始化这个链表
	}
	return r.next
}

// 获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

// 获取第n个节点
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// 添加新的节点，返回被连接的节点的后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

func linkNewTest() {
	// 第一个节点
	r := &Ring{Value: -1}
	// 链接新的五个节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})
	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)
		// 移到下一个节点
		node = node.Next()
		//  如果节点回到了起点，结束
		if node == r {
			return
		}
	}
}

func main3() {
	linkNewTest()
}
