package main

import "fmt"

//SqStack struct
//定义顺序栈的类型
type SqStack struct {
	//顺序栈的元素
	data []int
	//栈顶指针
	top int
	//栈底指针
	base int
	//最大个数
	maxSize int
}

//InitStack function is initnial stack
//初始化顺序栈,传入InitSize定义当前数据个数,最大个数设置为40
func InitStack(InitSize int) *SqStack {

	return &SqStack{data: make([]int, InitSize), maxSize: 40, top: 0, base: 0}
}

//StackEmpty method to know is stackempty
//判断是否空
func (s *SqStack) StackEmpty() bool {
	if s.base == s.top { //栈顶指针和栈底指针相等
		return true
	}
	return false

}

//StackLength method is getting length
//求顺序栈长度,直接使用栈顶指针减去栈底指针,得到指针相隔多少就是元素个数
func (s *SqStack) StackLength() int {
	if s.base == s.top {
		return 0
	}
	return s.top - s.base
}

//Clear method is clearing elements
//清空栈,直接把栈顶指针移到栈底即可,注意只是清空,栈在内存中还是存在的，并且忽略里面到底存了什么。
//清空栈的意思就是栈顶指针和栈底指针都指向栈底,代表里面没有元素了
func (s *SqStack) Clear() bool {
	if s.base == s.top {
		return true
	}
	s.top = s.base
	return true

}

//Push method is pushing an element into stack
//压栈操作,将一个元素入栈,向上移动指针
func (s *SqStack) Push(e int) bool {
	if s.top-s.base == s.maxSize {
		return false
	}
	s.data[s.top] = e
	s.top++
	return true
}

//Pop function is pop top element
//弹栈操作,将栈顶元素弹出
func (s *SqStack) Pop() bool {
	if s.base == s.top {
		return false
	}
	// x = s.data[s.top]
	s.top-- //可以使用for循环(参考遍历的循环),将所有元素都弹出
	return true
}

//GetElem is get top element
//读取栈顶元素
func (s *SqStack) GetElem() int {
	if s.base == s.top {
		return -1
	}
	x := s.data[s.top]
	return x
}

//Traverse function is get all elements
//遍历栈元素
func (s *SqStack) Traverse() {

	for s.top != s.base {
		fmt.Println(s.data[s.top])
		s.top--
	}

}

// func main() {
// 	S := InitStack(10)

// 	S.Push(12)
// 	S.Push(1212)
// 	S.Push(1242)
// 	S.Push(212)
// 	S.Push(12345)
// 	fmt.Println("栈底位置", S.base)
// 	fmt.Println("栈顶位置", S.top)
// 	fmt.Println("当前栈元素个数", S.StackLength())
// 	S.Pop()
// 	fmt.Println("栈顶元素", S.GetElem())
// 	fmt.Println(S.StackLength())
// 	fmt.Printf("S.top的类型%T", S.top)
// 	fmt.Println()
// 	S.Traverse()
// }
