package main

import (
	"fmt"
	"sync"
)

//可变长数组

//在golang语言中可变长数组就是切片  slice

func test01() {
	//创建一个容量为2的切片
	array := make([]int, 0, 2)
	fmt.Println("cap", cap(array), "len", len(array), "array", array)

	//虽然append但是没有赋予原来的变量array
	//在 Go 语言中，append 函数会返回一个新的切片，而不会修改原始切片的容量
	_ = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array", array)
	_ = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array", array)
	_ = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array", array)

	fmt.Println("----------------------")

	//赋予回原来的变量
	array = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array", array)
	array = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array", array)
	array = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array", array)
	array = append(array, 1, 1, 1, 1, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array", array)
	array = append(array, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array", array)
}

// 创建一个 len 个元素，容量为 cap 的可变长数组
type Array struct {
	array []int      //固定大小的数组，用满容量和满大小的切片来替代
	len   int        //真正长度
	cap   int        //容量
	lock  sync.Mutex //为了并发安全使用的锁
}

// 新建一个可变长数组
func Make(len, cap int) *Array {
	s := new(Array)
	if len > cap {
		panic("len larger than cap")
	}

	//把切片当数组用
	array := make([]int, cap, cap)

	//元数据
	s.array = array
	s.cap = cap
	s.len = len
	return s
}

// Append 增加一个元素
func (a *Array) Append(element int) {
	//并发锁
	a.lock.Lock()
	defer a.lock.Unlock()

	//大小等于容量，表示没多余位置了
	if a.len == a.cap {
		//没容量了数组要扩容
		newCap := 2 * a.len

		//需要判断之前的容量，如果之前就是0新容量为1
		if a.cap == 0 {
			newCap = 1
		}

		newArray := make([]int, newCap, newCap)

		//把老数组的数据移动到新数组
		for k, v := range a.array {
			newArray[k] = v
		}

		//替换数组
		a.array = newArray
		a.cap = newCap
	}

	//把元素放在数组里
	a.array[a.len] = element
	//真实长度+1
	a.len++
}

// AppendMany 增加多个元素
func (a *Array) AppendMany(element ...int) {
	for _, v := range element {
		a.Append(v)
	}
}

// Get获取某个下标的元素
func (a *Array) Get(index int) int {
	//越界了
	if a.len == 0 || index > a.len {
		panic("index out of range")
	}
	return a.array[index]
}

// Len返回长度
func (a *Array) Len() int {
	return a.len
}

// Cap返回容量
func (a *Array) Cap() int {
	return a.cap
}

// Print 辅助打印
func Print(array *Array) (result string) {
	result = "["
	for i := 0; i < array.Len(); i++ {
		// 第一个元素
		if i == 0 {
			result = fmt.Sprintf("%s%d", result, array.Get(i))
			continue
		}
		result = fmt.Sprintf("%s %d", result, array.Get(i))
	}
	result = result + "]"
	return
}

func test02() {
	// 创建一个容量为3的动态数组
	a := Make(0, 3)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
	// 增加一个元素
	a.Append(10)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
	// 增加一个元素
	a.Append(9)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
	// 增加多个元素
	a.AppendMany(8, 7)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))

	fmt.Println(a.Get(2))
}

func main4() {
	//test01()
	test02()
}
