package main

import "sync"

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
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()

	s.m[item] = struct{}{} //向字典中加入这个键
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
