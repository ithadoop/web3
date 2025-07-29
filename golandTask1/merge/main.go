package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		v1, v2 := intervals[i], intervals[j]
		return v1[0] < v2[0] || v1[0] == v2[0] && v1[1] < v2[1]
	})
	stack := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		cnt1 := stack[len(stack)-1]
		cnt2 := intervals[i]
		if cnt1[0] <= cnt2[0] && cnt1[1] >= cnt2[0] {
			cnt1[1] = max(cnt1[1], cnt2[1])
			stack[len(stack)-1] = cnt1
		} else {
			stack = append(stack, cnt2)
		}
	}
	return stack
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result := merge(intervals)
	fmt.Println(result) // 输出: [[1,6],[8,10],[15,18]]
}