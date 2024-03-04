# Meta2k24-GoLang-Project

### Folder Structure

```bash
Student-DB
├── main.go
├── db.txt
└── features
    └── features.go
```
---

### main.go starter code

```go 
package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello, World!")
}

```
---

### features.go 

#### Structs

```go
package features

//Note these are final dependencies and if they give any erro just comment unused ones for time being.
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Student struct {
	Name   string
	RollNo string
	Marks  int
}

type Class struct {
	Engineers []Student
}

```
---

#### Taking input from user
  
  ```go
  func inputName() string {
	fmt.Println("Enter your first name")
	var name string
	fmt.Scanln(&name)
	return name
}

func inputRoll() string {
	fmt.Println("Enter your roll no")
	var roll string
	fmt.Scanln(&roll)
	return roll
}

func inputMarks() int {
	fmt.Println("Enter your marks")
	var marks int
	fmt.Scanln(&marks)
	return marks
}
```
--- 

#### New Class Function

```go
func (c *Class) NewClass() {

	// fmt.Println(entry)
	var studentArray []Student
	file, err := os.Open("db.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		marks, _ := strconv.Atoi(fields[2])

		entry := Student{
			Name:   fields[0],
			RollNo: fields[1],
			Marks:  marks,
		}

		studentsArray = append(studentsArray, entry)

	}

	// fmt.Println(studentsArray)
	c.Engineers = studentsArray

}
```
---

#### Add Student Function

```go
func (c *Class) AddStudent() {

	name := inputName()

	roll := inputRoll()

	marks := inputMarks()

	entry := Student{
		Name:   name,
		RollNo: roll,
		Marks:  marks,
	}

	c.Engineers = append(c.Engineers, entry)

	fmt.Println("Student Added Successfully")
	fmt.Println()

	//Update this data in the db.txt
}
```
---

#### Show Students Function

```go
func (c *Class) ShowStudents() {
	for _, student := range c.Engineers {
		name, roll, marks := student.Name, student.RollNo, student.Marks
		fmt.Printf("Name: %s\nRoll No: %s\nMarks: %d\n", name, roll, marks)
		fmt.Println()
	}
}
```
---




