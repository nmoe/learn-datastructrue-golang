package main

import (
	"fmt"
	"laborlearn/datastructure/BTree"
	
	
)

func main() {
	var root BTree.BTnode //首字母大写才可被使用
	//按顺序(二叉树结点顺序)添加结点
	root = *BTree.NewBtree("m") //使用Btree.go中的NewBtree()方法添加结点
	//root = BTree.BTnode{Data: "a"}  //这也是一种添加结点方法
	root.Left = &BTree.BTnode{Data: "b"}
	root.Right = &BTree.BTnode{Data: "c"}
	root.Left.Left = &BTree.BTnode{Data: "d"}

	root.Left.Right = &BTree.BTnode{Data: "f"}
	root.Left.Right.Left = &BTree.BTnode{Data: "e"}
	root.Right.Left = &BTree.BTnode{Data: "g"}
	root.Right.Right = &BTree.BTnode{Data: "i"}
	root.Right.Left.Right = &BTree.BTnode{Data: "h"}

	root.PreTraverse()
	fmt.Println()
	root.InorderTraverse()
	fmt.Println()
	root.PostTraverse()
 
}
