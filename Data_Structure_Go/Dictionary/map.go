package main

import "fmt"

// golang语言自带的字典
func main1() {
	//创建一个容量为4的字典map
	m := make(map[string]int64, 4)

	//放入三个键值对
	m["dog"] = 1
	m["cat"] = 2
	m["bird"] = 3

	fmt.Println(m)

	//查找 bird
	which := "bird"

	v, ok := m[which]

	if ok {
		//找到了
		fmt.Println("find:", which, "value:", v)
	} else {
		//没有找到
		fmt.Println("not found:", which)
	}

	//查找 cow
	which = "cow"

	v, ok = m[which]

	if ok {
		//找到了
		fmt.Println("find:", which, "value:", v)
	} else {
		//没有找到
		fmt.Println("not found:", which)
	}

}
