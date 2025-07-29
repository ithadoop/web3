// 包声明
package main

// 引入包声明
import "fmt"

// 函数声明
func printInConsole(s string) {
	fmt.Println(s)
}

// 全局变量声明
var str string = "Hello world!"

/* 136. 只出现一次的数字
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
示例 1 ：
输入：nms = [2,2,1]
输出：1

示例 2 ：
输入：nums = [4,1,2,1,2]
输出：4

示例 3 ：
输入：nums = [1]
输出：1

提示：
1 <= nums.length <= 3 * 104
-3 * 104 <= nums[i] <= 3 * 104
除了某个元素只出现一次以外，其余每个元素均出现两次。
*/
/* 位运算解法（本题最佳解法）
异或运算的规律
任何数和其自身做异或运算，结果是0， a⊕a=0
任何数和 0 做异或运算，结果仍然是原来的数，即 a⊕0=a。
异或运算满足交换律和结合律，即 a⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=b。
数组中的全部元素的异或运算结果即为数组中只出现一次的数字
算法详述
初始化一个为0的变量 tmp = 0
遍历数组，用tmp去和每个元素挨个做异或运算
返回tmp，即为这个数组唯一的只出现一次的数字 */
func singleNumber(nums []int) int {
	// 初始化结果为0，然后对数组中的所有数字进行异或操作
	result := 0 //
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {
	nums := []int{1}

	fmt.Println("数组为", nums, "，只出现一次的数字：", singleNumber(nums))

	nums1 := []int{2, 2, 1}
	fmt.Println("数组为", nums1, "，只出现一次的数字：", singleNumber(nums1))

	nums2 := []int{4, 1, 2, 1, 2}
	fmt.Println("数组为", nums2, "，只出现一次的数字：", singleNumber(nums2))

}
