package main

import (
	"fmt"
)

//BubbleSort method
//冒泡排序
func BubbleSort(a []int) {
	n := len(a)
	for i := 0; i < n-1; i++ {  //总共比较i趟
		flag1 := false
		for j := n - 1; j > i; j-- { //for j:=1; j<=n-i; j++  第i趟需要比较n-i次
			if a[j-1] > a[j] {
				//Swap(&a[j-1], &a[j])
				a[j-1], a[j] = a[j], a[j-1]
			}

			flag1 = true
		}
		if flag1 == false {
			break
		}
	}
	fmt.Println(a)

}

//Swap method
//交换值,注意指针传递,才会真正改变实参的值
func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp

}

// func Swap1(a, b int) {
// 	temp := a
// 	a = b
// 	b = temp

// }

//冒泡排序简单的一个递归算法
func bubbleSort(arr []int, len int) {
	if len == 1 {
		return
	}
	for i := 0; i < len-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	bubbleSort(arr, len-1)
}


//
//快速排序递归算法
func qsort(array []int, low, high int) {
    if low < high {
        m := partition(array, low, high)
        // fmt.Println(m)
        qsort(array, low, m-1)
        qsort(array, m+1, high)
    }
}

//划分方法
func partition(array []int, low, high int) int {
    key := array[low]
    tmpLow := low
    tmpHigh := high
    for {
        //查找小于等于key的元素，该元素的位置一定是tmpLow到high之间，因为array[tmpLow]及左边元素小于等于key，不会越界
        for array[tmpHigh] > key {
            tmpHigh--
        }
        //找到大于key的元素，该元素的位置一定是low到tmpHigh+1之间。因为array[tmpHigh+1]必定大于key
        for array[tmpLow] <= key && tmpLow < tmpHigh {
            tmpLow++
        }

        if tmpLow >= tmpHigh {
            break
        }
        // swap(array[tmpLow], array[tmpHigh])
        array[tmpLow], array[tmpHigh] = array[tmpHigh], array[tmpLow]
        fmt.Println(array)
    }
    array[tmpLow], array[low] = array[low], array[tmpLow]
    return tmpLow
}

// func main() {
// 	//冒泡排序
// 	a := []int{1, 4, 1, 12, 45, 22, 14, 0, 1}
// 	BubbleSort(a)
// 	arr := []int{9, 4, 5, 1, 22, 8, 4, 7, 0, 2}
// 	bubbleSort(arr, len(arr))
// 	fmt.Println(arr)
// 	//快速排序
// 	var sortArray = []int{3, 41, 24, 76, 11, 45, 3, 3, 64, 21, 69, 19, 36}
//     fmt.Println(sortArray)
//     qsort(sortArray, 0, len(sortArray)-1)
//     fmt.Println(sortArray)
// }
