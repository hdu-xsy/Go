package LinkedList

import "fmt"

type LNode struct {
	data int
	Next *LNode
}

type LinkList *LNode

func InitList(L LinkList) {
	L = new(LNode)
	if L == nil {
		return
	}
}

func ListEmpty(L LinkList) bool {
	return L != nil && L.Next == nil
}

func ListLength(L LinkList) int {
	var i = 0
	var p LinkList
	if L != nil {
		p = L.Next
		for{
			if p == nil {
				break
			}
			i++
			p = p.Next
		}
	}
	return i
}

func GetElem(L LinkList,i int,e *int) bool{
	var j = 1
	var p LinkList = L.Next
	for {
		if !(p != nil && j<i) {
			break
		}
		j++
		p = p.Next
	}
	if p == nil || j>i {
		return false
	}
	*e = p.data
	return true
}

func ListInsert(L LinkList,i int,e int) bool{
	var p,s LinkList
	var j int
	p = L
	j = 0
	for {
		if !(p != nil && j<i-1) {
			break
		}
		p = p.Next
		j++
	}
	if p == nil || j > i-1 {
		return false
	}
	s = new(LNode)
	if s == nil {
		return false
	}
	s.data = e
	s.Next = p.Next
	p.Next = s
	return true
}

func ListDelete(L LinkList,i int) bool{
	var pre,p LinkList
	var j int
	pre = L
	j = 1
	for {
		if !(pre.Next !=nil && j<i) {
			break
		}
		pre = pre.Next
		j++
	}
	if pre.Next == nil || j > i {
		return false
	}
	p = pre.Next
	pre.Next = p.Next
	return true
}

func ListAdd(L LinkList,e int) {
	for {
		if L.Next == nil {
			break
		}
		L = L.Next
	}
	var p = new(LNode)
	p.data = e
	L.Next = p
}

func GetList(L LinkList) {
	for {
		if L.Next == nil {
			break
		}
		L = L.Next
		fmt.Print(L.data, " ")
	}
	fmt.Println()
}