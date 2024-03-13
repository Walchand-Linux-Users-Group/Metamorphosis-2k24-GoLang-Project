package main

import (
	"fmt"
	"os"
	features "student-db/features"
)

func displayChoices() {
	fmt.Println("Choose one of the following")
	fmt.Println("2. Show Students")
	fmt.Println("3. Add New Student")
	fmt.Println("4. Delete student")
	fmt.Println("5. Update student")
	fmt.Println("6. Exit")
}

func main() {
	fmt.Println("Welcome to Student DB!!")

	var c features.Class
	c.NewClass()

	for {
		displayChoices()
		var choices int
		fmt.Scanln(&choices)
		switch choices {
		case 2:
			c.ShowStudents()
		case 3:
			c.AddStudent()
		case 4:
			c.DeleteStudent()
		case 5:
			c.UpdateStudent()
		case 6:
			os.Exit(0)
		default:
			fmt.Println("Invalid input!")
		}
	}

}
