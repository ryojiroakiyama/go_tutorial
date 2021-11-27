package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	testname("array")
	{
		coast := [5]string{"zushi", "yuigahama", "inamura", "shichiri"}
		coast2 := [...]string{"zushi", "yuigahama", "inamura", "shichiri"}
		fmt.Println("Coast :", coast)
		fmt.Println("size  :", len(coast))
		fmt.Println("Coast2:", coast2)
		fmt.Println("size2 :", len(coast2))
		numbers := [...]int{99: -1} // specify -> numbets[99] = -1
		fmt.Println("Numbers :", numbers)
	}
	testname("multiple array")
	{
		var ThreeD [2][3][2]int
		fmt.Println("init: ", ThreeD)
		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ {
				for k := 0; k < 2; k++ {
					ThreeD[i][j][k] = (i + 1) * (j + 1) * (k + 1)
				}
			}
		}
		fmt.Println("after loop: ", ThreeD)
	}
	testname("slice")
	{
		months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
		quarter2 := months[3:6]
		quarter2Extended := quarter2[:5]
		quarter2ToEnd := months[3:]
		fmt.Println(months, len(months), cap(months))
		fmt.Println(quarter2, len(quarter2), cap(quarter2))
		fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))
		fmt.Println(quarter2ToEnd, len(quarter2ToEnd), cap(quarter2ToEnd))
		months = append(months, "messenger")
		fmt.Println(months, len(months), cap(months))

		remove := 2
		copy_months := make([]string, 13)
		copy(copy_months, months)
		if remove < len(copy_months) {
			fmt.Println("copy months:Before", copy_months, "Remove", copy_months[remove])
			copy_months = append(copy_months[:remove], copy_months[remove+1:]...)
			fmt.Println("copy months:After ", copy_months)
		}
	}
	testname("map without make")
	{
		studentsAge := map[string]int{
			"john": 32,
			"bob":  31,
		}
		studentsAge["jiro"] = 25
		fmt.Println(studentsAge)
	}
	testname("map with make")
	{
		//var studentsAge map[string]int // error because of nil map?
		studentsAge := make(map[string]int) // empty map can add element
		studentsAge["john"] = 32
		studentsAge["bob"] = 31
		studentsAge["jiro"] = 25
		fmt.Println(studentsAge)
		fmt.Println("yusei's age is", studentsAge["yusei"])
		age, exist := studentsAge["yusei"]
		fmt.Println("yusei's age is", age, ", exist flag:", exist)
		delete(studentsAge, "yusei")
		delete(studentsAge, "jiro")
		fmt.Println(studentsAge)
		//	_, age	(key ignore)
		//	name	(value ignore)
		for name, age := range studentsAge {
			fmt.Println(name, "\t", age)
		}
	}
	testname("struct")
	{
		type Employee struct {
			ID        int
			FirstName string
			LastName  string
			Address   string
		}

		employee := Employee{LastName: "Doe", FirstName: "John"}
		fmt.Println(employee)
		var employeeCopy *Employee
		employeeCopy = &employee
		employeeCopy.FirstName = "David"
		fmt.Println(employee)
	}
	testname("struct nest")
	{
		type Person struct {
			ID        int
			FirstName string
			LastName  string
			Address   string
		}
		type Employee struct {
			Person
			ManagerID int
		}
		type Contractor struct {
			Person
			CompanyID int
		}

		employee := Employee{
			Person: Person{
				FirstName: "John",
			},
		}
		employee.LastName = "Doe"
		fmt.Println(employee)
	}
	testname("struct JSON")
	{
		type Person struct {
			ID        int
			FirstName string `json:"name"`
			LastName  string
			Address   string `json:"address,omitempty"`
		}

		type Employee struct {
			Person
			ManagerID int
		}

		type Contractor struct {
			Person
			CompanyID int
		}

		employees := []Employee{
			Employee{
				Person: Person{
					LastName: "Doe", FirstName: "John",
				},
			},
			Employee{
				Person: Person{
					LastName: "Campbell", FirstName: "David",
				},
			},
		}

		data, _ := json.Marshal(employees)
		fmt.Printf("%s\n", data)

		var decoded []Employee
		json.Unmarshal(data, &decoded)
		fmt.Println(decoded)
	}
}

func testname(s string) {
	fmt.Println("---------------[", s, "]")
}
