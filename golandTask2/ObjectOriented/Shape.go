/* 向对象
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
 */

package main 
type Shape interface {
	Area() float64
	Perimeter() float64
}
//长方形
type Rectangle struct {
	Width  float64
	Height float64
}
//圆形
type Circle struct {
	Radius float64
}
//长方形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
//长方形周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
//圆形面积
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

//圆形周长
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func printShapeInfo(s Shape) {
	println("Area:", s.Area())
	println("Perimeter:", s.Perimeter())
}

func main() {
	// 创建长方形实例
	rect := Rectangle{Width: 5, Height: 10}
	// 创建圆形实例
	circ := Circle{Radius: 7}

	// 调用长方形的方法
	printShapeInfo(rect)
	// 调用圆形的方法
	printShapeInfo(circ)
}
