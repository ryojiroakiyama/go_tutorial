package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}
	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

func test1() {
	testname("base goroutine apis")
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}
	//ch := make(chan string)
	ch := make(chan string, 10)

	for _, api := range apis {
		//checkAPI(api)
		go checkAPI(api, ch)
	}
	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}
	//time.Sleep(3 * time.Second)
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func send(ch chan string, message string) {
	ch <- message
}

func test2() {
	testname("channel buffer")
	size := 2
	ch := make(chan string, size)
	send(ch, "one")
	send(ch, "two")
	go send(ch, "three")
	go send(ch, "four")
	fmt.Println("All data sent to the channel ...")

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Done!")
}

func write(ch chan<- string, message string) {
	fmt.Printf("Writing: %#v\n", message)
	ch <- message
}

func read(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
}

func test3() {
	testname("channel direction")
	ch := make(chan string, 1)
	write(ch, "Hello World!")
	read(ch)
}

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done processing!"
}

func replicate(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Done replicating!"
}

func test4() {
	testname("select")
	ch1 := make(chan string)
	ch2 := make(chan string)
	go process(ch1)
	go replicate(ch2)
	for i := 0; i < 2; i++ {
		select {
		case process := <-ch1:
			fmt.Println(process)
		case replicate := <-ch2:
			fmt.Println(replicate)
		}
	}
}

func main() {
	//test1()
	//test2()
	//test3()
	test4()
}

func testname(s string) {
	fmt.Println("---------------[", s, "]")
}
