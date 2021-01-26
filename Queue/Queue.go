package Queue

import (
	//"container/list"
	"fmt"
)

//SqQueue struct
//顺序队列的结构,但实际上是循环队列
type SqQueue struct{
	data [10]int //定义有20个元素的整型数组作为队列元素
	front int    //队首指针
	rear int     //队尾指针
	maxSize int
}

//InitQueue method to initqueue
//初始化队列,将队首指针和队尾指针指向都指向队尾
func (Q *SqQueue) InitQueue(maxSize int){
	 Q.front=Q.rear
	 Q.maxSize=maxSize
}

//isEmpty method to know whether empty
//判断是否为空
func (Q *SqQueue) isEmpty()bool{
	if Q.front==Q.rear {
		return true
	}
		return false
}

//EnQueue method is to push element
//将元素从队尾插入队列
func (Q *SqQueue) EnQueue(e int) bool{
	if (Q.rear+1)%Q.maxSize==Q.front{
		return false
	}
	Q.data[Q.rear] = e
	Q.rear = (Q.rear+1) % Q.maxSize
	return true
}

//DeQueue method is 
//将元素弹出队列
func (Q *SqQueue) DeQueue()bool{
	if Q.front==Q.rear{
		return false
	}
	//x := Q.data[Q.front]   //可以使用for循环将弹出元素打印出来
	Q.front=(Q.front+1)%Q.maxSize
	return true
}


//Traverse method is get all elements
//遍历所有元素
func (Q *SqQueue) Traverse() {
	for Q.front!=Q.rear{
		x:=Q.data[Q.front]
		Q.front++
	fmt.Println(x)
	}
}

//GetLength method
//求队列元素个数(队列长度)
func (Q *SqQueue) GetLength() int{
	return (Q.rear-Q.front+Q.maxSize)%Q.maxSize
}

//GetHead method to GetHead element
//获取队头元素
func (Q *SqQueue) GetHead() int{
	if Q.front!=Q.rear{
		return Q.data[Q.front]
	}
	return -1

	
}





// func main(){
// 	var queue SqQueue
// 	queue.InitQueue(20)
// 	queue.EnQueue(12)
// 	queue.EnQueue(22)
// 	queue.EnQueue(32)
// 	//queue.Traverse()
	
// 	// fmt.Println(queue.EnQueue(32))
// 	// fmt.Println(queue.maxSize)
// 	//fmt.Println(queue.isEmpty())
// 	//fmt.Println(queue.DeQueue())
// 	fmt.Println(queue.rear)
// 	fmt.Println(queue.front)

// 	// for queue.front!=queue.rear{  //for循环打印弹出的元素,先进先出
// 	// 	fmt.Println(queue.data[queue.front])
// 	// 	queue.front=(queue.front+1)%queue.maxSize
// 	// }
	
// }