package main

import "fmt"

// 插入排序
func sort(array []int) {
	// 将序列分为有序和无序两部分，将无序序列中每个元素依次插入有序序列中即可实现插入排序，通常以第一个元素为有序序列，其后其它元素为无序序列
	for i := 1; i < len(array); i++ {
		sentinel := array[i]
		for j := i - 1; j >= 0; j-- {
			if sentinel < array[j] {
				array[j+1], array[j] = array[j], sentinel
			} else {
				break
			}
		}
	}

}

func main() {

	array := []int{8, 231, 244, 32, 221, 248, 8, 153, 142, 235}

	// 排序前： [8 231 244 32 221 248 8 153 142 235]
	fmt.Println("排序前：", array)

	// 插入排序
	sort(array)

	// 排序后： [8 8 32 142 153 221 231 235 244 248]
	fmt.Println("排序后：", array)

}
