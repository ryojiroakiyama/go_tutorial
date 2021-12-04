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
	testname("Rome numerals")
	{
		/* input version */
		//var input string
		//fmt.Print("Enter Rome num: ")
		//fmt.Scanf("%s", &input) // Scanln(&input)
		//fmt.Println(Convert_Rome(input))
		fmt.Println(Convert_Rome("MCM"))
		fmt.Println(Convert_Rome("MCLX"))
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

func Convert_Rome(input string) (number int) {
	Rome := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	for i := 0; i < len(input); i++ {
		i_num := Rome[rune(input[i])]
		if i == len(input)-1 {
			number += i_num
			break
		}
		switch next_num := Rome[rune(input[i+1])]; {
		case i_num < next_num:
			number += next_num - i_num
			i++
		default:
			number += i_num
		}
	}
	if number == 0 {
		fmt.Println("error")
	}
	return
}

/* Rome numerals another code */
//package main

//import (
//    "fmt"
//)

//func romanToArabic(numeral string) int {
//    romanMap := map[rune]int{
//        'M': 1000,
//        'D': 500,
//        'C': 100,
//        'L': 50,
//        'X': 10,
//        'V': 5,
//        'I': 1,
//    }

//    arabicVals := make([]int, len(numeral)+1)

//    for index, digit := range numeral {
//        if val, present := romanMap[digit]; present {
//            arabicVals[index] = val
//        } else {
//            fmt.Printf("Error: The roman numeral %s has a bad digit: %c\n", numeral, digit)
//            return 0
//        }
//    }

//    total := 0

//    for index := 0; index < len(numeral); index++ {
//        if arabicVals[index] < arabicVals[index+1] {
//            arabicVals[index] = -arabicVals[index]
//        }
//        total += arabicVals[index]
//    }

//    return total
//}

//func main() {
//    fmt.Println("MCLX is", romanToArabic("MCLX"))
//    fmt.Println("MCMXCIX is ", romanToArabic("MCMXCIX"))
//    fmt.Println("MCMZ is", romanToArabic("MCMZ"))
//}
