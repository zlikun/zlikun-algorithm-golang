package main

import "fmt"

// 希尔排序
func sort(array []int) {
	for gap := len(array) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(array); i += gap {
			sentinel := array[i]
			for j := i - gap; j >= 0; j -= gap {
				if sentinel < array[j] {
					array[j+gap], array[j] = array[j], sentinel
				} else {
					break
				}
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
