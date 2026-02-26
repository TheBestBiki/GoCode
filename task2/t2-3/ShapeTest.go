package main

import (
	"fmt"
	"math"
)

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。考察点 ：接口的定义与实现、面向对象编程风格。
*/

func main() {
	// 创建 Rectangle 实例
	rect := Rectangle{Width: 5.0, Height: 3.0}
	fmt.Printf("矩形 - 宽度: %.2f, 高度: %.2f\n", rect.Width, rect.Height)
	fmt.Printf("面积: %.2f\n", rect.Area())
	fmt.Printf("周长: %.2f\n\n", rect.Perimeter())

	// 创建 Circle 实例
	circle := Circle{Radius: 4.0}
	fmt.Printf("圆形 - 半径: %.2f\n", circle.Radius)
	fmt.Printf("面积: %.2f\n", circle.Area())
	fmt.Printf("周长: %.2f\n", circle.Perimeter())
}

// Shape 接口定义
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle 结构体
type Circle struct {
	Radius float64
}

// Area Rectangle 的 Area 方法实现
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter Rectangle 的 Perimeter 方法实现
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area Circle 的 Area 方法实现
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter Circle 的 Perimeter 方法实现
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
