package main

import "fmt"

// 堆排序的实现（最大堆，顶层的节点值更大）
type Heap struct {
	Size  int   // 堆的大小
	Array []int // 使用静态数组来实现二叉树
}

// 初始化堆
func NewHeap(arr []int) *Heap {
	h := new(Heap)
	h.Array = make([]int, len(arr))
	h.Size = len(arr)
	return h
}

// 插入元素
func (h *Heap) Push(x int) {
	if h.Size >= len(h.Array) {
		// 扩展数组长度
		h.Array = append(h.Array, x)
	} else {
		h.Array[h.Size] = x
	}

	i := h.Size // 需要插入的索引位置
	for i > 0 { // 找到x的位置
		parent := (i - 1) / 2 // 父节点的位置
		if x <= h.Array[parent] {
			break // 如果比父节点大则满足大堆的条件直接可以插入
		} else {
			h.Array[i] = h.Array[parent] // 如果比父节点大，则需要与父节点交换使其满足大堆的条件
			i = parent                   // 一直进行，直到满足这个节点的父节点比他大为止或者成为根节点
		}
	}

	h.Array[i] = x
	h.Size++
}

// 移除大堆中的元素
func (h *Heap) Pop() int {
	if h.Size == 0 {
		return -1 // 树为空返回-1表示不能够继续移除了
	}

	ret := h.Array[0] // 每次取的元素都是根节点即为最大的元素

	h.Size--
	x := h.Array[h.Size] // 将最后个节点取出放入根节点使其自动进行匹配以满足大堆的条件
	h.Array[0] = x

	i := 0
	for {
		a := 2*i + 1 // 左子树
		b := 2*i + 2 // 右子树

		if a >= h.Size {
			break // 若是超出堆的大小表示不存在后面的子树了
		}

		if b < h.Size && h.Array[b] > h.Array[a] {
			a = b
		}

		if x >= h.Array[a] {
			break // 如果这个节点已经比子树都大了那么已经满足条件了
		}

		h.Array[i] = h.Array[a] // 将大的元素往上换
		i = a                   // 进行交换直达找到左右子树都比自己小或者到达末尾
	}

	h.Array[i] = x // 确定刚取的尾元素的位置
	return ret
}

func HeapSort() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}

	// 构建最大堆
	h := NewHeap(make([]int, len(list)))
	for _, v := range list {
		h.Push(v)
	}

	// 所有元素都取出后的顺序就是一个从大到小的序列（每次取出的元素都是放在列表的最后面）
	sorted := make([]int, len(list))
	for i := range sorted {
		sorted[i] = h.Pop()
	}

	// 打印排序后的值
	fmt.Println(sorted)
}
