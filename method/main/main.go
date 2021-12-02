package main

import (
	"fmt"
	"github.com/myuser/geometry"
	"strings"
)

type triangle struct {
	size int
}

type square struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (s square) perimeter() int {
	return s.size * 4
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

type coloredTriangle struct {
	triangle
	color string
}

func (t coloredTriangle) perimeter() int {
	return t.size * 3 * 2
}

func main() {
	testname("base method")
	{
		t := triangle{3}
		s := square{4}
		fmt.Println("Perimeter (triangle): ", t.perimeter())
		fmt.Println("Perimeter (square): ", s.perimeter())
		t.doubleSize()
		fmt.Println("Size: (triangle): ", t.size)
	}
	testname("make method another type")
	{
		s := upperstring("Learning Go!")
		fmt.Println(s)
		fmt.Println(s.Upper())
	}
	testname("method overload")
	{
		t := coloredTriangle{triangle{3}, "blue"}
		fmt.Println("Size: ", t.size)
		fmt.Println("Perimeter: ", t.perimeter()) // if coloredTriangle don't have perimeter(), go use triangle's perimeter() automaticaly
		fmt.Println("Perimeter (triangle): ", t.triangle.perimeter())
	}
	testname("external package")
	{
		t := geometry.Triangle{}
		t.SetSize(3)
		fmt.Println("Perimeter: ", t.Perimeter())
		t.doubleSize() // cannot refer
	}
}

func testname(s string) {
	fmt.Println("---------------[", s, "]")
}
