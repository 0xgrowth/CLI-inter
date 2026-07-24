package main

import (
	"fmt"
	"time"
)

type Student struct {
    ID    int
    Name  string
    Age   int
    Score float64
}
func addStudent(students []Student) []Student {
	a := Student{
		ID: 0,
		Name: "John",
		Age: 19,
		Score: 50.7,
	}
	b := Student{
		ID: 1,
		Name: "Susan",
		Age: 18,
		Score: 78.7,
	}
	c := Student{
		ID: 2,
		Name: "Jake",
		Age: 19,
		Score: 89.5,
	}
	d := Student{
		ID: 3,
		Name: "Jane",
		Age: 20,
		Score: 50.0,
	}
	fmt.Println(a,b,c,d)
}

func displayStudents(students []Student) {

}

func dent() {
	students := []Student{}
	for {
		menu()
	
		// read user's choice
	
		switch choice {
		case 1:
			// add
			time.Sleep(time.Second)
			students = addStudent(students)
		case 2:
			// view
			time.Sleep(time.Second)
			displayStudents(student)
		case 6:
			return
		}
	}

}