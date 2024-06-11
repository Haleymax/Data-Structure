package main

import (
	"fmt"
	"sync"
)

// 集合
type Set struct {
	m            map[int]struct{} //采用字典来实现
	len          int              //集合的大小
	sync.RWMutex                  //加入锁保证读写的安全
}

// 创建一个空集合
func NewSet(cap int64) *Set {
	temp := make(map[int]struct{}, cap)
	return &Set{
		m: temp, //将这个m指定为新创建的字典
	}
}

// 添加一个新的元素
func (s *Set) Add(key int) {
	s.Lock()
	defer s.Unlock()

	s.m[key] = struct{}{} //向字典中加入这个键
	s.len = len(s.m)
}

// 移除一个元素
func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()

	//集合没元素直接返回
	if s.len == 0 {
		return
	}

	delete(s.m, item) //从字典中删除这个键
	s.len = len(s.m)
}

// 检查元素是否存在
func (s *Set) Has(e int) bool {
	s.RLock()
	defer s.RUnlock()

	_, ret := s.m[e]
	return ret
}

// 检查集合大小
func (s *Set) Len() int {
	return s.len
}

// 检查集合是否为空
func (s *Set) IsEmpty() bool {
	if s.len == 0 {
		return true
	}
	return false
}

// 清除集合所有元素
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()

	s.m = map[int]struct{}{}
	s.len = 0
}

// 将集合转换为列表
func (s *Set) ToSlice() []int {
	s.RLock()
	defer s.RUnlock()

	list := make([]int, 0, s.len)
	for key := range s.m {
		list = append(list, key)
	}
	return list
}

func main4() {
	//创建一个集合
	s := NewSet(10)

	s.Add(1)
	s.Add(2)
	s.Add(3)
	fmt.Println(s.m)

	s.Remove(2)
	fmt.Println(s.m)

	if s.IsEmpty() {
		fmt.Println("empty set")
	}

	if s.Has(2) {
		fmt.Println("2 exists")
	} else {
		fmt.Println("2 not exists")
	}

	fmt.Println(s.ToSlice())
}
