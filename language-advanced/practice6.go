package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int64
}

type Employee struct {
	EmployeeId int64
	Person Person
}

func (e Employee) PrintInfo() {
		fmt.Printf("%v, %v, %v", e.Person.Name, e.Person.Age, e.EmployeeId)
}

func main() {
	e := Employee{EmployeeId:1, Person: Person{Name: "xxx", Age: 10}}
	e.PrintInfo()
}

