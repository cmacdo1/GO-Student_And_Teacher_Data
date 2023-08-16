// studentAndTeacher Data populates structs and prints out the info
package main

import (
	"fmt"
)

// Defining of Address, Student and Teacher structs with the different types of data it is to hold
type Address struct {
	Street  string
	City    string
	State   string
	zipCode int
}

type Student struct {
	Studentid int
	Name      string
	Address
}

type Teacher struct {
	EmployeeID int
	Name       string
	Salary     float64
	Address
}

// Accesses the Student and Address structs and prints out the information
func printStudentInfo(s Student) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Student ID:", s.Studentid)
	fmt.Println("Street:", s.Street)
	fmt.Println("City:", s.City)
	fmt.Println("State:", s.State)
	fmt.Println("Zip Code:", s.zipCode)
}

// Accesses the Teacher and Address structs and prints out the information
func printTeacherInfo(t Teacher) {
	fmt.Println("Name:", t.Name)
	fmt.Println("Employee ID:", t.EmployeeID)
	fmt.Println("Salary:", t.Salary)
	fmt.Println("Street:", t.Street)
	fmt.Println("City:", t.City)
	fmt.Println("State:", t.State)
	fmt.Println("Zip Code:", t.zipCode)
}

func main() {
	// the two fmt.Println lines are to clean up the data and make it more readable on the screen
	fmt.Println()
	fmt.Println("-------------------------")
	// creation of the first student in Student and Address struct
	s1 := Student{Name: "Chris MacDonald"}
	s1.Studentid = 12345
	s1.Street = "1522 S Price Rd"
	s1.City = "Tempe"
	s1.State = "Arizona"
	s1.zipCode = 85282
	// Access the function to print the students data
	printStudentInfo(s1)

	fmt.Println("-------------------------")
	// creation of the second student in Student and Address struct
	s2 := Student{Name: "Lisa Jannings"}
	s2.Studentid = 413573
	s2.Street = "4726 E Broadway Rd"
	s2.City = "Tampa"
	s2.State = "Florida"
	s2.zipCode = 33065
	// Access the function to print the students data
	printStudentInfo(s2)

	fmt.Println("-------------------------")
	// creation of the first teacher in Teacher and Address struct
	t1 := Teacher{Name: "John Doe"}
	t1.EmployeeID = 98765
	t1.Salary = 85461.23
	t1.Street = "123 W Main St"
	t1.City = "Phoenix"
	t1.State = "Arizona"
	t1.zipCode = 85045
	// Access the function to print the teachers data
	printTeacherInfo(t1)

	fmt.Println("-------------------------")
	// creation of the first teacher in Teacher and Address struct
	t2 := Teacher{Name: "Jane Smith"}
	t2.EmployeeID = 41065
	t2.Salary = 105300.46
	t2.Street = "9123 N Recker Dr"
	t2.City = "Columbus"
	t2.State = "Ohio"
	t2.zipCode = 43081
	// Access the function to print the teachers data
	printTeacherInfo(t2)
	// Println used to make the data more readable on the screen
	fmt.Println()
}
