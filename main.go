package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

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

//增加结点数据
func addNode(data int, link *DoubleLink) {
	hash := myHash(data)
	node := &Node{}
	node.Data = data
	//散列表添加，isExist 判断此数据是否已经存在于缓存链表中
	isExist := false
	if table[hash] == nil {
		table[hash] = node
	} else {
		//遍历散列表此位置上对应的链表
		tempNode := table[hash]
		for {
			if tempNode == nil {
				break
			}
			if tempNode.Data == data { //数据已经存在就不用再次添加了
				isExist = true
				node = tempNode
				break
			}
			if tempNode.hnext == nil {
				tempNode.hnext = node
				break
			} else {
				tempNode = tempNode.hnext
			}
		}
	}
	//双向链表缓存分已存在数据添加和直接添加两种情况
	if isExist {
		//不是排在链表最后一个，就删除当前位置，添加到链尾去
		if node.Next != nil {
			fmt.Println("数据" + strconv.Itoa(node.Data) + "已存在")
			prevNode := node.Prev
			nextNode := node.Next
			//链头位置
			if link.Head == node {
				link.Head = nextNode
				nextNode.Prev = nil
			} else {
				prevNode.Next = nextNode
				nextNode.Prev = prevNode
			}
			node.Next = nil
			node.Prev = nil

			link.Tail.Next = node
			fmt.Println("老链尾：" + strconv.Itoa(link.Tail.Data))
			node.Prev = link.Tail
			node.Next = nil
			link.Tail = node
			fmt.Println("新链尾：" + strconv.Itoa(node.Data))

			//oldTailNode:=link.Tail
			//oldTailNode.Next=node
			//fmt.Println("老链尾："+strconv.Itoa(oldTailNode.Data))
			//node.Prev=oldTailNode
			//node.Next=nil
			//link.Tail=node
			//fmt.Println("新链尾："+strconv.Itoa(node.Data))
		}
	} else {
		//双向链表链头添加
		if link.NodeNum == 0 {
			link.Head = node
			link.Tail = node
			link.NodeNum++
			return
		} else if link.NodeNum < 10 && link.NodeNum > 0 { //链尾直接添加
			oldTail := link.Tail
			oldTail.Next = node
			node.Prev = oldTail
			link.Tail = node
			link.NodeNum++
			return
		} else { //满了先删最老的信息，即链表头结点，再添加最新数据到尾结点
			oldhead := link.Head
			newHead := link.Head.Next
			oldhead.Next = nil
			newHead.Prev = nil
			link.Head = newHead
			deleteData := oldhead.Data
			deleteIndex := myHash(deleteData) //散列表同时删除最老结点
			tempNode := table[deleteIndex]
			if tempNode.Data == deleteData { //散列表中的链表第一个结点就是要删除的结点
				table[deleteIndex] = tempNode.hnext

			} else {
				for {
					if tempNode.hnext.Data == deleteData {
						tempNode.hnext = tempNode.hnext.hnext
						break
					} else {
						tempNode = tempNode.hnext
					}
				}
			}
			link.Tail.Next = node
			node.Prev = link.Tail
			node.Next = nil
			link.Tail = node
		}
	}
}

//遍历双向链表，即缓存内容
func iterLink(link *DoubleLink) {
	tempNode := link.Head
	fmt.Println("缓存遍历开始，总数：" + strconv.Itoa(link.NodeNum))
	for {
		fmt.Print(strconv.Itoa(tempNode.Data) + ",")
		if tempNode.Next == nil {
			break
		} else {
			tempNode = tempNode.Next
		}
	}
}

//遍历散列表内容
func iterTable() {
	for i := 0; i < 10; i++ {
		tempNode := table[i]
		fmt.Println("\r\n第" + strconv.Itoa(i) + "行：")
		for {
			if tempNode == nil {
				break
			} else {
				fmt.Print(strconv.Itoa(tempNode.Data) + ";")
				tempNode = tempNode.hnext
			}
		}
	}
}

var table [10]*Node

func main() {
	//大小为10的数组代替散列表
	link := &DoubleLink{0, nil, nil}
	for i := 1; i <= 250; i++ {
		j := rand.Intn(20)
		fmt.Println("随机数：" + strconv.Itoa(j))
		addNode(j, link)
		iterLink(link)
	}
	iterTable()
}
