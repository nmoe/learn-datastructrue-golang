package main

import "fmt"

// SqList  is a sqlist file
//定义顺序表
type SqList struct{
	//定义顺序表的最大长度
	maxSize int
	//定义当前顺序表的长度
	length int
	//顺序表的元素
	data []int
}

//InitSqList is to initmylist
//初始化顺序表
func InitSqList(maxSize int) *SqList{
	return &SqList{maxSize: maxSize, data: make([]int, maxSize)}
}

//IsEmpty is or not empty
//检查顺表是否为空
//我选择的方法是查看当前顺序表的长度是否为0
func (mySqList *SqList) IsEmpty() bool{
	return mySqList.length == 0
}

//IsFull is or not full
//判断当前顺序表长度是否已达到最大值
func (mySqList *SqList) IsFull() bool{
	return mySqList.length == mySqList.maxSize
}

//ListInsert function is to insert an element
//在i位置插入元素e
func (mySqList *SqList) ListInsert(i, e int) bool{
	if i < 1 || i > mySqList.length + 1 {
		fmt.Println("不合法,请检查i")
		return false
	}
	if i > mySqList.maxSize {
		fmt.Println("不合法,请检查i")
		return false
	}
	for j := mySqList.length; j >= i; j--{
		mySqList.data[j] = mySqList.data[j-1]
	}
	mySqList.data[i-1] = e
	mySqList.length++
	return true
}

//ListAppend function is to append an element
//在尾部追加一个元素,需要和最大长度比较
func (mySqList *SqList) ListAppend(e int) bool{
	if mySqList.IsFull() {
		fmt.Println("sorry list is maxsize")
		return false
	}
	mySqList.data[mySqList.length] = e
	mySqList.length++
	return true
}

//ListDelete function is to remove en element
//此方法是用来删除指定位置元素 
//method 1
func (mySqList *SqList) ListDelete(i int) bool{
	if i < 1 || i > mySqList.length + 1 {
		fmt.Println("i不合法,请检查i")
		return false
	}
	if i > mySqList.maxSize {
		fmt.Println("i不合法,请检查i")
	}
	for j := i-1; j < mySqList.length-1; j++{
		mySqList.data[j] = mySqList.data[j+1]
	}
	mySqList.data[mySqList.length-1] = 0
	mySqList.length--
	return true
		
	
}

//ListDelete function is to remove element
//method2
// func (mySqList *SqList) ListDelete(i int) bool{
// 	if i < 1 || i > mySqList.length + 1 {
// 		fmt.Println("i不合法,请检查i")
// 		return false
// 	}
// 	if i > mySqList.maxSize {
// 		fmt.Println("i不合法,请检查i")
// 	}
//区别主要在这里，有两种思路
// 	for( j := i; j < mySqList.length; j++){
// 		mySqList.data[j-1] = mySqList.data[j]
// 	}
// 	mySqList.length--
// 	return true
// }

//ListFind function is to find an element
//此方法用来查找指定位置的元素
func (mySqList *SqList) ListFind(i int) bool{
	if i < 1 || i > mySqList.length + 1 {
		fmt.Println("i不合法,请检查i")
		return false
	}
	if i > mySqList.maxSize {
		fmt.Println("i不合法,请检查i")
		return false
	}
	fmt.Println(mySqList.data[i-1])
	return true
}
func main(){
	
	list := InitSqList(20)  //设置最大长度为10
	list.ListInsert(1,22)
	list.ListAppend(12)
	list.ListAppend(122)
	list.ListAppend(133)
	list.ListAppend(143)
	list.ListAppend(153)
	list.ListAppend(173)
	list.ListAppend(183)
	list.ListAppend(193)
	list.ListAppend(163)
	fmt.Println(&list)
	fmt.Println(*list)
	fmt.Println(*&list.length)
	list.ListDelete(5)
	//fmt.Println(*&list.data[2])
	fmt.Println(*list)
	list.ListFind(3)
}
