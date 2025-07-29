// 包声明
package main

// 引入包声明
import "fmt"

// 函数声明
//先定义一个栈结构
type stack struct {
	top  int
	data []byte
}

//栈结构方法-入栈
func (s *stack) push(b byte) {
	s.data = append(s.data, b)
	s.top++
}

//栈结构方法-出栈
func (s *stack) pop() byte {
	b := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.top--
	return b
}

//栈结构方法-是否为空栈
func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

//判断是否为有效的括号
func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}
	byteslice := []byte(s)
	helpstack := &stack{top: -1}
	for i := 0; i < len(s); i++ {
		if byteslice[i] == '(' || byteslice[i] == '[' || byteslice[i] == '{' {
			helpstack.push(byteslice[i])
		} else if byteslice[i] == ')' || byteslice[i] == ']' || byteslice[i] == '}' {
			if helpstack.isEmpty() {
				return false
			} else {
				b1 := helpstack.pop()
				if !match(b1, byteslice[i]) {
					return false
				}
			}
		} else {
			return false
		}
	}
	return helpstack.isEmpty()
}

//判断字符是否匹配
func match(b1, b2 byte) bool {
	if (b1 == '(' && b2 == ')') || (b1 == '[' && b2 == ']') || (b1 == '{' && b2 == '}') {
		return true
	}
	return false
}

func main() {
	var s string = "([])"
	fmt.Println(s, "判断结果：", isValid(s))
	s = "()[]{}"
	fmt.Println(s, "判断结果：", isValid(s))

	s = "(]"
	fmt.Println(s, "判断结果：", isValid(s))

	s = "([])"
	fmt.Println(s, "判断结果：", isValid(s))

	s = "([)]"
	fmt.Println(s, "判断结果：", isValid(s))

}
