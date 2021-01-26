package BTree

import (
	"fmt"
	
)
type BTnode struct{
	Data string		//存储数据
	Left *BTnode    //指向左子树的结点
	Right *BTnode	//指向右子树的结点
}

//NewBtree function 
//添加结点构造树
func NewBtree(data string) *BTnode{
	return &BTnode{Data: data}
}

//PreTraverse method
//先序遍历-递归算法
func (root *BTnode)PreTraverse(){
	if root == nil{
		return 
	}
	fmt.Print(root.Data+" ")
	root.Left.PreTraverse()
	root.Right.PreTraverse()
} 

//InorderTraverse method
//中序遍历递归算法
func (root *BTnode) InorderTraverse(){
	if root == nil{
		return
	}
	root.Left.InorderTraverse()
	fmt.Print(root.Data+" ")
	root.Right.InorderTraverse()
}

//PostTraverse method
//后序遍历递归算法
func (root *BTnode) PostTraverse(){
	if root ==nil{
		return
	}
	root.Left.PostTraverse()
	root.Right.PostTraverse()
	fmt.Print(root.Data+" ")
}

//
//


// func main(){
// 	root:=&BTnode{
// 		Data: "A",
// 		left: nil,
// 		right: nil,
// 	}
// 	root.left = new(BTnode)
// 	root.left.Data="B"
// 	root.right = new(BTnode)
// 	root.right.Data="C"
// 	root.left.left = new(BTnode)
// 	root.left.left.Data="D"
// 	root.left.right = new(BTnode)
// 	root.left.right.Data="F"
// 	root.left.right.left = new(BTnode)
// 	root.left.right.left.Data="E"
// 	root.right.left = new(BTnode)
// 	root.right.left.Data="G"
// 	root.right.right = new(BTnode)
// 	root.right.right.Data="I"
// 	root.right.left.right = new(BTnode)
// 	root.right.left.right.Data="H"
// 	root.PreTraverse()
// }