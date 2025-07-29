package main

import "fmt"

func twoSum(nums []int, target int) []int {

	tempMap := make(map[int]int)
	for index, value := range nums {
		tempMap[value] = index
	}
	for index, value := range nums {
		secondNum := target - value
		secondValue, exists := tempMap[secondNum]
		if exists && secondValue != index {
			return []int{index, secondValue}
		}

	}
	return nil
}

//main执行函数
func main() {
	fmt.Println("数组：[2,7,11,15],的输出结果是：",twoSum([]int{2,7,11,15},9))
	fmt.Println("数组：[3,2,4],的输出结果是：", twoSum([]int{3, 2, 4}, 6))
	fmt.Println("数组：[3,3],的输出结果是：", twoSum([]int{3, 3}, 6))

}
