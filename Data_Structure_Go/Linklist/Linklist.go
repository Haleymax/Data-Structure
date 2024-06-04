package main

import "fmt"

//这是一个最简单的链表实现

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func main() {
	//创建一个新的节点
	node := new(LinkNode)
	node.Data = 2

	//创建第二个节点
	node1 := new(LinkNode)
	node1.Data = 3
	node.NextNode = node1 //将第二个节点连接到第一个节点上

	//创建第三个节点
	node2 := new(LinkNode)
	node2.Data = 4
	node1.NextNode = node2 //将第三个节点连接到第二个节点上

	//输出链表的三个节点
	nowNode := node

	for {
		if nowNode != nil {
			fmt.Println(nowNode.Data)
			nowNode = nowNode.NextNode
			continue
		}
		break
	}

}
