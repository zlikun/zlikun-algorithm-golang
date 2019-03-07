package main

import "fmt"

// 快速排序
func sort(array []int, left int, right int) {

	if left >= right {
		return
	}

	// 以最左边元素为支点，小于支点元素交换到支点左边，大于等于支点元素交换到支点右边
	pivot := array[left]
	// 分别定义左右游标，用于分别左右遍历数组，找到需要交换的元素并记录元素位置
	low, high := left, right

	// 完成一轮比较、交换
	for low < high {
		// 从右向左遍历，找到比支点小的元素，high记录其位置
		for low < high && pivot <= array[high] {
			high--
		}
		// 交换high元素到支点位置，支点位置元素由支点变量记录，空出high位置
		array[low] = array[high]

		// 从左向右遍历，找到比支点大（含等于）的元素，low记录其位置
		for low < high && pivot > array[low] {
			low++
		}
		// 交换low位元素到high位，空出low位置，该位置将用于记录支点元素
		array[high] = array[low]
	}

	// 一轮比较完成后，low位置空出，用于记录支点元素（支点右移）
	array[low] = pivot

	// 以新支点位置为界，左右两个子序列分别递归实现同样逻辑，支点本身有序，不参与递归
	sort(array, left, low-1)
	sort(array, low+1, right)

}

func main() {
	array := []int{8, 231, 244, 32, 221, 248, 8, 153, 142, 235}

	// 排序前： [8 231 244 32 221 248 8 153 142 235]
	fmt.Println("排序前：", array)

	// 快速排序
	sort(array, 0, len(array)-1)

	// 排序后： [8 8 32 142 153 221 231 235 244 248]
	fmt.Println("排序后：", array)
}
