package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			fmt.Println("前", nums)
			nums[slow] = nums[fast]
			fmt.Println("后", nums)
			slow++
		}
	}

	return slow
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))                      // 输出: 2
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})) // 输出: 5

}
