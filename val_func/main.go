package main

import "strconv"
import "fmt"

func main() {

	// confirm int size
	{
		var int_32 int32 = 2147483647
		//var int_32 int32 = 2147483648; // overflow
		var int_normal int = 9223372036854775807
		//var int_normal int = 9223372036854775808; // overflow->original int size 64bit
		fmt.Println(int_32, int_normal)
	}
	// convert type
	{
		i, _ := strconv.Atoi("-42")
		s := strconv.Itoa(-42)
		fmt.Println(i, s)
	}
}
