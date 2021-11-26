package main

import (
	"fmt"
)

func main() {
	testname("array")
	{
		coast := [5]string{"zushi", "yuigahama", "inamura", "shichiri"}
		coast2 := [...]string{"zushi", "yuigahama", "inamura", "shichiri"}
		fmt.Println("Coast :", coast)
		fmt.Println("size  :", len(coast))
		fmt.Println("Coast2:", coast2)
		fmt.Println("size2 :", len(coast2))
		numbers := [...]int{99: -1} // specify -> numbets[99] = -1
		fmt.Println("Numbers :", numbers)
	}
}

func testname(s string) {
	fmt.Println("---------------[", s, "]")
}
