package main

import (
	"fmt"
)

type Student struct {
    ID    int
    Name  string
    Age   int
    Score float64
}
func addStudent(students []Student) []Student {
    var student Student

    fmt.Print("Enter student ID: ")
    fmt.Scan(&student.ID)

    fmt.Print("Enter student name: ")
    fmt.Scan(&student.Name)

    fmt.Print("Enter student age: ")
    fmt.Scan(&student.Age)

    fmt.Print("Enter student score: ")
    fmt.Scan(&student.Score)

    students = append(students, student)

    fmt.Println("\nStudent added successfully!")

    return students
}

func displayStudents(students []Student) {
	for _, student := range students {
		fmt.Printf(
			"ID: %d | Name: %s | Age: %d | Score: %.2f\n",
			student.ID,
			student.Name,
			student.Age,
			student.Score,
		)
	}
}

func Dent() {
	students := []Student{}
	
}