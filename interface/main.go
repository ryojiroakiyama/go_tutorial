package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

// for base interface
type Shape interface {
	Perimeter() float64
	Area() float64
}
type Square struct {
	size float64
}

func (s Square) Area() float64 {
	return s.size * s.size
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printInformation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

// for stringer interface
type Stringer interface {
	String() string
}

type Person struct {
	Name, Country string
}

func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

// for extend existing implementation
type GitHubResponse []struct {
	FullName string `json:"full_name"`
}

/* standard interface
** type Writer interface {
** 		Write(p []byte) (n int, err error)
** }
 */
type customWriter struct{}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GitHubResponse
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
}

// for custom server API
type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	testname("base interface")
	{
		var s Shape = Square{3}
		printInformation(s)

		c := Circle{6}
		printInformation(c)
	}
	testname("stringer interface")
	{
		rs := Person{"John Doe", "USA"}
		ab := Person{"Mark Collins", "United Kingdom"}
		fmt.Printf("%s\n%s\n", rs, ab) // Printf use Stringer interface
	}
	testname("extend existing implementation")
	{
		resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		writer := customWriter{}
		io.Copy(writer, resp.Body)
	}
	testname("custom server API")
	{
		db := database{"Go T-Shirt": 25, "Go Jacket": 55}
		log.Fatal(http.ListenAndServe("localhost:8000", db))
	}
}

func testname(s string) {
	fmt.Println("---------------[", s, "]")
}
