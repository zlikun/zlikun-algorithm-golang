package main

import "fmt"

// 冒泡排序，优化一：减少比较轮数
func sort2(array []int) {
	for i := 0; i < len(array)-1; i++ {
		flag := true
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}

func main() {

	array := []int{8, 231, 244, 32, 221, 248, 8, 153, 142, 235}

	// 排序前： [8 231 244 32 221 248 8 153 142 235]
	fmt.Println("排序前：", array)

	// 冒泡排序
	sort2(array)

	// 排序后： [8 8 32 142 153 221 231 235 244 248]
	fmt.Println("排序后：", array)

}
