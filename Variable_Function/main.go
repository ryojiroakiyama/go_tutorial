package main

import "strconv"
import "fmt"
import "os"

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
		i, _ := strconv.Atoi("-42") // Atoi return 2 vlues
		s := strconv.Itoa(-42)
		fmt.Println(i, s)
	}
	// receive imput
	{
		sum, mul := sum(os.Args[1], os.Args[2])
		fmt.Println("Sum:", sum)
		fmt.Println("Mul:", mul)
	}
	// func pointer
	{
		firstName := "original name"
		updateName(&firstName)
		fmt.Println(firstName)
	}
}

func sum(number1 string, number2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	return
}

func updateName(name *string) {
	*name = "updated name"
}
