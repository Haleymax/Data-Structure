package main

import "fmt"

// 冒泡排序
func BubbleSort(list []int) []int {
	n := len(list)

	for i := 0; i < n-1; i++ {

		flag := false //如果一个循环都没有交换一个数据那么已经完成了排序

		//冒泡排序的每一次循序都会确定最大或最小的一个元素
		for j := 0; j < n-i-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp

				flag = true
			}
		}

		if flag == false {
			return list
		}
	}
	return list
}

// 测试
func test01() {
	fmt.Println("冒泡排序的结果如下所示：")
	list := []int{5, 9, 1, 6, 8, 6, 12, 18, 16, 22}
	list = BubbleSort(list)
	fmt.Println(list)
	fmt.Println("--------------------------------")
}
