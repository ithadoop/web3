/* 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。 */
package main

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person     // 组合 Person 结构体
	EmployeeID string
}

func (e Employee) PrintInfo() {
	println("Name:", e.Name)
	println("Age:", e.Age)
	println("EmployeeID:", e.EmployeeID)
}

func main() {
	// 创建 Person 实例
	person := Person{Name: "Alice", Age: 30}
	// 创建 Employee 实例，组合 Person
	employee := Employee{Person: person, EmployeeID: "E12345"}

	// 调用 PrintInfo 方法输出员工信息
	employee.PrintInfo()
}