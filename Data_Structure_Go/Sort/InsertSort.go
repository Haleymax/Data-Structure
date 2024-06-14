package main

import "fmt"

// 插入排序（从小到大）
func InsertSort(list []int) []int {
	n := len(list)
	for i := 1; i < n; i++ {
		deal := list[i] //待排序的数
		j := i - 1

		//向后扫描
		for j >= 0 && list[j] > deal {
			list[j+1] = list[j] //将大的元素向后移动找到待插入元素的正确位置
			j--
		}
		list[j+1] = deal

	}
	return list
}

func test03() {
	fmt.Println("插入排序的结果如下所示：")
	list := []int{5, 9, 1, 6, 8, 6, 12, 18, 16, 22}
	list = BubbleSort(list)
	fmt.Println(list)
	fmt.Println("--------------------------------")
}
