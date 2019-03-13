package main

import "fmt"

// 选择排序
func sort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(array); j++ {
			if array[minIndex] > array[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			array[minIndex], array[i] = array[i], array[minIndex]
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
