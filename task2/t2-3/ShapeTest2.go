package main

import (
	"fmt"
	"math"
)

/*
关于使用值接收者实现接口 和 使用指针接收者实现接口 的区别
推荐使用： 使用指针接收者实现接口
*/

/*
使用值接收者实现接口或者使用指针接收者实现接口， 都认为是实现了Shape接口是吧？？
Shape有2个方法，如果一个方法使用值接收者，一个使用指针接收者，也算是实现了Shape的所有接口是吧？？

答：
1. 关于第一个问题：
使用值接收者实现接口或者使用指针接收者实现接口，都认为是实现了Shape接口是吧
答案：是的，但这有个重要前提：
如果所有方法都用值接收者 → T 和 *T 都实现了接口
如果所有方法都用指针接收者 → 只有 *T 实现了接口
如果混合使用 → 只有 *T 实现了接口
2. 关于第二个问题：
Shape有2个方法，如果一个方法使用值接收者，一个使用指针接收者，也算是实现了Shape的所有接口是吧
答案：部分正确，但有限制。混合使用，只有 *T 才能实现接口。即需要注意在赋值给变量的时候，可能会出现编译错误

// MixedShape 的情况
mixed := MixedShape{Width: 5, Height: 3, Radius: 2}
// var shape Shape = mixed  // ❌ 编译错误！MixedShape 值没有实现 Shape

// 指针接收者实现的需要注意加上&号
mixedPtr := &MixedShape{Width: 5, Height: 3, Radius: 2}
var shape Shape = mixedPtr  // ✅ 可以！*MixedShape 实现了 Shape

*/

type Shape2 interface {
	Area() float64
	Perimeter() float64
}

// Rectangle2 使用值接收者实现接口
type Rectangle2 struct {
	Width  float64
	Height float64
}

func (r Rectangle2) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle2) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle2 使用指针接收者实现接口
type Circle2 struct {
	Radius float64
}

func (c *Circle2) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle2) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// Rectangle2 可以用值或指针调用
	rect1 := Rectangle2{Width: 5, Height: 3}
	var Shape21 Shape2 = rect1 // 值可以直接赋值给接口
	fmt.Printf("矩形面积: %.2f\n", Shape21.Area())

	rect2 := &Rectangle2{Width: 4, Height: 6}
	var Shape22 Shape2 = rect2 // 指针也可以赋值给接口
	fmt.Printf("矩形面积: %.2f\n", Shape22.Area())

	// Circle2 只能用指针调用（因为方法是 ptr receiver）
	Circle2 := &Circle2{Radius: 3}
	var Shape23 Shape2 = Circle2 // 只有指针可以赋值给接口
	fmt.Printf("圆形面积: %.2f\n", Shape23.Area())

	// 下面这行会编译错误！
	// Circle2Value := Circle2{Radius: 3}
	// var Shape24 Shape2 = Circle2Value  // 错误：Circle2 值没有实现 Shape2 接口
}
