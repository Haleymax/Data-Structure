package main

import "fmt"

// 切分
func partition(array []int, low int, high int) int {
	index := array[low] //将low指向的元素作为基准元素

	i := low
	for j := low + 1; j <= high; j++ {
		if array[j] <= index {
			i++
			temp := array[i]
			array[i] = array[j]
			array[j] = temp
		}
	}

	//确定基准元素的位置
	temp := array[low]
	array[low] = array[i]
	array[i] = temp
	return i
}

func QuickSort(array []int, low int, high int) {
	if low < high {

		//进行划分
		loc := partition(array, low, high)

		QuickSort(array, low, loc-1)  //对左边划分的区间进行排序
		QuickSort(array, loc+1, high) //对右边区间进行排序
	}
}

func QuickSorttest() {
	list := []int{5}
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)
	list1 := []int{5, 9}
	QuickSort(list1, 0, len(list1)-1)
	fmt.Println(list1)
	list2 := []int{5, 9, 1}
	QuickSort(list2, 0, len(list2)-1)
	fmt.Println(list2)
	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort(list3, 0, len(list3)-1)
	fmt.Println(list3)
}
