package main

import "fmt"

// 冒泡排序
func sort(array []int) {
	// 外层循环控制比较轮数
	for i := 0; i < len(array)-1; i++ {
		// 内层循环控制每轮比较，大数向右移，直到有序位（一轮仍会比较完，所以有优化空间）
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

}

func main() {

	array := []int{8, 231, 244, 32, 221, 248, 8, 153, 142, 235}

	// 排序前： [8 231 244 32 221 248 8 153 142 235]
	fmt.Println("排序前：", array)

	// 冒泡排序
	sort(array)

	// 排序后： [8 8 32 142 153 221 231 235 244 248]
	fmt.Println("排序后：", array)

}
