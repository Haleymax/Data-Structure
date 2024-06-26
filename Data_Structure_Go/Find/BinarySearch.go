package main

import "fmt"

func BinarySearch(slice []int, target int) int {
	left, right := 0, len(slice)-1

	for left <= right {
		mid := (left + right) / 2
		if slice[mid] == target {
			return mid
		} else if slice[mid] > target {
			right = mid - 1 //查找左侧区间
		} else {
			left = mid + 1 //查找右侧区间
		}
	}
	return -1 //没有找到
}

func binarySearch() {
	sortedSlice := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	target := 11
	index := BinarySearch(sortedSlice, target)

	if index != -1 {
		fmt.Printf("找到了目标值 %d，索引为 %d\n", target, index)
	} else {
		fmt.Printf("没有找到目标值 %d\n", target)
	}

	target = 20
	index = BinarySearch(sortedSlice, target)
	if index != -1 {
		fmt.Printf("找到了目标值 %d，索引为 %d\n", target, index)
	} else {
		fmt.Printf("没有找到目标值 %d\n", target)
	}
}
