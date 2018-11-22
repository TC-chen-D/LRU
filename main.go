package main

import "fmt"

//LRU 淘汰算法
//散列表+双向链表
//前驱和后继指针是为了将结点串在双向链表中，hnext 指针是为了将结点串在散列表的拉链中

//双向节点+散列表的拉链节点
type Node struct {
	Prev  *Node //前驱指针
	Data  int   //数据
	Next  *Node //后继指针h
	hnext *Node //散列表的链表的指针
}

//双向链表，头结点，尾结点,缓存大小由结点数决定，这里定最大15个
type DoubleLink struct {
	NodeNum int
	Head    *Node
	Tail    *Node
}

//模拟hash算法，这里只做除余计算，为了散列表的链表都能有多个数据
func myHash(data int) int {
	return data % 10
}

//增加结点
func addNode(data int, link DoubleLink) {
	hash := myHash(data)
	node := &Node{}
	node.Data = data
	if table[hash] == nil {
		table[hash] = node
	} else {
		//遍历散列表此位置上对应的链表
		for {
			tempNode := table[hash]
			fmt.Println(tempNode)
		}
	}

}

var table [10]*Node

func main() {
	//大小为10的数组代替散列表

}
