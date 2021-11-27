package main

import (
	"fmt"
)

func main() {
	testname("Fibonacci")
	{
		Fibo := FibonacciSlice(0)
		fmt.Println(Fibo)
		Fibo = FibonacciSlice(6)
		fmt.Println(Fibo)
		Fibo = FibonacciSlice(7)
		fmt.Println(Fibo)
		Fibo = FibonacciSlice(8)
		fmt.Println(Fibo)
	}
}

func testname(s string) {
	fmt.Println("---------------[", s, "]")
}

func FibonacciSlice(cnt int) (Fibo []int) {
	lower_limit := 2
	for idx := 0; idx < lower_limit && lower_limit <= cnt; idx++ {
		Fibo = append(Fibo, 1)
	}
	for idx := lower_limit; idx < cnt; idx++ {
		Fibo = append(Fibo, Fibo[idx-1]+Fibo[idx-2])
	}
	return
}
