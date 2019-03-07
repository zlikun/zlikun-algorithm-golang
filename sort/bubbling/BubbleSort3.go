package main

import "fmt"

// 冒泡排序，优化二：减少每轮比较次数
func sort3(array []int) {
	border := len(array) - 1
	for i := 0; i < len(array)-1; i++ {
		flag := true
		lastIndex := -1
		for j := 0; j < border; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				flag = false
				lastIndex = j
			}
		}
		border = lastIndex
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
	sort3(array)

	// 排序后： [8 8 32 142 153 221 231 235 244 248]
	fmt.Println("排序后：", array)

}
