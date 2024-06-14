package main

import "fmt"

// 选择排序（从小到大）
func SelectSort(list []int) []int {
	n := len(list)

	for i := 0; i < n-1; i++ {

		//取每一轮开始的数为最小值若在遍历中找打比他小将最小值转让
		min := list[i]
		minIndex := i //记录最小值的索引

		//每一次遍历可以确定一个最大值和最小值
		for j := i; j < n; j++ {
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}

		//最小值发生了变化
		if minIndex != i {
			temp := list[i]
			list[i] = list[minIndex]
			list[minIndex] = temp
		}

	}
	return list
}

func test02() {
	fmt.Println("选择排序的结果如下所示：")
	list := []int{5, 9, 1, 6, 8, 6, 12, 18, 16, 22}
	list = BubbleSort(list)
	fmt.Println(list)
	fmt.Println("--------------------------------")
}
