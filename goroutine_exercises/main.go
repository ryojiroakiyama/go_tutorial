package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(number float64, ch chan<- string) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}
	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)
	ch <- fmt.Sprintf("Fib(%v): %v\n", number, x)
}

func exercise1() {
	start := time.Now()
	ch := make(chan string, 20)
	for i := 1; i < 15; i++ {
		go fib(float64(i), ch)
	}
	for i := 1; i < 15; i++ {
		fmt.Printf(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func fib2(chQuit <-chan bool, chFib <-chan bool) {
	x, y := 1.0, 1.0
	cnt := 1
	for {
		select {
		case <-chQuit:
			return
		case <-chFib:
			fmt.Printf("Fib(%v): %v\n", cnt, x)
			cnt++
			x, y = y, x+y
		}
	}
}

func exercise2() {
	start := time.Now()
	chQuit := make(chan bool)
	chFib := make(chan bool)

	go fib2(chQuit, chFib)
	var name string
	for {
		fmt.Scanf("%s", &name)
		if name == "quit" {
			chQuit <- true
			break
		} else if name == "" {
			chFib <- true
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func main() {
	//exercise1()
	exercise2()
}

/* exericise2 example */
//var quit = make(chan bool)

//func fib(c chan int) {
//    x, y := 1, 1

//    for {
//        select {
//            case c <- x:
//                x, y = y, x+y
//            case <-quit:
//                fmt.Println("Done calculating Fibonacci!")
//            return
//        }
//    }
//}

//func main() {
//    start := time.Now()

//    command := ""
//    data := make(chan int)

//    go fib(data)

//    for {
//        num := <-data
//        fmt.Println(num)
//        fmt.Scanf("%s", &command)
//        if command == "quit" {
//            quit <- true
//            break
//        }
//    }

//    time.Sleep(1 * time.Second)

//    elapsed := time.Since(start)
//    fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
//}
