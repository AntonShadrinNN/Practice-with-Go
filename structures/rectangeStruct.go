package main

import "fmt"

type Point struct {
	x, y int
}

type Rectangle struct {
	p1, p2 Point
}

func calculatePerimeter(rect Rectangle) (perimeter int) {
	perimeter = ((rect.p2.x - rect.p1.x) + (rect.p2.y - rect.p1.y)) << 1
	return
}

func calculateArea(rect Rectangle) (area int) {
	return ((rect.p2.x - rect.p1.x) * (rect.p2.y - rect.p1.y))
}

func main() {
	rect := Rectangle{Point{1, 4}, Point{5, 7}}
	fmt.Printf("Rectangle perimeter is %d\n", calculatePerimeter(rect))
	fmt.Printf("Rectangle area is %d\n", calculateArea(rect))
}
