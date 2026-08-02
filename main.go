package main

import (
	"fmt"
	"time"
)

func menu() {
	r := `==== Student Manager ====

	1. Add Student
	2. View Students
	3. Exit

	Select ? 
	`
	fmt.Println(r)
}

func main(){
	students := []Student{}
	for {
		menu()
		var choice int

		fmt.Print("Select: ")
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Invalid Input.")
			continue
		}
		// read user's choice
	
		switch choice {
		case 1:
			// add
			students = addStudent(students)
			time.Sleep(time.Second)
		case 2:
			// view
			displayStudents(students)
			time.Sleep(time.Second)
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid Option.")
			time.Sleep(time.Second)
		}
	}
}