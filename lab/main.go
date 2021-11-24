package main

import "fmt"

func main() {
	var int_32 int32 = 2147483647
	//var int_32 int32 = 2147483648;n
	var int_normal int = 9223372036854775807
	//var int_normal int = 9223372036854775808;
	fmt.Println(int_32, int_normal)
}
