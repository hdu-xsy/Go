package main

import(
	"fmt"
	"../LinkedList"
)

func main() {
	var L = new(LinkedList.LNode)
	LinkedList.InitList(L)
	var a,i,e int
	for {
		fmt.Println("请输入功能1.插入 2.删除 3.查找 4.尾部添加 5.遍历")
		fmt.Scanln(&a)
		switch a {
		case 1:
			fmt.Println("插入的数字")
			fmt.Scanln(&e)
			fmt.Println("插入的位置")
			fmt.Scanln(&i)
			if LinkedList.ListInsert(L, i, e) {
				fmt.Println("插入成功")
			} else {
				fmt.Println("插入失败")
			}
		case 2:
			fmt.Println("删除的位置")
			fmt.Scanln(&i)
			if LinkedList.ListDelete(L, i) {
				fmt.Println("删除成功")
			} else {
				fmt.Println("删除失败")
			}
		case 3:
			fmt.Println("查找的位置")
			fmt.Scanln(&i)
			if LinkedList.GetElem(L, i, &e) {
				fmt.Println("查找的结果为", e)
			} else {
				fmt.Println("查找失败")
			}

		case 4:
			fmt.Println("插入的数字")
			fmt.Scanln(&e)
			LinkedList.ListAdd(L,e)
		case 5:
			LinkedList.GetList(L)
		}
	}
}