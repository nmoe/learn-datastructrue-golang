package main

import (
	"fmt"
)

//ElemType replace int type
//type ElemType int

//LNode is linkedlist node struct
//结点结构包括数据域和指针域
type LNode struct {
	data int    //数据域
	next *LNode //next指针域
}

//InitList is initnial a list
//初始化单链表(生成头节点),头节点数据域任意也可为空，指针域暂且设置为nil
func InitList() *LNode {
	return &LNode{0, nil}
}

//GetElem function is
//查找第i个位置的元素,并返回该值
func (head *LNode) GetElem(i int) *LNode {
	
	p := head.next
	if p == nil {
		return &LNode{0, nil}
	}
	for j := 1; j < i; j++ {
		p = p.next
	}
	return p

}

//FindElem method 
//按照值查找是否存在
func (head *LNode) FindElem(e int) bool{
	p := head
	for p != nil {
		if p.data == e{
			return true
		}
		p = p.next
	}
	return false
}

//InsertElem function
//在指定位置i插入元素e
func (head *LNode) InsertElem(i, e int) bool {
	p := head
    j := 1
    for nil != p && j < i {
        p = p.next
        j++
    }
	
	s := &LNode{data: e}
	s.next = p.next
	p.next = s
	return true
}

//Traverse method
//遍历链表
func (head *LNode) Traverse() {
	pointer := head.next
	for pointer != nil{
		// fmt.Println("头节点的属性(包括data和next指针",head)
		// fmt.Println("当前指针的属性(包括data和next指针",pointer)
		//fmt.Println("当前指针的数据值",pointer.data)
		fmt.Println(pointer.data)
		pointer = pointer.next
		
	}
	fmt.Println("遍历完成")
}

//GetLength method
//求链表表长
func (head *LNode) GetLength() int{
	p := head
	i := 0
	for p !=nil{
		i++
		p = p.next
	}
	return i
}

//DeleteElem method 
//删除指定位置i的元素
func (head *LNode) DeleteElem(i int) bool{
	p := head
	j := 1
	for p != nil && j < i{   //找到第i个结点
		p = p.next   
		j++
	}
	if nil == p || j > i {
        fmt.Println("pls check i:", i)
        return false
	}
	p.next = p.next.next  //删除第i个结点,将被删除结点的后继节点,与前一个结点链接
	return true
}
func main(){
	list := InitList()
	list.InsertElem(1,99)
	list.InsertElem(2,89)
	list.InsertElem(3,999)
	list.InsertElem(4,789)
	list.InsertElem(5,456)
	list.InsertElem(6,123)
	list.Traverse()
	fmt.Println("第三个结点值:",list.GetElem(3).data)
	list.DeleteElem(4)
	fmt.Println("删除指定位置结点后的链表:")
	list.Traverse()
	fmt.Println(list.FindElem(123))
	fmt.Println("该链表的长度为:",list.GetLength())
}
