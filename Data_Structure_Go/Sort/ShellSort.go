package main

import "fmt"

// 简单的希尔排序实现（每一次排序的增量为上一次的一半）
func ShellSort(list []int) []int {
	n := len(list)

	//初始增量为数组长度的一半
	for step := n / 2; step >= 1; step /= 2 {

		//中间每一次排序采用插入排序的过程
		for i := step; i < n; i += step {
			for j := i - step; j >= 0; j -= step {
				if list[j+step] < list[j] {
					temp := list[j]
					list[j] = list[j+step]
					list[j+step] = temp
					continue
				}
				break
			}
		}
	}
	return list
}

func test04() {
	fmt.Println("希尔排序的结果如下所示：")
	list := []int{5, 9, 1, 6, 8, 6, 12, 18, 16, 22}
	list = BubbleSort(list)
	fmt.Println(list)
	fmt.Println("--------------------------------")
}
