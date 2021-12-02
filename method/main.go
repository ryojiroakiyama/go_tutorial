package main

import (
	"fmt"
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
}

func testname(s string) {
	fmt.Println("---------------[", s, "]")
}
