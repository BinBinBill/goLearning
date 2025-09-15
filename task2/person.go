package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeId int
}

func (e Employee) PrintInfo() {
	fmt.Printf("Employee Info:\nName: %s\nAge: %d\nEmployee ID: %d\n",
		e.Name, e.Age, e.EmployeeId)
}
