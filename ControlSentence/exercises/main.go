package main

import (
	"fmt"
)

func main() {
	testname("FizzBuzz")
	{
		for i := 1; i <= 100; i++ {
			t := i % 3
			f := i % 5
			switch {
			case t == 0 && f == 0:
				fmt.Print("FizzBuzz")
			case t == 0:
				fmt.Print("Fizz")
			case f == 0:
				fmt.Print("Buzz")
			default:
				fmt.Print(i)
			}
			if i%10 == 0 {
				fmt.Println()
			} else {
				fmt.Print(" ")
			}
		}
	}
	testname("Prime numbers less than 20:")
	{
		for number := 1; number <= 20; number++ {
			if findprimes(number) {
				fmt.Printf("%v ", number)
			}
		}
		fmt.Println()
	}
	testname("panic")
	{
		val := 0
		for {
			fmt.Print("Enter number: ")
			fmt.Scanf("%d", &val)
			// switch is better
			if val == 0 {
				fmt.Println("0 is neither negative nor positive")
			} else if val > 0 {
				fmt.Println("You entered:", val)
			} else {
				panic("You entered negative number")
			}
		}
	}
}

func findprimes(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	if number != 1 {
		return true
	} else {
		return false
	}
}

func testname(s string) {
	fmt.Println("---------------[", s, "]")
}
