package main

type Shape2D interface {
	Area() float32
	Perimeter() float32
}

type Triangle struct {
	Base   float32
	Height float32
}

func (t Triangle) Area() float32 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float32 {
	return t.Base + t.Height + (t.Base*t.Base + t.Height*t.Height)
}

func DrawShape(s Shape2D) {
	tri, ok := s.(Triangle)
	if ok {
		println("Drawing a triangle with base:", tri.Base, "and height:", tri.Height)
	} else {
		println("Unknown shape")
	}
}

func main() {
	triangle := Triangle{
		Base:   3,
		Height: 4,
	}

	DrawShape(triangle)
	println("Area:", triangle.Area())
	println("Perimeter:", triangle.Perimeter())
}
