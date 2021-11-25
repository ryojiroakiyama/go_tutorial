package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"regexp"
	"time"
)

func createnumber() int {
	return 10
}

func main() {

	// if
	{
		if num := createnumber(); num < 0 {
			fmt.Println(num, "if negative")
		} else if num < 10 {
			fmt.Println(num, "is one digit")
		} else {
			fmt.Println(num, "is multiple digits")
		}
	}

	// switch
	{
		sec := time.Now().Unix()
		rand.Seed(sec)
		i := rand.Int31n(10)

		switch i {
		case 0:
			fmt.Println("zero")
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		default:
			fmt.Println(i)
		}
	}
	// switch
	{
		var to_everyone = regexp.MustCompile("^Broad cast message")
		var to_jim = regexp.MustCompile(".*Dear jim")
		var to_me = regexp.MustCompile(".*Dear akiyama")
		message := "Broad cast message~~~~~Dear akiyama"

		switch {
		case to_everyone.MatchString(message):
			fmt.Println("to everyone")
		case to_jim.MatchString(message):
			fmt.Println("to jim")
		case to_me.MatchString(message):
			fmt.Println("to me")
		default:
			fmt.Println("missing destination")
		}
	}
	// for
	{
		sum := 0
		for i := 5; i < 5; i++ {
			sum += i
		}
		fmt.Println("sum is ", sum)
	}
	// for
	{
		num := 0
		for num == 0 {
			num = rand.Int() % 2
		}
		fmt.Println(num)
	}
	// defer
	{
		newfile, error := os.Create("learnGo.txt")
		if error != nil {
			fmt.Println("error : fail to create file")
			return
		}
		defer newfile.Close()

		if _, error = io.WriteString(newfile, "learn go"); error != nil {
			fmt.Println("error : fail to write to file")
			return
		}
		newfile.Sync()
	}
}
