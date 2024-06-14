package main

import "fmt"

// 归并排序自顶向下的实现
func MergeSort(list []int, start, end int) []int {
	if start >= end {
		//如果只有一个元素的话直接返回
		return list[start : end+1]
	}

	mid := (start + end) / 2

	left := MergeSort(list, start, mid)
	right := MergeSort(list, mid, end)
	Merge(left, right)

	return list
}

// 合并分出来的两个数组
func Merge(left, right []int) []int {
	var result []int
	i, j := 0, 0 //表示左边数组的索引，j表示右边数组的指针

	//将两个数组的元素合并到result数组中
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i]) //从小到大的顺序排序
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	//合并剩余的元素,剩下的元素一定是更大的元素
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func test05() {
	fmt.Println("归并排序的结果如下所示：")
	list := []int{5, 9, 1, 6, 8, 6, 12, 18, 16, 22}
	list = BubbleSort(list)
	fmt.Println(list)
	fmt.Println("--------------------------------")
}
