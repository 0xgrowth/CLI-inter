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
	Dent()
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
			if len(students) == 0 {
				fmt.Println("No students found.")
				time.Sleep(time.Second)
				return
			}
			displayStudents(student)
		case 6:
			return
		}
	}
}