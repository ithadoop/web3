package main

/* 指针
1题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
 */

import (
	"fmt"
)

func increaseByTen(num *int) {

	*num += 10 // 使用指针修改值
}

/* 2题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。 */

func doubleSliceElements(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2 // 使用指针修改切片中的元素
	}
}

func main() {
	var value int = 5
	fmt.Println("Before:", value) // 输出：Before: 5

	increaseByTen(&value)        // 传递指针
	fmt.Println("After:", value) // 输出：After: 15

	value = 10
	fmt.Println("Before:", value) // 输出：Before: 5

	increaseByTen(&value)        // 传递指针
	fmt.Println("After:", value) // 输出：After: 15

	
}
