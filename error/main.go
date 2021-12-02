package main

import (
	"errors"
	"fmt"
	//"github.com/rs/zerolog"
	//"github.com/rs/zerolog/log"
	"log"
	"os"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	testname("simple error handle")
	{
		employee, err := getInformation(1001)
		if errors.Is(err, ErrNotFound) {
			fmt.Printf("NOT FOUND: %v\n", err)
		} else if err != nil {
			fmt.Print(err)
		} else {
			fmt.Println(employee)
		}
	}
	testname("log")
	{
		log.SetPrefix("main(): ")
		log.Print("I'm a log")
		//log.Fatal("I'm a Fatal error log")
		//log.Panic("I'm a Panic error log")
	}
	testname("log to file")
	{
		file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		log.SetOutput(file)
		log.Print("I'm a log")
		log.SetOutput(os.Stdout)
	}
	//testname("flame work")
	//{
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	//log.Print("I'm a log message")

	//log.Debug().
	//Int("EmployeeID", 1001).
	//Msg("Getting employee information")

	//log.Debug().
	//Str("Name", "John").
	//Send()
	//}
}

var ErrNotFound = errors.New("Employee not found!")

func getInformation(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}
	employee, err := apiCallEmployee(1000)
	if err != nil {
		return nil, fmt.Errorf("Got an error when getting the employee information: %v", err)
	}
	return employee, err
}

/* retry when error */
//func getInformation(id int) (*Employee, error) {
//    for tries := 0; tries < 3; tries++ {
//        employee, err := apiCallEmployee(1000)
//        if err == nil {
//            return employee, nil
//        }
//        fmt.Println("Server is not responding, retrying ...")
//        time.Sleep(time.Second * 2)
//    }
//    return nil, fmt.Errorf("server has failed to respond to get the employee information")
//}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}

func testname(s string) {
	fmt.Println("---------------[", s, "]")
}
