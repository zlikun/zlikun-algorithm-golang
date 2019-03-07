package main

import "fmt"

// 冒泡排序，优化三：鸡尾酒排序，两边同时冒泡（小向左，大向右）
func sort4(array []int) {
	left, right := 0, len(array)-1
	for left < right {
		// 小数左移
		for j := right; j > left; j-- {
			if array[j-1] > array[j] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
		left++
		// 大数右移
		for j := left; j < right; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
		right--
	}

}

func main() {

	array := []int{8, 231, 244, 32, 221, 248, 8, 153, 142, 235}

	// 排序前： [8 231 244 32 221 248 8 153 142 235]
	fmt.Println("排序前：", array)

	// 冒泡排序
	sort4(array)

	// 排序后： [8 8 32 142 153 221 231 235 244 248]
	fmt.Println("排序后：", array)

}
