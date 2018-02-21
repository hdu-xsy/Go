package ArrayList

import (
	"fmt"
	"os"
)

type LNode struct {
	elem []int
	length int
	listsize int
}

type ArrayList *LNode

func InitList(L ArrayList) bool{
	L.elem = make([]int,10)
	if L.elem == nil {
		os.Exit(-1)
	}
	L.length = 0
	L.listsize	= 10
	return true
}

func ListEmpty(L ArrayList) bool {
	return L.length == 0
}

func ListLength(L ArrayList) int {
	return L.length
}

func GetElem(L ArrayList,i int,e *int) bool{
	if i<1|| i>L.length {
		return false
	} else {
		*e = L.elem[i-1]
	}
	return true
}

func ListInsert(L ArrayList,i int,e int) bool{
	var newbase []int
	if i<1|| i>L.length+1 {
		return false
	}
	if L.length >= L.listsize {
		newbase = make([]int,L.listsize+5)
		copy(newbase,L.elem)
		if newbase == nil {
			os.Exit(-1)
		}
		L.elem = newbase
		L.listsize += 5
	}
	for a :=L.length-1;a>i-2;a=a-1 {
		L.elem[a+1] = L.elem[a]
	}
	L.elem[i-1] = e
	L.length++
	return true
}

func ListDelete(L ArrayList,i int) bool{
	if i<1 || i >L.length {
		return false
	}
	for a:=i-1;a<L.length-1;a++ {
		L.elem[a] = L.elem[a+1]
	}
	L.length--
	return true
}

func ListAdd(L ArrayList,e int) {
	ListInsert(L,L.length+1,e)
}

func GetList(L ArrayList) {
	for i:=0;i<L.length;i++ {
		fmt.Print(L.elem[i]," ")
	}
	fmt.Println()
}