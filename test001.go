// main.go

package main

import (
	"fmt"
)

// 冒泡排序算法实现
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 交换元素
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
