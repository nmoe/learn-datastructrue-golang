package main

import (
	"fmt"
	//"sort"
)

//InsertSort method
//直接插入排序
func InsertSort(a []int) {
	var i, j int
	for i = 1; i < len(a); i++ { //第一次是认为第一个元素已经是有序的,故从第二个位置开始

		temp := a[i]                                //将当前元素设为哨兵,临时变量temp,可以考虑直接赋值给a[0],代码改进的地方
		for j = i - 1; j >= 0 && temp < a[j]; j-- { //比较元素大小,从后往前比较
			a[j+1] = a[j] //向后挪位置
		}
		a[j+1] = temp //将哨兵元素赋值到插入位置

	}

	fmt.Println(a)
}

//InsertSort1 method
//折半插入排序
func InsertSort1(a []int) {
	var i, j, low, high, mid int
	for i = 2; i < len(a); i++ {
		temp := a[i]
		low, high = 1, i-1
		if low <= high {
			mid = (high + low) / 2
			if a[mid] > temp {
				high = mid - 1
			} else {
				low = mid + 1
			}
			for j = i - 1; j >= high+1; j-- {
				a[j+1] = a[j]
			}
			a[high+1] = temp
		}
	}
	fmt.Println(a)
}

//ShellSort method
//希尔排序
func ShellSort(a []int) {
	n := len(a)
	
	for dk := n / 2; dk >= 1; dk = dk / 2 {
		for i := dk + 1; i < n; i++ {
			temp := a[i]
			j:=i-dk
			for ; j >= 0 && temp < a[j]; j -= dk {
				a[j+dk] = a[j]
				
			}
			a[j+dk] = temp
		}
	}
	fmt.Println(a)
}


func shellSort(arr []int) {
    len := len(arr)
    // 增量
    inc := 2
    // 步长
    step := len / inc
    for ; step > 0; step /= inc {
        for r := step; r < len; r ++ {
            // 最小下标为0
            l := r - step
            tmp := arr[r]
            for ; l >= 0 && tmp < arr[l]; l -= step {
                arr[l+step] = arr[l]
            }
            arr[l+step] = tmp
        }
	}
	fmt.Println(arr)
}
// func main() {
// 	a := []int{1, 2, 6, 4, 78, 50}
// 	fmt.Println(a[1:5])
// 	InsertSort(a)
// 	InsertSort1(a)
// 	fmt.Printf("%T\n", a)
// 	fmt.Println(a[1:5])
// 	ShellSort(a)
// 	shellSort(a)
// }
