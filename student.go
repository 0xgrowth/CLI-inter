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
	if _, err := fmt.Scan(&student.ID); err != nil {
		fmt.Println("Invalid ID.")
		return students
	}
	for _, s := range students {
		if s.ID == student.ID {
			fmt.Println("Student ID already exists.")
			return students
		}
	}

    fmt.Print("Enter student name: ")
	if _, err := fmt.Scan(&student.Name); err != nil {
		fmt.Println("Invalid Name")
		return students
	}

    fmt.Print("Enter student age: ")
	fmt.Scan(&student.Age)
	if _, err := fmt.Scan(&student.Age); err != nil {
		fmt.Println("Invalid age.")
		return students
	}
	
	if student.Age <= 0 {
		fmt.Println("Age must be positive.")
		return students
	}


    fmt.Print("Enter student score: ")
    fmt.Scan(&student.Score)
	if _, err := fmt.Scan(&student.Score); err != nil {
		fmt.Println("Invalid score.")
		return students
	}
	
	if student.Score < 0 || student.Score > 100 {
		fmt.Println("Score must be between 0 and 100.")
		return students
	}

    fmt.Println("\nStudent added successfully!")

	students = append(students, student)
    return students
}

func displayStudents(students []Student) {
	if len(students) == 0 {
		fmt.Println("No students found.")
		return
	}

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
