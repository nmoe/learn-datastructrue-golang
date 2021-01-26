package main

import "fmt"

//简单选择排序
func selectsort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ { //一共进行n-1趟
		min := i                     //记录最小元素位置
		for j := i + 1; j < n; j++ { //在arr[i...n-1]中选择最小的元素
			if arr[j] < arr[min] {
				min = j //更新最小的元素位置
			}

		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i] //与第i个位置交换
		}
	}
	fmt.Println(arr)
}

//堆排序
// Heap 定义堆排序过程中使用的堆结构
type Heap struct {
	arr  []int // 用来存储堆的数据
	size int   // 用来标识堆的大小
}

// adjustHeap 用于调整堆，保持堆的固有性质
func adjustHeap(h Heap, parentNode int) {
	leftNode := parentNode*2 + 1
	rightNode := parentNode*2 + 2

	maxNode := parentNode
	if leftNode < h.size && h.arr[maxNode] < h.arr[leftNode] {
		maxNode = leftNode
	}
	if rightNode < h.size && h.arr[maxNode] < h.arr[rightNode] {
		maxNode = rightNode
	}

	if maxNode != parentNode {
		h.arr[maxNode], h.arr[parentNode] = h.arr[parentNode], h.arr[maxNode]
		adjustHeap(h, maxNode)
	}
}

// createHeap 用于构造一个堆
func createHeap(arr []int) (h Heap) {
	h.arr = arr
	h.size = len(arr)

	for i := h.size / 2; i >= 0; i-- {
		adjustHeap(h, i)
	}
	return
}

// heapSort 使用堆对数组进行排序
func heapSort(arr []int) {
	h := createHeap(arr)

	for h.size > 0 {
		// 将最大的数值调整到堆的末尾
		h.arr[0], h.arr[h.size-1] = h.arr[h.size-1], h.arr[0]
		// 减少堆的长度
		h.size--
		// 由于堆顶元素改变了，而且堆的大小改变了，需要重新调整堆，维持堆的性质
		adjustHeap(h, 0)
	}
}

func main() {
	a := []int{1, 2, 10, 2, 4, 5, 52}
	selectsort(a)
}
